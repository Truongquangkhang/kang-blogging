package comment

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/comment"
	"kang-blogging/internal/blogging/domain/detection_client"
	"kang-blogging/internal/blogging/domain/role"
	"kang-blogging/internal/blogging/domain/user"
	"kang-blogging/internal/common/constants"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/model"
	"kang-blogging/internal/common/utils"
	"os"
	"strconv"
	"time"
)

type UpdateCommentParams struct {
	CommentID string
	UserID    string
	Content   string
}

type UpdateCommentResult struct {
	Comment model.Comment
}

type UpdateCommentHandler decorator.UsecaseHandler[UpdateCommentParams, UpdateCommentResult]

type updateCommentHandler struct {
	commentRepo     comment.Repository
	userRepo        user.Repository
	roleRepo        role.Repository
	detectionClient detection_client.IClientAdapter
}

func NewUpdateCommentHandler(
	commentRepo comment.Repository,
	userRepo user.Repository,
	roleRepo role.Repository,
	detectionClient detection_client.IClientAdapter,
	logger *logrus.Entry,
	metrics decorator.MetricsClient,
) UpdateCommentHandler {
	if commentRepo == nil {
		panic("commentRepo is nil")
	}
	if detectionClient == nil {
		panic("detectionClient is nil")
	}
	if roleRepo == nil {
		panic("roleRepo is nil")
	}
	if userRepo == nil {
		panic("userRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[UpdateCommentParams, UpdateCommentResult](
		updateCommentHandler{
			commentRepo:     commentRepo,
			userRepo:        userRepo,
			detectionClient: detectionClient,
			roleRepo:        roleRepo,
		},
		logger,
		metrics,
	)
}

func (g updateCommentHandler) Handle(ctx context.Context, param UpdateCommentParams) (UpdateCommentResult, error) {
	err := param.Validate()
	if err != nil {
		return UpdateCommentResult{}, err
	}
	cmt, err := g.commentRepo.GetCommentById(ctx, param.CommentID)
	if err != nil {
		return UpdateCommentResult{}, err
	}
	if cmt == nil {
		return UpdateCommentResult{}, errors.NewNotFoundError("comment not found")
	}

	// check authentication
	if cmt.UserID != param.UserID {
		roleUser, err := g.roleRepo.GetRoleByUserId(ctx, param.UserID)
		if err != nil {
			return UpdateCommentResult{}, err
		}
		if roleUser.Name != constants.ADMIN_ROLE {
			return UpdateCommentResult{}, errors.NewForbiddenDefaultError()
		}
	} else {
		u, err := g.userRepo.GetUserByID(ctx, param.UserID)
		if err != nil {
			return UpdateCommentResult{}, err
		}
		if utils.PointerInt64ToValue(u.ExpireWarningTime) >= time.Now().Unix() ||
			!u.IsActive {
			return UpdateCommentResult{}, errors.NewBadRequestError("You have been banned comment. PLease try again later")
		}
		cmt.User = *u
	}

	// check content
	mustDetectComment, err := strconv.ParseBool(os.Getenv("TOXICITY_DETECTION_USE"))
	if err == nil && mustDetectComment {
		prediction, err := g.detectionClient.DetectToxicComment(ctx, param.Content)
		if err != nil || prediction == nil {
			logrus.Error("Failed to detect toxicity comment", err)
		} else {
			marshall, err := json.Marshal(prediction)
			if err != nil {
				logrus.Error("Failed to marshal toxicity comment", err)
			}
			cmt.IsToxicity = prediction.IsToxicComment
			cmt.Prediction = utils.ToStringPointerValue(fmt.Sprintf("%s", string(marshall)))
		}
	}

	cmt, err = g.commentRepo.UpdateComment(ctx, cmt)
	if err != nil {
		return UpdateCommentResult{}, err
	}

	return UpdateCommentResult{
		Comment: *cmt,
	}, nil
}

func (p *UpdateCommentParams) Validate() error {
	if p.UserID == "" || p.Content == "" || p.CommentID == "" {
		return errors.NewBadRequestDefaultError()
	}
	return nil
}

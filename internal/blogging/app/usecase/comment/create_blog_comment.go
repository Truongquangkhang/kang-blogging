package comment

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/comment"
	"kang-blogging/internal/blogging/domain/detection_client"
	"kang-blogging/internal/blogging/domain/user"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/model"
	"kang-blogging/internal/common/utils"
	"os"
	"strconv"
	"time"
)

type CreateBlogCommentParams struct {
	BlogID         string
	UserID         string
	Content        string
	ReplyCommentID *string
}

type CreateBlogCommentResult struct {
	Comment model.Comment
}

type CreateBlogCommentHandler decorator.UsecaseHandler[CreateBlogCommentParams, CreateBlogCommentResult]

type createBlogCommentHandler struct {
	commentRepo     comment.Repository
	userRepo        user.Repository
	detectionClient detection_client.IClientAdapter
}

func NewCreateBlogCommentHandler(
	commentRepo comment.Repository,
	userRepo user.Repository,
	detectionClient detection_client.IClientAdapter,
	logger *logrus.Entry,
	metrics decorator.MetricsClient,
) CreateBlogCommentHandler {
	if commentRepo == nil {
		panic("commentRepo is nil")
	}
	if detectionClient == nil {
		panic("detectionClient is nil")
	}
	if userRepo == nil {
		panic("userRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[CreateBlogCommentParams, CreateBlogCommentResult](
		createBlogCommentHandler{
			commentRepo:     commentRepo,
			detectionClient: detectionClient,
			userRepo:        userRepo,
		},
		logger,
		metrics,
	)
}

func (g createBlogCommentHandler) Handle(ctx context.Context, param CreateBlogCommentParams) (CreateBlogCommentResult, error) {
	err := param.Validate()
	if err != nil {
		return CreateBlogCommentResult{}, err
	}
	u, err := g.userRepo.GetUserByID(ctx, param.UserID)
	if err != nil {
		return CreateBlogCommentResult{}, err
	}
	if utils.PointerInt64ToValue(u.ExpireWarningTime) >= time.Now().Unix() ||
		!u.IsActive {
		return CreateBlogCommentResult{}, errors.NewBadRequestError("You have been banned comment. PLease try again later")
	}
	isToxicity := false
	var predictionJson *string
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
			isToxicity = prediction.IsToxicComment
			predictionJson = utils.ToStringPointerValue(fmt.Sprintf("%s", string(marshall)))
		}
	}

	commentId := utils.GenUUID()
	level := 0
	if param.ReplyCommentID != nil {
		level = 1
	}
	// insert responseComment
	_, err = g.commentRepo.InsertComment(ctx, &model.Comment{
		ID:             commentId,
		Content:        param.Content,
		Level:          level,
		ReplyCommentID: param.ReplyCommentID,
		IsToxicity:     isToxicity,
		UserID:         param.UserID,
		BlogID:         param.BlogID,
		Prediction:     predictionJson,
	})
	if err != nil {
		return CreateBlogCommentResult{}, err
	}

	if isToxicity {
		return CreateBlogCommentResult{}, errors.NewBadRequestError("Your comment is toxic")
	}

	responseComment, err := g.commentRepo.GetCommentById(ctx, commentId)
	if err != nil {
		return CreateBlogCommentResult{}, err
	}

	// update before response
	u.TotalComments = u.TotalComments + 1
	responseComment.User = *u
	return CreateBlogCommentResult{
		Comment: *responseComment,
	}, nil
}

func (p *CreateBlogCommentParams) Validate() error {
	emptyStr := ""
	if p.UserID == "" || p.Content == "" || p.BlogID == "" || p.ReplyCommentID == &emptyStr {
		return errors.NewBadRequestError("Invalid params")
	}
	return nil
}

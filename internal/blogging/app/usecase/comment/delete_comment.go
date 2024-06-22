package comment

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/comment"
	"kang-blogging/internal/blogging/domain/role"
	"kang-blogging/internal/blogging/domain/user"
	"kang-blogging/internal/common/constants"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/errors"
)

type DeleteCommentParams struct {
	CommentID string
	UserID    string
}

type DeleteCommentResult struct {
}

type DeleteCommentHandler decorator.UsecaseHandler[DeleteCommentParams, DeleteCommentResult]

type deleteCommentHandler struct {
	commentRepo comment.Repository
	userRepo    user.Repository
	roleRepo    role.Repository
}

func NewDeleteCommentHandler(
	commentRepo comment.Repository,
	userRepo user.Repository,
	roleRepo role.Repository,
	logger *logrus.Entry,
	metrics decorator.MetricsClient,
) DeleteCommentHandler {
	if commentRepo == nil {
		panic("commentRepo is nil")
	}
	if roleRepo == nil {
		panic("roleRepo is nil")
	}
	if userRepo == nil {
		panic("userRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[DeleteCommentParams, DeleteCommentResult](
		deleteCommentHandler{
			commentRepo: commentRepo,
			userRepo:    userRepo,
			roleRepo:    roleRepo,
		},
		logger,
		metrics,
	)
}

func (g deleteCommentHandler) Handle(ctx context.Context, param DeleteCommentParams) (DeleteCommentResult, error) {
	err := param.Validate()
	if err != nil {
		return DeleteCommentResult{}, err
	}

	cmt, err := g.commentRepo.GetCommentById(ctx, param.CommentID)
	if err != nil {
		return DeleteCommentResult{}, err
	}
	if cmt == nil {
		return DeleteCommentResult{}, errors.NewNotFoundError("comment not found")
	}

	if cmt.UserID != param.UserID {
		roleUser, err := g.roleRepo.GetRoleByUserId(ctx, param.UserID)
		if err != nil {
			return DeleteCommentResult{}, err
		}
		if roleUser.Name != constants.ADMIN_ROLE {
			return DeleteCommentResult{}, errors.NewForbiddenDefaultError()
		}
	}

	err = g.commentRepo.DeleteComment(ctx, param.CommentID, cmt.IsDeprecated)
	return DeleteCommentResult{}, err
}

func (p *DeleteCommentParams) Validate() error {
	if p.UserID == "" || p.CommentID == "" {
		return errors.NewBadRequestDefaultError()
	}
	return nil
}

package comment

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/blog_comments"
	"kang-blogging/internal/blogging/domain/comment"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/model"
	"kang-blogging/internal/common/utils"
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
	commentRepo      comment.Repository
	blogCommentsRepo blog_comments.Repository
}

func NewCreateBlogCommentHandler(
	commentRepo comment.Repository,
	blogCommentsRepo blog_comments.Repository,
	logger *logrus.Entry,
	metrics decorator.MetricsClient,
) CreateBlogCommentHandler {
	if commentRepo == nil {
		panic("commentRepo is nil")
	}
	if blogCommentsRepo == nil {
		panic("blogCommentsRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[CreateBlogCommentParams, CreateBlogCommentResult](
		createBlogCommentHandler{
			commentRepo:      commentRepo,
			blogCommentsRepo: blogCommentsRepo,
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
		IsToxicity:     false,
		UserID:         param.UserID,
	})
	responseComment, err := g.commentRepo.GetCommentById(ctx, commentId)
	if err != nil {
		return CreateBlogCommentResult{}, err
	}
	// insert blog responseComment
	_, err = g.blogCommentsRepo.InsertBlogComment(ctx, &model.BlogComment{
		ID:        utils.GenUUID(),
		CommentID: commentId,
		BlogID:    param.BlogID,
	})
	if err != nil {
		return CreateBlogCommentResult{}, err
	}

	return CreateBlogCommentResult{
		Comment: *responseComment,
	}, nil
}

func (p *CreateBlogCommentParams) Validate() error {
	if p.UserID == "" || p.Content == "" || p.BlogID == "" {
		return errors.NewBadRequestError("Invalid params")
	}
	return nil
}

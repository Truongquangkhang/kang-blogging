package comment

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/domain/comment"
	"kang-blogging/internal/common/model"
	"kang-blogging/internal/common/utils"
)

func (r *CommentRepository) GetBlogComments(
	ctx context.Context, param comment.ParamGetBlogComments,
) ([]comment.ResultGetBlogComments, int32, error) {
	var comments []model.Comment
	var count int64
	limit, offset := utils.PagePageSizeToLimitOffset(param.Page, param.PageSize)
	query := r.gdb.DB().WithContext(ctx).Model(&model.Comment{}).
		Preload("User").
		Joins("join blog_comments on comments.id = blog_comments.comment_id").
		Where("blog_comments.blog_id = ?", param.BlogID)

	// Count the number of comments with level 0
	countQuery := r.gdb.DB().WithContext(ctx).Model(&model.Comment{}).
		Preload("User").
		Joins("join blog_comments on comments.id = blog_comments.comment_id").
		Where("blog_comments.blog_id = ? AND comments.level = 0", param.BlogID).Count(&count)
	err := countQuery.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Limit(int(limit)).Offset(int(offset)).Find(&comments).Error
	if err != nil {
		return nil, 0, err
	}
	return mapCommentsToResults(comments), int32(count), nil
}

func mapCommentsToResults(comments []model.Comment) []comment.ResultGetBlogComments {
	var rs []comment.ResultGetBlogComments
	mapRootCommentToReplies := map[string][]model.Comment{}
	for _, c := range comments {
		if c.ReplyCommentID == nil && c.Level == 0 {
			rs = append(rs, comment.ResultGetBlogComments{
				Comment: c,
			})
		} else if c.ReplyCommentID != nil && c.Level > 0 {
			mapRootCommentToReplies[utils.ToStringValue(c.ReplyCommentID)] =
				append(mapRootCommentToReplies[utils.ToStringValue(c.ReplyCommentID)], c)
		}
	}
	var response []comment.ResultGetBlogComments
	for _, c := range rs {
		response = append(response, comment.ResultGetBlogComments{
			Comment: c.Comment,
			Replies: mapRootCommentToReplies[c.Comment.ID],
		})
	}
	return response
}

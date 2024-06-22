package comment

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/domain/comment"
	"kang-blogging/internal/common/model"
	"kang-blogging/internal/common/utils"
)

func (r *CommentRepository) GetCommentsByParams(
	ctx context.Context,
	params comment.ParamsGetComments,
) ([]model.Comment, int32, error) {
	var comments []model.Comment
	var count int64
	query := r.gdb.DB().WithContext(ctx).Preload("User").
		Where("is_deprecated = false").
		Model(&model.Comment{})

	limit, offset := utils.PagePageSizeToLimitOffset(params.Page, params.PageSize)
	if params.SearchName != nil {
		query = query.Where("comments.content LIKE ?", "%"+*params.SearchName+"%")
	}
	if params.IsToxicity != nil {
		query = query.Where("comments.is_toxicity = ?", *params.IsToxicity)
	}
	if len(params.UserIds) > 0 {
		query = query.Where("comments.user_id IN (?)", params.UserIds)
	}

	errCount := query.Count(&count).Error
	if errCount != nil {
		return nil, 0, errCount
	}
	if params.SortBy != nil {
		if *params.SortBy == "created_at" {
			query = query.Order("created_at DESC")
		}
		if *params.SortBy == "total_reply" {
			query = query.
				Select("comments.*, COUNT(DISTINCT(c1.reply_comment_id)) AS total_reply").
				Joins("left join comments c1 on comments.id = c1.reply_comment_id").
				Group("comments.id").Order("total_reply DESC")
			//query = query.Order("total_reply DESC")
		}
	}

	err := query.
		Limit(int(limit)).Offset(int(offset)).
		Find(&comments).Error
	if err != nil {
		return nil, 0, err
	}
	return comments, int32(count), nil
}

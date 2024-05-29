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
	query := r.gdb.DB().WithContext(ctx).Preload("User").Model(&model.Comment{})
	limit, offset := utils.PagePageSizeToLimitOffset(params.Page, params.PageSize)
	if params.SearchName != nil {
		query = query.Where("content LIKE ?", "%"+*params.SearchName+"%")
	}
	if params.IsToxicity != nil {
		query = query.Where("is_toxicity = ?", *params.IsToxicity)
	}
	if len(params.UserIds) > 0 {
		query = query.Where("user_id IN (?)", params.UserIds)
	}
	if params.SortBy != nil {
		if *params.SortBy == "created_at" {
			query = query.Order("created_at DESC")
		}
	}

	err := query.Count(&count).
		Limit(int(limit)).Offset(int(offset)).
		Find(&comments).Error
	if err != nil {
		return nil, 0, err
	}
	return comments, int32(count), nil
}

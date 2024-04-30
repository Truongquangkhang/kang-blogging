package blog

import (
	"context"
	"fmt"
	"kang-blogging/internal/blogging/domain/blog"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/model"
	"kang-blogging/internal/common/utils"
)

func (r BlogRepository) GetBlogsByParam(
	ctx context.Context,
	param blog.BlogsParams,
) ([]model.Blog, int32, error) {
	limit, offset := utils.PagePageSizeToLimitOffset(param.Page, param.PageSize)
	var blogs []model.Blog
	var total int64
	query := r.gdb.DB().WithContext(ctx).Model(model.Blog{})

	if param.SearchBy != nil && param.SearchName != nil {
		switch *param.SearchBy {
		case "title":
			query = query.Where("title LIKE ?", fmt.Sprintf("%%%s%%", *param.SearchName))
		case "summary":
			query = query.Where("summary LIKE ?", fmt.Sprintf("%%%s%%", *param.SearchName))
		default:
			return nil, 0, errors.NewBadRequestError("invalid search name")
		}
	}
	if len(param.AuthorIds) > 0 {
		query = query.Where("author_id IN (?)", param.AuthorIds)
	}
	err := query.Offset(int(offset)).Limit(int(limit)).Find(&blogs).
		Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	return blogs, int32(total), nil
}

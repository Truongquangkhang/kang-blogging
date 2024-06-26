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
	if len(param.CategoryIds) > 0 {
		query = query.Joins("join blog_categories on blog_categories.blog_id = blogs.id").
			Where("blog_categories.category_id IN (?)", param.CategoryIds)
	}
	if param.Published != nil {
		query = query.Where("published = ?", *param.Published)
	}
	if param.IsDeprecated != nil {
		query = query.Where("blogs.is_deprecated = ?", *param.IsDeprecated)
	}
	if param.SortBy != nil {
		switch *param.SortBy {
		case "updated_at":
			query = query.Order("updated_at desc")
		case "total_comment":
			query = query.Order("total_blog_comments desc")
		}
	}

	err := query.
		Preload("User").Preload("Categories").
		Count(&total).Offset(int(offset)).Limit(int(limit)).
		Select("blogs.*, count(comments.id) as total_blog_comments").
		Joins("left join comments on comments.blog_id = blogs.id").
		Group("blogs.id").
		Find(&blogs).Error

	//logger.Info(query)
	if err != nil {
		return nil, 0, err
	}
	return blogs, int32(total), nil
}

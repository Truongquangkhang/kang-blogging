package blog

import (
	"fmt"
	"golang.org/x/net/context"
	"kang-blogging/internal/common/model"
)

func (r BlogRepository) UpdateBlog(
	ctx context.Context,
	blog *model.Blog,
	categoryIds []string,
) (*model.Blog, error) {
	if len(categoryIds) > 0 {
		var categories []model.Category
		err := r.gdb.DB().WithContext(ctx).
			Where("id IN (?)", categoryIds).Find(&categories).Error
		if err != nil {
			return nil, fmt.Errorf("get an error when get categories: %v", err)
		}
		err = r.gdb.DB().WithContext(ctx).
			Where("blog_id IN (?)", blog.ID).Delete(model.BlogCategories{}).Error
		blog.Categories = categories
	}

	err := r.gdb.DB().WithContext(ctx).Save(blog).Error
	if err != nil {
		return nil, err
	}
	return blog, nil
}

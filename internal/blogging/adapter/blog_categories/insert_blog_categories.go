package blog_categories

import (
	"context"
	"kang-blogging/internal/common/model"
)

func (r *BlogCategoriesRepository) InsertBlogCategories(
	ctx context.Context,
	blogCategories []model.BlogCategories,
) ([]model.BlogCategories, error) {
	err := r.gdb.DB().WithContext(ctx).Create(&blogCategories).Error
	if err != nil {
		return nil, err
	}
	return blogCategories, nil
}

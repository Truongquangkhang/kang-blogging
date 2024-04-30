package category

import (
	"context"
	"kang-blogging/internal/common/model"
)

func (r *CategoryRepository) GetCategories(
	ctx context.Context,
	categoryIds []string,
) ([]model.Category, error) {
	var categories []model.Category
	query := r.gdb.DB().WithContext(ctx)
	if len(categoryIds) > 0 {
		query = query.Where("id IN (?)", categoryIds)
	}
	err := query.Find(&categories).Error
	return categories, err
}

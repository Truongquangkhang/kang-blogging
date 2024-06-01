package category

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/common/model"
)

func (r *CategoryRepository) InsertCategory(
	ctx context.Context,
	category model.Category,
) (*model.Category, error) {
	err := r.gdb.DB().WithContext(ctx).
		Model(&model.Category{}).
		Create(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

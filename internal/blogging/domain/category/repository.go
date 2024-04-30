package category

import (
	"context"
	"kang-blogging/internal/common/model"
)

type Repository interface {
	GetCategories(
		ctx context.Context,
		categoryIds []string,
	) ([]model.Category, error)
}

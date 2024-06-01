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

	GetCategoriesByParam(
		ctx context.Context,
		param ParamGetCategories,
	) ([]ResultGetCategories, int32, error)

	InsertCategory(
		ctx context.Context,
		category model.Category,
	) (*model.Category, error)

	UpdateCategory(
		ctx context.Context,
		category model.Category,
	) (*model.Category, error)
}

type ParamGetCategories struct {
	Page         int32
	PageSize     int32
	SearchByName *string
	SortBy       *string
}

type ResultGetCategories struct {
	ID        string
	Name      string
	BlogCount int32
}

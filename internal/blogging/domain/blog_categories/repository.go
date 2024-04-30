package blog_categories

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/common/model"
)

type Repository interface {
	InsertBlogCategories(
		ctx context.Context,
		blogCategories []model.BlogCategories,
	) ([]model.BlogCategories, error)
}

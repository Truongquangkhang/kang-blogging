package category

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/category"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/model"
)

type UpdateCategoryParams struct {
	CategoryID  string
	Name        *string
	Description *string
}

type UpdateCategoryResult struct {
	Category model.Category
}

type UpdateCategoryHandler decorator.UsecaseHandler[UpdateCategoryParams, UpdateCategoryResult]

type updateCategoryHandler struct {
	categoryRepo category.Repository
}

func NewUpdateCategoryHandler(
	categoryRepo category.Repository,
	logger *logrus.Entry,
	metrics decorator.MetricsClient,
) UpdateCategoryHandler {
	if categoryRepo == nil {
		panic("categoryRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[UpdateCategoryParams, UpdateCategoryResult](
		updateCategoryHandler{
			categoryRepo: categoryRepo,
		},
		logger,
		metrics,
	)
}

func (g updateCategoryHandler) Handle(ctx context.Context, param UpdateCategoryParams) (UpdateCategoryResult, error) {
	err := param.Validate()
	if err != nil {
		return UpdateCategoryResult{}, err
	}
	categories, err := g.categoryRepo.GetCategories(ctx, []string{param.CategoryID})
	if err != nil {
		return UpdateCategoryResult{}, err
	}
	if len(categories) <= 0 {
		return UpdateCategoryResult{}, errors.NewNotFoundError("Category not found")
	}

	// Update
	c := categories[0]
	if param.Name != nil {
		c.Name = *param.Name
	}
	if param.Description != nil {
		c.Description = param.Description
	}
	categoryUpdated, err := g.categoryRepo.UpdateCategory(ctx, c)
	if err != nil {
		return UpdateCategoryResult{}, err
	}

	return UpdateCategoryResult{
		Category: *categoryUpdated,
	}, err
}

func (p *UpdateCategoryParams) Validate() error {
	if p.Name == nil && p.Description == nil {
		return errors.NewBadRequestDefaultError()
	}
	return nil
}

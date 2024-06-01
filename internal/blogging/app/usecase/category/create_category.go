package category

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/category"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/model"
	"kang-blogging/internal/common/utils"
)

type CreateCategoryParams struct {
	Name        string
	Description *string
}

type CreateCategoryResult struct {
	Category model.Category
}

type CreateCategoryHandler decorator.UsecaseHandler[CreateCategoryParams, CreateCategoryResult]

type createCategoryHandler struct {
	categoryRepo category.Repository
}

func NewCreateCategoryHandler(
	categoryRepo category.Repository,
	logger *logrus.Entry,
	metrics decorator.MetricsClient,
) CreateCategoryHandler {
	if categoryRepo == nil {
		panic("categoryRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[CreateCategoryParams, CreateCategoryResult](
		createCategoryHandler{
			categoryRepo: categoryRepo,
		},
		logger,
		metrics,
	)
}

func (g createCategoryHandler) Handle(ctx context.Context, param CreateCategoryParams) (CreateCategoryResult, error) {
	err := param.Validate()
	if err != nil {
		return CreateCategoryResult{}, err
	}
	c, err := g.categoryRepo.InsertCategory(ctx, model.Category{
		ID:          utils.GenUUID(),
		Name:        param.Name,
		Description: param.Description,
	})
	if err != nil {
		return CreateCategoryResult{}, err
	}

	return CreateCategoryResult{
		Category: *c,
	}, err
}

func (p *CreateCategoryParams) Validate() error {
	if p.Name == "" {
		return errors.NewBadRequestError("name is required")
	}
	return nil
}

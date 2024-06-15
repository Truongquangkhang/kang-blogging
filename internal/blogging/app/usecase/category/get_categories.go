package category

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/category"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/model"
)

type GetCategoriesParams struct {
	Page       int32
	PageSize   int32
	SearchName *string
	SortBy     *string
}

type GetCategoriesResult struct {
	Categories []category.ResultGetCategories
	Pagination model.Pagination
}

type GetCategoriesHandler decorator.UsecaseHandler[GetCategoriesParams, GetCategoriesResult]

type getCategoriesHandler struct {
	categoryRepo category.Repository
}

func NewGetCategoriesHandler(
	categoryRepo category.Repository,
	logger *logrus.Entry,
	metrics decorator.MetricsClient,
) GetCategoriesHandler {
	if categoryRepo == nil {
		panic("categoryRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[GetCategoriesParams, GetCategoriesResult](
		getCategoriesHandler{
			categoryRepo: categoryRepo,
		},
		logger,
		metrics,
	)
}

func (g getCategoriesHandler) Handle(ctx context.Context, param GetCategoriesParams) (GetCategoriesResult, error) {
	err := param.Validate()
	if err != nil {
		return GetCategoriesResult{}, err
	}

	categories, count, err := g.categoryRepo.GetCategoriesByParam(ctx, category.ParamGetCategories{
		Page:         param.Page,
		PageSize:     param.PageSize,
		SearchByName: param.SearchName,
		SortBy:       param.SortBy,
	})
	if err != nil {
		return GetCategoriesResult{}, err
	}
	return GetCategoriesResult{
		Categories: categories,
		Pagination: model.Pagination{
			Total:    count,
			Page:     param.Page,
			PageSize: param.PageSize,
		},
	}, err
}

func (p *GetCategoriesParams) Validate() error {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
	if p.SortBy != nil {
		if *p.SortBy != "total_blog" && *p.SortBy != "created_at" {
			return errors.NewBadRequestError("invalid param sortBy")
		}
	}
	return nil
}

package violation

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/violation"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/model"
	"strings"
)

type GetViolationParams struct {
	Page     int32
	PageSize int32
	Type     *string
	UserIDs  *string
}

type GetViolationResult struct {
	Violations []model.Violation
	Pagination model.Pagination
}

type GetViolationHandler decorator.UsecaseHandler[GetViolationParams, GetViolationResult]

type getViolationHandler struct {
	violationRepo violation.Repository
}

func NewGetViolationHandler(
	violationRepo violation.Repository,
	logger *logrus.Entry,
	metrics decorator.MetricsClient,
) GetViolationHandler {
	if violationRepo == nil {
		panic("violationRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[GetViolationParams, GetViolationResult](
		getViolationHandler{
			violationRepo: violationRepo,
		},
		logger,
		metrics,
	)
}

func (g getViolationHandler) Handle(ctx context.Context, param GetViolationParams) (GetViolationResult, error) {
	err := param.Validate()
	if err != nil {
		return GetViolationResult{}, err
	}
	paramsGetViolations := violation.ParamsGetViolations{
		Page:     param.Page,
		PageSize: param.PageSize,
		Type:     param.Type,
	}
	if param.UserIDs != nil {
		paramsGetViolations.UserIDs = strings.Split(*param.UserIDs, ",")
	}
	violations, total, err := g.violationRepo.GetViolationByParams(ctx, paramsGetViolations)
	if err != nil {
		return GetViolationResult{}, err
	}

	return GetViolationResult{
		Violations: violations,
		Pagination: model.Pagination{
			Page:     param.Page,
			PageSize: param.PageSize,
			Total:    total,
		},
	}, err
}

func (p *GetViolationParams) Validate() error {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
	if p.Type != nil {
		if *p.Type != "comment" && *p.Type != "blog" {
			return errors.NewBadRequestError("type must be comment, blog")
		}
	}
	return nil
}

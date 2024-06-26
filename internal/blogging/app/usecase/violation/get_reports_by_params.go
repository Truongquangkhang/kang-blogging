package violation

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/report"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/model"
	"strings"
)

type GetReportsParams struct {
	Page     int32
	PageSize int32
	Type     *string
	IsClosed *bool
	UserIDs  *string
}

type GetReportsResult struct {
	Reports    []model.Report
	Pagination model.Pagination
}

type GetReportsHandler decorator.UsecaseHandler[GetReportsParams, GetReportsResult]

type getReportsHandler struct {
	reportRepo report.Repository
}

func NewGetReportsHandler(
	reportRepo report.Repository,
	logger *logrus.Entry,
	metrics decorator.MetricsClient,
) GetReportsHandler {
	if reportRepo == nil {
		panic("violationRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[GetReportsParams, GetReportsResult](
		getReportsHandler{
			reportRepo: reportRepo,
		},
		logger,
		metrics,
	)
}

func (g getReportsHandler) Handle(ctx context.Context, param GetReportsParams) (GetReportsResult, error) {
	err := param.Validate()
	if err != nil {
		return GetReportsResult{}, err
	}
	paramsGetReports := report.ParamsGetReports{
		Page:     param.Page,
		PageSize: param.PageSize,
		Type:     param.Type,
		IsClosed: param.IsClosed,
	}
	if param.UserIDs != nil {
		paramsGetReports.UserIDs = strings.Split(*param.UserIDs, ",")
	}

	reports, total, err := g.reportRepo.GetReport(ctx, paramsGetReports)
	if err != nil {
		return GetReportsResult{}, err
	}

	return GetReportsResult{
		Reports: reports,
		Pagination: model.Pagination{
			Page:     param.Page,
			PageSize: param.PageSize,
			Total:    total,
		},
	}, err
}

func (p *GetReportsParams) Validate() error {
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

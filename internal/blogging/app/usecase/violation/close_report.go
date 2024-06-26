package violation

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/report"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/errors"
)

type CloseReportParams struct {
	ReportID string
}

type CloseReportResult struct {
}

type CloseReportHandler decorator.UsecaseHandler[CloseReportParams, CloseReportResult]

type closeReportHandler struct {
	reportRepo report.Repository
}

func NewCloseReportHandler(
	reportRepo report.Repository,
	logger *logrus.Entry,
	metrics decorator.MetricsClient,
) CloseReportHandler {
	if reportRepo == nil {
		panic("reportRepo is nil")
	}

	return decorator.ApplyUsecaseDecorators[CloseReportParams, CloseReportResult](
		closeReportHandler{
			reportRepo: reportRepo,
		},
		logger,
		metrics,
	)
}

func (g closeReportHandler) Handle(ctx context.Context, param CloseReportParams) (CloseReportResult, error) {
	err := param.Validate()
	if err != nil {
		return CloseReportResult{}, err
	}

	err = g.reportRepo.UpdateStatusReport(ctx, param.ReportID, false)
	return CloseReportResult{}, err
}

func (p *CloseReportParams) Validate() error {
	if p.ReportID == "" {
		return errors.NewBadRequestError("invalid params")
	}
	return nil
}

package violation

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/comment"
	"kang-blogging/internal/blogging/domain/report"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/model"
	"kang-blogging/internal/common/utils"
)

type CreateReportParams struct {
	Type        string
	TargetID    string
	Reason      string
	Description *string
	ReporterID  string
}

type CreateReportResult struct {
}

type CreateReportHandler decorator.UsecaseHandler[CreateReportParams, CreateReportResult]

type createReportHandler struct {
	reportRepo  report.Repository
	commentRepo comment.Repository
}

func NewCreateReportHandler(
	reportRepo report.Repository,
	commentRepo comment.Repository,
	logger *logrus.Entry,
	metrics decorator.MetricsClient,
) CreateReportHandler {
	if reportRepo == nil {
		panic("reportRepo is nil")
	}
	if commentRepo == nil {
		panic("commentRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[CreateReportParams, CreateReportResult](
		createReportHandler{
			reportRepo:  reportRepo,
			commentRepo: commentRepo,
		},
		logger,
		metrics,
	)
}

func (g createReportHandler) Handle(ctx context.Context, param CreateReportParams) (CreateReportResult, error) {
	err := param.Validate()
	if err != nil {
		return CreateReportResult{}, err
	}
	if param.Type == "comment" {
		cmt, err := g.commentRepo.GetCommentById(ctx, param.TargetID)
		if err != nil {
			return CreateReportResult{}, err
		}
		if cmt == nil {
			return CreateReportResult{}, errors.NewNotFoundError("comment not found")
		}
	}

	reportId := utils.GenUUID()
	rp := model.Report{
		ID:             reportId,
		ReportType:     param.Type,
		ReportTargetID: param.TargetID,
		ReporterID:     param.ReporterID,
		Reason:         param.Reason,
		Description:    param.Description,
	}
	_, err = g.reportRepo.CreateReport(ctx, &rp)
	if err != nil {
		return CreateReportResult{}, err
	}

	return CreateReportResult{}, err
}

func (p *CreateReportParams) Validate() error {
	if p.Type == "" || p.Reason == "" || p.TargetID == "" || p.ReporterID == "" {
		return errors.NewBadRequestError("invalid params")
	}
	return nil
}

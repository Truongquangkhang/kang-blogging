package report

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/common/model"
)

type Repository interface {
	CreateReport(
		ctx context.Context,
		report *model.Report,
	) (*model.Report, error)

	GetReport(
		ctx context.Context,
		params ParamsGetReports,
	) ([]model.Report, int32, error)
}

type ParamsGetReports struct {
	Page     int32
	PageSize int32
	Type     *string
	UserIDs  []string
	IsClosed *bool
}

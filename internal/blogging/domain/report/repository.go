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
}

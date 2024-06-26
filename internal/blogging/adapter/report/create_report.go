package report

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/common/model"
)

func (r *ReportRepository) CreateReport(
	ctx context.Context,
	report *model.Report,
) (*model.Report, error) {
	err := r.gdb.DB().
		WithContext(ctx).
		Model(&model.Report{}).
		Create(&report).Error
	return report, err
}

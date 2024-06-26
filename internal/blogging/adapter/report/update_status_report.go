package report

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/common/model"
)

func (r *ReportRepository) UpdateStatusReport(
	ctx context.Context,
	reportId string,
	currentStatus bool,
) error {
	return r.gdb.DB().WithContext(ctx).
		Model(&model.Report{}).
		Where("id = ?", reportId).
		Updates(map[string]interface{}{"is_closed": !currentStatus}).
		Error
}

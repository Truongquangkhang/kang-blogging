package report

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/domain/report"
	"kang-blogging/internal/common/model"
	"kang-blogging/internal/common/utils"
)

func (r *ReportRepository) GetReport(
	ctx context.Context,
	params report.ParamsGetReports,
) ([]model.Report, int32, error) {
	var reports []model.Report
	var total int64
	limit, offset := utils.PagePageSizeToLimitOffset(params.Page, params.PageSize)

	query := r.gdb.DB().WithContext(ctx).Model(&model.Report{})

	if params.Type != nil {
		query = query.Where("report_type = ?", *params.Type)
	}
	if params.IsClosed != nil {
		query = query.Where("is_closed = ?", *params.IsClosed)
	}
	if len(params.UserIDs) > 0 {
		query = query.Where("reporter_id IN (?)", params.UserIDs)
	}

	err := query.Count(&total).
		Preload("User").
		Order("created_at DESC").
		Limit(int(limit)).Offset(int(offset)).Find(&reports).Error

	if err != nil {
		return nil, 0, err
	}
	return reports, int32(total), nil
}

package violation

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/domain/violation"
	"kang-blogging/internal/common/model"
	"kang-blogging/internal/common/utils"
)

func (r *ViolationRepository) GetViolationByParams(
	ctx context.Context,
	params violation.ParamsGetViolations,
) ([]model.Violation, int32, error) {
	var violations []model.Violation
	var total int64
	limit, offset := utils.PagePageSizeToLimitOffset(params.Page, params.PageSize)

	query := r.gdb.DB().WithContext(ctx).Model(&model.Violation{})
	if params.Type != nil {
		query = query.Where("violation_type = ?", *params.Type)
	}
	if len(params.UserIDs) > 0 {
		query = query.Where("user_id IN (?)", params.UserIDs)
	}
	err := query.
		Preload("User").
		Count(&total).
		Order("created_at DESC").
		Limit(int(limit)).Offset(int(offset)).
		Find(&violations).
		Error
	if err != nil {
		return nil, 0, err
	}
	return violations, int32(total), nil
}

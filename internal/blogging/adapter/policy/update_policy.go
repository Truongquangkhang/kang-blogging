package policy

import (
	"golang.org/x/net/context"
	"gorm.io/gorm/clause"
	"kang-blogging/internal/common/model"
)

func (r *PolicyRepository) UpdatePolicies(
	ctx context.Context,
	policies []model.Policy,
) ([]model.Policy, error) {
	err := r.gdb.DB().WithContext(ctx).Model(&model.Policy{}).
		Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(policies).Error
	return policies, err
}

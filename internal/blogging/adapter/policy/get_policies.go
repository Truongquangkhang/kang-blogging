package policy

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/common/model"
)

func (r *PolicyRepository) GetPolicies(
	ctx context.Context,
) ([]model.Policy, error) {
	var policies []model.Policy
	err := r.gdb.DB().WithContext(ctx).Model(&model.Policy{}).
		Find(&policies).Error
	if err != nil {
		return nil, err
	}
	return policies, nil
}

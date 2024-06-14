package follow

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/common/model"
)

func (r *FollowRepository) DeleteFollow(
	ctx context.Context,
	follow *model.Follow,
) (*model.Follow, error) {
	err := r.gdb.DB().WithContext(ctx).Delete(&follow).Error
	return follow, err
}

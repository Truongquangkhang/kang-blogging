package follow

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/common/model"
)

func (r *FollowRepository) GetFollowedIdsByFollowerId(
	ctx context.Context,
	followerId string,
) ([]string, error) {
	var followedIds []string
	err := r.gdb.DB().WithContext(ctx).Model(&model.Follow{}).
		Select("followed_id").
		Where("follower_id = ?", followerId).
		Find(&followedIds).Error
	if err != nil {
		return nil, err
	}
	return followedIds, nil
}

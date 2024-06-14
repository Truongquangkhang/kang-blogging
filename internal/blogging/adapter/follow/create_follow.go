package follow

import (
	"golang.org/x/net/context"
	"gorm.io/gorm/clause"
	"kang-blogging/internal/common/model"
)

func (r *FollowRepository) CreateFollow(
	ctx context.Context,
	follow *model.Follow,
) (*model.Follow, error) {
	err := r.gdb.DB().WithContext(ctx).
		Clauses(clause.OnConflict{
			DoNothing: true,
		}).
		Create(&follow).Error
	return follow, err
}

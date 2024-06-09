package user

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/common/model"
	"time"
)

func (r UserRepository) ChangeStatus(
	ctx context.Context,
	userId string,
	currentStatus bool,
) error {
	now := time.Now()
	err := r.gdb.DB().WithContext(ctx).Model(&model.User{}).
		Where("id = ?", userId).
		Updates(map[string]interface{}{
			"is_active":  !currentStatus,
			"deleted_at": now,
		}).Error

	return err
}

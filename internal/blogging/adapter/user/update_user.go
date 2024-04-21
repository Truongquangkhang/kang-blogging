package user

import (
	"context"
	"kang-blogging/internal/common/model"
)

func (u UserRepository) UpdateUser(
	ctx context.Context,
	user *model.User,
) (*model.User, error) {
	err := u.gdb.DB().WithContext(ctx).Save(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

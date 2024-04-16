package user

import (
	"context"
	"kang-blogging/internal/common/model"
)

func (u UserRepository) InsertUser(
	ctx context.Context,
	user *model.User,
) (*model.User, error) {
	err := u.gdb.DB().WithContext(ctx).Create(&user).Error
	if err != nil {
		return nil, err
	}
	return nil, err
}

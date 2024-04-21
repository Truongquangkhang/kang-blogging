package user

import (
	"context"
	"kang-blogging/internal/blogging/domain/user"
	"kang-blogging/internal/common/model"
)

func (u UserRepository) UpdateUser(
	ctx context.Context,
	user *user.UserInfo,
) (*model.User, error) {
	mapUpdate := map[string]interface{}{}
	if user.Name != nil {
		mapUpdate["name"] = *user.Name
	}
	if user.DisplayName != nil {
		mapUpdate["display_name"] = *user.DisplayName
	}
	if user.Avatar != nil {
		mapUpdate["avatar"] = *user.Avatar
	}
	if user.Email != nil {
		mapUpdate["email"] = *user.Email
	}
	if user.Gender != nil {
		mapUpdate["gender"] = *user.Gender
	}
	if user.PhoneNumber != nil {
		mapUpdate["phone_number"] = *user.PhoneNumber
	}
	var rs model.User
	err := u.gdb.DB().WithContext(ctx).Model(&rs).
		Where("id = ?", user.ID).Updates(mapUpdate).Error
	if err != nil {
		return nil, err
	}
	return &rs, nil
}

package user

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"kang-blogging/internal/common/model"
)

func (u UserRepository) GetUserByID(
	ctx context.Context,
	id string,
) (*model.User, error) {
	var rs *model.User
	var countComments int32
	errCountComment := u.gdb.DB().WithContext(ctx).Model(&model.Comment{}).
		Select("COUNT(1)").Where("user_id = ?", id).
		Find(&countComments).Error
	if errCountComment != nil {
		return nil, errCountComment
	}

	err := u.gdb.DB().WithContext(ctx).Model(model.User{}).
		First(&rs, "id = ?", id).Error

	if err != nil || rs == nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	rs.TotalComments = countComments
	return rs, nil
}

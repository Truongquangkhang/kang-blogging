package account

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"kang-blogging/internal/common/model"
)

func (u AccountRepository) GetAccountByUsername(
	ctx context.Context,
	username string,
) (*model.Account, error) {
	var rs *model.Account
	err := u.gdb.DB().WithContext(ctx).Model(model.Account{}).
		First(&rs, "username = ?", username).Error

	if err != nil || rs == nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil

}

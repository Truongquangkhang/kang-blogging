package account

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"kang-blogging/internal/blogging/domain/account"
	"kang-blogging/internal/common/model"
)

func (u AccountRepository) GetAccountByUsername(
	ctx context.Context,
	username string,
) (*account.Account, error) {
	var rs model.Account
	err := u.gdb.DB().WithContext(ctx).Model(model.Account{}).
		First(&rs, "username = ?", username).Error

	if err != nil || &rs.ID == nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &account.Account{
		Username: rs.Username,
		Password: rs.Password,
	}, nil

}

package account

import (
	"errors"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"kang-blogging/internal/common/model"
)

func (u AccountRepository) GetAccountById(
	ctx context.Context,
	id string,
) (*model.Account, error) {
	var account model.Account
	err := u.gdb.DB().WithContext(ctx).Model(&model.Account{}).
		Preload("User").
		Where("id = ?", id).First(&account).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &account, nil
}

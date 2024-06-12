package account

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/common/model"
)

func (u AccountRepository) UpdateAccount(
	ctx context.Context,
	account model.Account,
) (*model.Account, error) {
	err := u.gdb.DB().WithContext(ctx).Model(&model.Account{}).
		Where("id = ?", account.ID).
		Updates(map[string]interface{}{"password": account.Password}).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

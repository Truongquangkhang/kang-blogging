package account

import (
	"context"
	"fmt"
	"kang-blogging/internal/blogging/domain/account"
	"kang-blogging/internal/common/model"
)

func (u AccountRepository) GetAccountByUsername(
	ctx context.Context,
	username string,
) (*account.Account, error) {
	var rs model.Account
	err := u.gdb.DB().WithContext(ctx).Model(model.Account{}).
		Find(&rs, "username = ?", username).Error

	if err != nil || &rs.ID == nil {
		fmt.Printf("error: %v \n", err)
		fmt.Printf("rs: %v\n", rs)
		return nil, err
	}
	return &account.Account{
		Username: rs.Username,
		Password: rs.Password,
	}, nil

}

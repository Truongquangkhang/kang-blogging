package account

import (
	"context"
	"fmt"
	"kang-blogging/internal/blogging/domain/account"
	"kang-blogging/internal/common/model"
)

func (u AccountRepository) InsertAccount(
	ctx context.Context,
	id string,
	username string,
	password string,
) (*account.Account, error) {
	account := model.Account{
		ID:       id,
		Username: username,
		Password: password,
	}

	err := u.gdb.DB().Create(&account).Error
	if err != nil {
		logger.WithField("err", err).Error("Error while creating new user")
		return nil, err

	}
	fmt.Printf("account: %#v\n", account)
	return nil, nil
}

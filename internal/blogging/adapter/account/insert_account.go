package account

import (
	"context"
	"kang-blogging/internal/blogging/domain/account"
	"kang-blogging/internal/common/model"
)

func (u AccountRepository) InsertAccount(
	ctx context.Context,
	id string,
	username string,
	password string,
) (*account.Account, error) {
	acc := model.Account{
		ID:       id,
		Username: username,
		Password: password,
	}

	err := u.gdb.DB().Create(&acc).Error
	if err != nil {
		logger.WithField("err", err).Error("Error while creating new user")
		return nil, err
	}
	return &account.Account{
		ID:       id,
		Username: username,
		Password: password,
	}, nil
}

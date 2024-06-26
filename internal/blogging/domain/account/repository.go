package account

import (
	"context"
	"kang-blogging/internal/common/model"
)

type Repository interface {
	InsertAccount(
		ctx context.Context,
		id string,
		username string,
		password string,
	) (*Account, error)

	GetAccountByUsername(
		ctx context.Context,
		username string,
	) (*model.Account, error)

	GetRoleUserByID(
		ctx context.Context,
		id string,
	) (*model.Role, error)

	GetAccountById(
		ctx context.Context,
		id string,
	) (*model.Account, error)

	UpdateAccount(
		ctx context.Context,
		account model.Account,
	) (*model.Account, error)
}

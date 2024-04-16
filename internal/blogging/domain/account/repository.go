package account

import "context"

type Repository interface {
	InsertAccount(
		ctx context.Context,
		id string,
		username string,
		password string,
	) (*Account, error)

	GetAccountByUsername(ctx context.Context, username string) (*Account, error)
}

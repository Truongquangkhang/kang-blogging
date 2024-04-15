package account

type Repository interface {
	InsertAccount(
		username string,
		password string,
	) (*Account, error)
}

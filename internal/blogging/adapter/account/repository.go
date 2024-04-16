package account

import (
	"github.com/sirupsen/logrus"
	gormAdapter "kang-blogging/internal/common/db"
)

type AccountRepository struct {
	gdb gormAdapter.DBAdapter
}

const repoName = "UserRepository"

var logger *logrus.Entry = logrus.StandardLogger().
	WithField("logger", repoName)

func NewRepository() *AccountRepository {
	return &AccountRepository{
		gdb: gormAdapter.GetDBInstance(),
	}
}

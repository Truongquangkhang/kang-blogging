package user

import (
	"github.com/sirupsen/logrus"
	gormAdapter "kang-blogging/internal/common/db"
)

type UserRepository struct {
	gdb gormAdapter.DBAdapter
}

const repoName = "UserRepository"

var logger *logrus.Entry = logrus.StandardLogger().
	WithField("logger", repoName)

func NewRepository() *UserRepository {
	return &UserRepository{
		gdb: gormAdapter.GetDBInstance(),
	}
}

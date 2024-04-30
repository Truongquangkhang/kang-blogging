package blog

import (
	"github.com/sirupsen/logrus"
	gormAdapter "kang-blogging/internal/common/db"
)

type BlogRepository struct {
	gdb gormAdapter.DBAdapter
}

const repoName = "UserRepository"

var logger *logrus.Entry = logrus.StandardLogger().
	WithField("logger", repoName)

func NewRepository() *BlogRepository {
	return &BlogRepository{
		gdb: gormAdapter.GetDBInstance(),
	}
}

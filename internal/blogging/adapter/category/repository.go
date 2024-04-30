package category

import (
	"github.com/sirupsen/logrus"
	gormAdapter "kang-blogging/internal/common/db"
)

type CategoryRepository struct {
	gdb gormAdapter.DBAdapter
}

const repoName = "CategoryRepository"

var logger *logrus.Entry = logrus.StandardLogger().
	WithField("logger", repoName)

func NewRepository() *CategoryRepository {
	return &CategoryRepository{
		gdb: gormAdapter.GetDBInstance(),
	}
}

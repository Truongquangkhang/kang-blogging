package role

import (
	"github.com/sirupsen/logrus"
	gormAdapter "kang-blogging/internal/common/db"
)

type RoleRepository struct {
	gdb gormAdapter.DBAdapter
}

const repoName = "RoleRepository"

var logger *logrus.Entry = logrus.StandardLogger().
	WithField("logger", repoName)

func NewRepository() *RoleRepository {
	return &RoleRepository{
		gdb: gormAdapter.GetDBInstance(),
	}
}

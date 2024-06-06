package policy

import (
	"github.com/sirupsen/logrus"
	gormAdapter "kang-blogging/internal/common/db"
)

type PolicyRepository struct {
	gdb gormAdapter.DBAdapter
}

const repoName = "PolicyRepository"

var logger *logrus.Entry = logrus.StandardLogger().
	WithField("logger", repoName)

func NewRepository() *PolicyRepository {
	return &PolicyRepository{
		gdb: gormAdapter.GetDBInstance(),
	}
}

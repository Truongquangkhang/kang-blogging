package violation

import (
	"github.com/sirupsen/logrus"
	gormAdapter "kang-blogging/internal/common/db"
)

type ViolationRepository struct {
	gdb gormAdapter.DBAdapter
}

const repoName = "ViolationRepository"

var logger *logrus.Entry = logrus.StandardLogger().
	WithField("logger", repoName)

func NewRepository() *ViolationRepository {
	return &ViolationRepository{
		gdb: gormAdapter.GetDBInstance(),
	}
}

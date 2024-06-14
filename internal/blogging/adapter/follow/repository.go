package follow

import (
	"github.com/sirupsen/logrus"
	gormAdapter "kang-blogging/internal/common/db"
)

type FollowRepository struct {
	gdb gormAdapter.DBAdapter
}

const repoName = "FollowRepository"

var logger *logrus.Entry = logrus.StandardLogger().
	WithField("logger", repoName)

func NewRepository() *FollowRepository {
	return &FollowRepository{
		gdb: gormAdapter.GetDBInstance(),
	}
}

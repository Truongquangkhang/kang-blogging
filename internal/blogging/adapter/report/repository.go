package report

import (
	"github.com/sirupsen/logrus"
	gormAdapter "kang-blogging/internal/common/db"
)

type ReportRepository struct {
	gdb gormAdapter.DBAdapter
}

const repoName = "ReportRepository"

var logger *logrus.Entry = logrus.StandardLogger().
	WithField("logger", repoName)

func NewRepository() *ReportRepository {
	return &ReportRepository{
		gdb: gormAdapter.GetDBInstance(),
	}
}

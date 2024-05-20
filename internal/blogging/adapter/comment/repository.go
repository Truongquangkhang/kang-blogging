package comment

import (
	"github.com/sirupsen/logrus"
	gormAdapter "kang-blogging/internal/common/db"
)

type CommentRepository struct {
	gdb gormAdapter.DBAdapter
}

const repoName = "CommentRepository"

var logger *logrus.Entry = logrus.StandardLogger().
	WithField("logger", repoName)

func NewRepository() *CommentRepository {
	return &CommentRepository{
		gdb: gormAdapter.GetDBInstance(),
	}
}

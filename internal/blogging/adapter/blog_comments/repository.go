package blog_comments

import (
	"github.com/sirupsen/logrus"
	gormAdapter "kang-blogging/internal/common/db"
)

type BlogCommentsRepository struct {
	gdb gormAdapter.DBAdapter
}

const repoName = "BlogCommentsRepository"

var logger *logrus.Entry = logrus.StandardLogger().
	WithField("logger", repoName)

func NewRepository() *BlogCommentsRepository {
	return &BlogCommentsRepository{
		gdb: gormAdapter.GetDBInstance(),
	}
}

package blog_categories

import (
	"github.com/sirupsen/logrus"
	gormAdapter "kang-blogging/internal/common/db"
)

type BlogCategoriesRepository struct {
	gdb gormAdapter.DBAdapter
}

const repoName = "BlogCategoriesRepository"

var logger *logrus.Entry = logrus.StandardLogger().
	WithField("logger", repoName)

func NewRepository() *BlogCategoriesRepository {
	return &BlogCategoriesRepository{
		gdb: gormAdapter.GetDBInstance(),
	}
}

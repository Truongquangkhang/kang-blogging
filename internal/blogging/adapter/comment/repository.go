package comment

import (
	"errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	gormAdapter "kang-blogging/internal/common/db"
	"kang-blogging/internal/common/model"
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

func (r *CommentRepository) GetCommentById(
	ctx context.Context, commentId string,
) (*model.Comment, error) {
	var comment *model.Comment
	err := r.gdb.DB().WithContext(ctx).Preload("User").Where("id = ?", commentId).First(&comment).Error
	if err != nil || comment == nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return comment, nil
}

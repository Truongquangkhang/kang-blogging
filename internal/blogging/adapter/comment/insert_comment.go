package comment

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/net/context"
	errors2 "kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/model"
)

func (r *CommentRepository) InsertComment(
	ctx context.Context,
	comment *model.Comment,
) (*model.Comment, error) {
	result := r.gdb.DB().WithContext(ctx).Create(&comment)
	if result.Error != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(result.Error, &mysqlErr) {
			if mysqlErr.Number == 1452 {
				return nil, errors2.NewBadRequestError("Invalid blog id")
			}
		}
		return &model.Comment{}, result.Error
	}
	return comment, nil
}

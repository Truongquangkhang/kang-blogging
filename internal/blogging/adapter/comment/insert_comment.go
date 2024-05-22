package comment

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/common/model"
)

func (r *CommentRepository) InsertComment(
	ctx context.Context,
	comment *model.Comment,
) (*model.Comment, error) {
	err := r.gdb.DB().WithContext(ctx).Create(&comment).Error
	if err != nil {
		return &model.Comment{}, err
	}
	return comment, nil
}

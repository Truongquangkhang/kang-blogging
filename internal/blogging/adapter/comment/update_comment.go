package comment

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/common/model"
)

func (r *CommentRepository) UpdateComment(
	ctx context.Context,
	comment *model.Comment,
) (*model.Comment, error) {
	err := r.gdb.DB().WithContext(ctx).Model(&model.Comment{}).
		Where("id = ?", comment.ID).
		Updates(&comment).Error
	return comment, err
}

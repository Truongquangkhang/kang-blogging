package comment

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/common/model"
)

func (g *CommentRepository) DeleteComment(
	ctx context.Context,
	commentID string,
) error {
	return g.gdb.DB().WithContext(ctx).
		Where("id = ?", commentID).Delete(&model.Comment{}).Error
}

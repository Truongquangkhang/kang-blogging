package comment

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/common/model"
	"time"
)

func (g *CommentRepository) DeleteComment(
	ctx context.Context,
	commentID string,
	currentStatus bool,
) error {
	now := time.Now()
	return g.gdb.DB().WithContext(ctx).
		Model(&model.Comment{}).
		Where("id = ?", commentID).
		Updates(map[string]interface{}{"is_deprecated": !currentStatus, "deleted_at": &now}).
		Error
}

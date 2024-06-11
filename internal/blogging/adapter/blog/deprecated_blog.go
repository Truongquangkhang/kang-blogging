package blog

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/common/model"
	"time"
)

func (r BlogRepository) ChangeDeprecatedBlog(
	ctx context.Context,
	blogId string,
	currentStatus bool,
) error {
	now := time.Now()
	err := r.gdb.DB().WithContext(ctx).Model(&model.Blog{}).
		Where("id = ?", blogId).
		Updates(map[string]interface{}{"is_deprecated": !currentStatus, "deleted_at": &now}).Error
	return err
}

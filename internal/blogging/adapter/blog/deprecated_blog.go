package blog

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/common/model"
	"time"
)

func (r BlogRepository) DeprecatedBlog(
	ctx context.Context,
	blogId string,
) error {
	now := time.Now()
	err := r.gdb.DB().WithContext(ctx).Model(&model.Blog{}).
		Where("id = ?", blogId).
		Updates(&model.Blog{IsDeprecated: true, DeletedAt: &now}).Error
	return err
}

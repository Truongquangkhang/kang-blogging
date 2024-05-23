package blog

import (
	"context"
	"kang-blogging/internal/common/model"
)

func (r *BlogRepository) InsertBlog(
	ctx context.Context,
	blog *model.Blog,
) (*model.Blog, error) {
	err := r.gdb.DB().WithContext(ctx).Omit("TotalBlogComments").
		Create(&blog).Error
	if err != nil {
		return nil, err
	}
	return blog, nil
}

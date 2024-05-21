package blog_comments

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/common/model"
)

func (r *BlogCommentsRepository) InsertBlogComment(
	ctx context.Context,
	blogComment *model.BlogComment,
) (*model.BlogComment, error) {
	err := r.gdb.DB().WithContext(ctx).Create(&blogComment).Error
	if err != nil {
		return &model.BlogComment{}, err
	}
	return blogComment, nil
}

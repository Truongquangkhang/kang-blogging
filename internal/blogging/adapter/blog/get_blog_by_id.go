package blog

import (
	"errors"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"kang-blogging/internal/common/model"
)

func (r BlogRepository) GetBlogByID(
	ctx context.Context,
	blogId string,
) (*model.Blog, error) {
	var blog *model.Blog
	err := r.gdb.DB().WithContext(ctx).
		Preload("User").Preload("Categories").
		Where("id = ?", blogId).First(&blog).Error
	if err != nil || blog == nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return blog, nil
}

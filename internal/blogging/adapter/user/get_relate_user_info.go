package user

import (
	"errors"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"kang-blogging/internal/blogging/domain/user"
	errors2 "kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/model"
)

func (r UserRepository) GetRelateInfoOfUser(
	ctx context.Context,
	userId string,
	ignoreBlogIsDraft bool,
) (*user.RelateUserInfo, error) {
	var totalComments int32
	var totalBlogs int32
	var blogs []model.Blog
	var comments []model.Comment
	var userInfo model.User

	// get user info
	errGetUser := r.gdb.DB().WithContext(ctx).Model(&model.User{}).
		Where("id = ?", userId).First(&userInfo).Error
	if errGetUser != nil {
		if errors.Is(errGetUser, gorm.ErrRecordNotFound) {
			return nil, errors2.NewNotFoundError("user not found")
		}
		return nil, errGetUser
	}
	// count comments and blog of this user
	errCountComment := r.gdb.DB().WithContext(ctx).Model(&model.Comment{}).
		Select("COUNT(1)").Where("user_id = ?", userId).Scan(&totalComments).Error
	if errCountComment != nil {
		return nil, errCountComment
	}
	errCountBlog := r.gdb.DB().WithContext(ctx).Model(&model.Blog{}).
		Select("COUNT(1)").Where("author_id = ? AND published = TRUE", userId).Scan(&totalBlogs).Error
	if errCountBlog != nil {
		return nil, errCountBlog
	}
	// get comments and blogs of this user
	errGetComment := r.gdb.DB().WithContext(ctx).Model(&model.Comment{}).
		Where("user_id = ? AND is_toxicity = false", userId).
		Limit(10).
		Find(&comments).Error
	if errGetComment != nil {
		return nil, errGetComment
	}
	queryGetBlogs := r.gdb.DB().WithContext(ctx).Model(&model.Blog{}).
		Preload("Categories").
		Where("author_id = ? AND is_deprecated = false", userId)
	if ignoreBlogIsDraft {
		queryGetBlogs = queryGetBlogs.Where("published = true")
	}
	errGetBlog := queryGetBlogs.Limit(10).Find(&blogs).Error
	if errGetBlog != nil {
		return nil, errGetBlog
	}

	return &user.RelateUserInfo{
		User:          userInfo,
		Blogs:         blogs,
		Comments:      comments,
		TotalComments: totalComments,
		TotalBlogs:    totalBlogs,
	}, nil
}

package user

import "kang-blogging/internal/common/model"

type UserInfo struct {
	ID          string
	Name        *string
	DisplayName *string
	Email       *string
	PhoneNumber *string
	Avatar      *string
	Gender      *bool
	BirthOfDay  *int64
}

type UserParams struct {
	Page          int32
	PageSize      int32
	SearchName    *string
	SearchBy      *string
	Follower      *bool
	Followed      *bool
	IsActive      *bool
	SortBy        *string
	CurrentUserID *string
}

type SystemInfo struct {
	TotalBlogs           int32
	TotalComments        int32
	TotalUsers           int32
	TotalCategories      int32
	BlogIncreaseInDay    int32
	CommentIncreaseInDay int32
	UserIncreaseInDay    int32
}

type RelateUserInfo struct {
	User           model.User
	TotalComments  int32
	TotalBlogs     int32
	TotalFollowers int32
	TotalFolloweds int32
	IsFollower     bool
	IsFollowed     bool
	Blogs          []model.Blog
	Comments       []model.Comment
}

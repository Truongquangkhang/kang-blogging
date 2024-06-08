package user

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
	Page         int32
	PageSize     int32
	SearchName   *string
	SearchBy     *string
	Following    *bool
	FollowedByMe *bool
	IsActive     *bool
	SortBy       *string
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

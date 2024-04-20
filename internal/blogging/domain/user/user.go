package user

type UserInfo struct {
	Name        string
	DisplayName string
	Email       string
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
}

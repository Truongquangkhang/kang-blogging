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

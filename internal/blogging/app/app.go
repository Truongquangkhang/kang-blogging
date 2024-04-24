package app

import (
	"kang-blogging/internal/blogging/app/usecase/blog"
	"kang-blogging/internal/blogging/app/usecase/iam"
	"kang-blogging/internal/blogging/app/usecase/user"
)

type Application struct {
	IAMUsecases IAMUsecases
	UserUsecase UserUsecase
	BlogUsecase BlogUsecase
}

type IAMUsecases struct {
	Login              iam.LoginHandler
	Register           iam.RegisterHandler
	CheckExistUsername iam.CheckExistUsernameHandler
}

type UserUsecase struct {
	GetUsers      user.GetUsersHandler
	GetUserDetail user.GetUserDetailHandler
	UpdateUser    user.UpdateUserHandler
}

type BlogUsecase struct {
	GetBlogs blog.GetBlogsHandler
}

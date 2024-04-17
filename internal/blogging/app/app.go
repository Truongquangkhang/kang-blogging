package app

import (
	"kang-blogging/internal/blogging/app/usecase/iam"
)

type Application struct {
	IAMUsecases IAMUsecases
}

type IAMUsecases struct {
	Login              iam.LoginHandler
	Register           iam.RegisterHandler
	CheckExistUsername iam.CheckExistUsernameHandler
}

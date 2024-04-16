package app

import (
	"kang-blogging/internal/blogging/app/usecase/iam"
)

type Application struct {
	IAMUsecases IAMUsecases
}

type IAMUsecases struct {
	Register           iam.RegisterHandler
	CheckExistUsername iam.CheckExistUsernameHandler
}

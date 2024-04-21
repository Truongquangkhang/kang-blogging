package user

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/user"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/model"
)

type GetUserDetailParams struct {
	ID          string
	Name        *string
	DisplayName *string
	Email       *string
	Avatar      *string
	PhoneNumber *string
	Gender      *bool
}

type GetUserDetailResult struct {
	User model.User
}

type GetUserDetailHandler decorator.UsecaseHandler[GetUserDetailParams, GetUserDetailResult]

type getUserDetailHandler struct {
	userRepo user.Repository
}

func NewGetUserDetailHandler(
	userRepo user.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) GetUserDetailHandler {
	if userRepo == nil {
		panic("userRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[GetUserDetailParams, GetUserDetailResult](
		&getUserDetailHandler{
			userRepo: userRepo,
		},
		logger,
		metricsClient,
	)
}

func (g getUserDetailHandler) Handle(
	ctx context.Context,
	param GetUserDetailParams,
) (GetUserDetailResult, error) {
	return GetUserDetailResult{}, nil
}

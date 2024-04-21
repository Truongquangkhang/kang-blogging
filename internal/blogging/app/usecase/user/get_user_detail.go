package user

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/user"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/model"
)

type GetUserDetailParams struct {
	ID string
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
	err := param.Validate()
	if err != nil {
		return GetUserDetailResult{}, err
	}
	u, err := g.userRepo.GetUserByID(ctx, param.ID)
	if err != nil || u == nil {
		return GetUserDetailResult{}, errors.NewNotFoundError("user not found")
	}

	return GetUserDetailResult{
		User: *u,
	}, nil
}

func (p GetUserDetailParams) Validate() error {
	if p.ID == "" {
		return errors.NewBadRequestError("user ID is required")
	}
	return nil
}

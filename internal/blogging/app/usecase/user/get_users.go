package user

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/user"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/model"
)

type GetUsersParams struct {
	Page         int32
	PageSize     int32
	SearchName   *string
	SearchBy     *string
	Following    *bool
	FollowedByMe *bool
}

type GetUsersResult struct {
	Users      []model.User
	Pagination model.Pagination
}

type GetUsersHandler decorator.UsecaseHandler[GetUsersParams, GetUsersResult]

type getUsersHandler struct {
	userRepo user.Repository
}

func NewGetUserHandler(
	userRepo user.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) GetUsersHandler {
	if userRepo == nil {
		panic("userRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[GetUsersParams, GetUsersResult](
		getUsersHandler{
			userRepo: userRepo,
		},
		logger,
		metricsClient)
}

func (g getUsersHandler) Handle(ctx context.Context, param GetUsersParams) (GetUsersResult, error) {
	err := param.Validate()
	if err != nil {
		return GetUsersResult{}, err
	}
	users, total, err := g.userRepo.GetUsers(ctx, user.UserParams{
		Page:         param.Page,
		PageSize:     param.PageSize,
		SearchBy:     param.SearchBy,
		Following:    param.Following,
		FollowedByMe: param.FollowedByMe,
		SearchName:   param.SearchName,
	})
	if err != nil {
		return GetUsersResult{}, err
	}

	return GetUsersResult{
		Users: users,
		Pagination: model.Pagination{
			Page:     param.Page,
			PageSize: param.PageSize,
			Total:    total,
		},
	}, nil
}

func (p *GetUsersParams) Validate() error {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
	return nil
}
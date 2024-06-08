package user

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/user"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/model"
	"kang-blogging/internal/common/utils"
)

type GetUsersParams struct {
	Page         int32
	PageSize     int32
	SearchName   *string
	SearchBy     *string
	Following    *bool
	FollowedByMe *bool
	IsActive     *bool
	SortBy       *string
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
		IsActive:     param.IsActive,
		SortBy:       param.SortBy,
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
	if p.SortBy == nil {
		p.SortBy = utils.ToStringPointerValue("created_at")
	} else {
		if *p.SortBy != "created_at" && *p.SortBy != "total_violation" &&
			*p.SortBy != "total_blog" && *p.SortBy != "total_comment" {
			return errors.NewBadRequestError("invalid params")
		}
	}
	return nil
}

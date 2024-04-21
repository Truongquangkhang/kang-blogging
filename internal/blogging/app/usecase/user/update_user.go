package user

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/user"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/model"
)

type UpdateUserParams struct {
	ID          string
	Name        *string
	DisplayName *string
	Email       *string
	Avatar      *string
	PhoneNumber *string
	Gender      *bool
}

type UpdateUserResult struct {
	User model.User
}

type UpdateUserHandler decorator.UsecaseHandler[UpdateUserParams, UpdateUserResult]

type updateUserHandler struct {
	userRepo user.Repository
}

func NewUpdateUserHandler(
	userRepo user.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) UpdateUserHandler {
	if userRepo == nil {
		panic("userRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[UpdateUserParams, UpdateUserResult](
		&updateUserHandler{
			userRepo: userRepo,
		},
		logger,
		metricsClient,
	)
}

func (u updateUserHandler) Handle(ctx context.Context, param UpdateUserParams) (UpdateUserResult, error) {
	err := param.Validate()
	if err != nil {
		return UpdateUserResult{}, err
	}
	rs, err := u.userRepo.UpdateUser(ctx, &user.UserInfo{
		ID:          param.ID,
		Name:        param.Name,
		DisplayName: param.DisplayName,
		Email:       param.Email,
		Avatar:      param.Avatar,
		PhoneNumber: param.PhoneNumber,
		Gender:      param.Gender,
	})
	if err != nil {
		return UpdateUserResult{}, err
	}
	return UpdateUserResult{
		User: *rs,
	}, nil
}

func (p *UpdateUserParams) Validate() error {
	if p.ID == "" {
		return errors.NewBadRequestError("id is required")
	}
	return nil
}

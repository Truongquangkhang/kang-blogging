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
	Description *string
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

	// get user
	user, err := u.userRepo.GetUserByID(ctx, param.ID)
	if err != nil || user == nil {
		return UpdateUserResult{}, errors.NewNotFoundError("User not found")
	}
	rs, err := u.userRepo.UpdateUser(ctx, mapUserUpdate(param, user))
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

func mapUserUpdate(p UpdateUserParams, user *model.User) *model.User {
	if p.Name != nil {
		user.Name = *p.Name
	}
	if p.DisplayName != nil {
		user.DisplayName = *p.DisplayName
	}
	if p.Email != nil {
		user.Email = *p.Email
	}
	if p.Avatar != nil {
		user.Avatar = p.Avatar
	}
	if p.PhoneNumber != nil {
		user.PhoneNumber = p.PhoneNumber
	}
	if p.Gender != nil {
		user.Gender = p.Gender
	}
	if p.Description != nil {
		user.Description = p.Description
	}
	return user
}

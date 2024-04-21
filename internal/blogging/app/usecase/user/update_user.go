package user

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/user"
	"kang-blogging/internal/common/decorator"
)

type UpdateUserParams struct {
}

type UpdateUserResult struct {
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
	return UpdateUserResult{}, nil
}

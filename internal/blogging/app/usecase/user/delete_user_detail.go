package user

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/domain/role"
	"kang-blogging/internal/blogging/domain/user"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/errors"
)

type DeleteUserDetailParams struct {
	UserID string
}

type DeleteUserDetailResult struct {
}

type DeleteUserDetailHandler decorator.UsecaseHandler[DeleteUserDetailParams, DeleteUserDetailResult]

type deleteUserDetailHandler struct {
	userRepo user.Repository
	roleRepo role.Repository
}

func NewDeleteUserDetailHandler(
	userRepo user.Repository,
	roleRepo role.Repository,
	logger *logrus.Entry,
	metrics decorator.MetricsClient,
) DeleteUserDetailHandler {
	if userRepo == nil {
		panic("blogRepo is nil")
	}
	if roleRepo == nil {
		panic("roleRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[DeleteUserDetailParams, DeleteUserDetailResult](
		deleteUserDetailHandler{
			userRepo: userRepo,
			roleRepo: roleRepo,
		},
		logger,
		metrics,
	)
}

func (g deleteUserDetailHandler) Handle(ctx context.Context, param DeleteUserDetailParams) (DeleteUserDetailResult, error) {
	err := param.Validate()
	if err != nil {
		return DeleteUserDetailResult{}, err
	}
	u, err := g.userRepo.GetUserByID(ctx, param.UserID)
	if err != nil {
		return DeleteUserDetailResult{}, err
	}
	if u == nil {
		return DeleteUserDetailResult{}, errors.NewNotFoundError("user not found")
	}

	err = g.userRepo.ChangeStatus(ctx, u.ID, u.IsActive)
	if err != nil {
		return DeleteUserDetailResult{}, err
	}

	return DeleteUserDetailResult{}, nil
}

func (p *DeleteUserDetailParams) Validate() error {
	return nil
}

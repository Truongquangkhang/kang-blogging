package iam

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/account"
	"kang-blogging/internal/blogging/domain/role"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/errors"
	util_password "kang-blogging/internal/common/utils/password"
)

type ChangePasswordParams struct {
	UserID             string
	OldPassword        *string
	NewPassword        string
	WithoutOldPassword bool
}

type ChangePasswordResult struct {
}

type ChangePasswordHandler decorator.UsecaseHandler[ChangePasswordParams, ChangePasswordResult]

type changePasswordHandler struct {
	roleRepository role.Repository
	accountRepo    account.Repository
}

func NewChangePasswordHandler(
	roleRepository role.Repository,
	accountRepo account.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) ChangePasswordHandler {
	if roleRepository == nil {
		panic("RoleRepository is nil")
	}
	if accountRepo == nil {
		panic("AccountRepository is nil")
	}
	return decorator.ApplyUsecaseDecorators[ChangePasswordParams, ChangePasswordResult](
		changePasswordHandler{
			roleRepository: roleRepository,
			accountRepo:    accountRepo,
		},
		logger,
		metricsClient,
	)
}

func (l changePasswordHandler) Handle(ctx context.Context, param ChangePasswordParams) (ChangePasswordResult, error) {
	err := param.Validate()
	if err != nil {
		return ChangePasswordResult{}, err
	}

	acc, err := l.accountRepo.GetAccountById(ctx, param.UserID)
	if err != nil {
		return ChangePasswordResult{}, err
	}
	if acc == nil {
		return ChangePasswordResult{}, errors.NewNotFoundError("account not found")
	}

	if !param.WithoutOldPassword {
		// compare old password
		if !util_password.CheckPasswordHash(*param.OldPassword, acc.Password) {
			return ChangePasswordResult{}, errors.NewBadRequestError("old password is invalid")
		}
	}

	hashPassword, err := util_password.HashPassword(param.NewPassword)
	if err != nil {
		return ChangePasswordResult{}, err
	}
	acc.Password = hashPassword

	_, err = l.accountRepo.UpdateAccount(ctx, *acc)
	if err != nil {
		return ChangePasswordResult{}, err
	}

	return ChangePasswordResult{}, nil
}

func (p *ChangePasswordParams) Validate() error {
	if p.NewPassword == "" {
		return errors.NewBadRequestError("Invalid params")
	}

	if !p.WithoutOldPassword && p.OldPassword == nil {
		return errors.NewBadRequestError("Invalid params")
	}
	return nil
}

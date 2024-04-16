package iam

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/account"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/errors"
)

type CheckExistUsernameParam struct {
	Username string
}

type CheckExistUsernameResult struct {
	AlreadyExists bool
}

type CheckExistUsernameHandler decorator.UsecaseHandler[CheckExistUsernameParam, CheckExistUsernameResult]

type checkExistUsernameHandler struct {
	accountRepo account.Repository
}

func NewCheckExistUsernameHandler(
	accountRepo account.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) CheckExistUsernameHandler {
	if accountRepo == nil {
		panic("accountRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[CheckExistUsernameParam, CheckExistUsernameResult](
		checkExistUsernameHandler{
			accountRepo: accountRepo,
		},
		logger,
		metricsClient)
}

func (c checkExistUsernameHandler) Handle(
	ctx context.Context,
	param CheckExistUsernameParam,
) (CheckExistUsernameResult, error) {
	err := param.Validate()
	if err != nil {
		return CheckExistUsernameResult{}, err
	}
	account, err := c.accountRepo.GetAccountByUsername(ctx, param.Username)
	if err != nil {
		return CheckExistUsernameResult{}, err
	}
	return CheckExistUsernameResult{
		AlreadyExists: account != nil,
	}, nil
}

func (p *CheckExistUsernameParam) Validate() error {
	if p.Username == "" {
		return errors.NewBadRequestError("invalid params")
	}
	return nil
}

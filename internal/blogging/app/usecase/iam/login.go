package iam

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/account"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/errors"
	util_password "kang-blogging/internal/common/utils/password"
)

type LoginParams struct {
	Username string
	Password string
}

type LoginResult struct {
}

type LoginHandler decorator.UsecaseHandler[LoginParams, LoginResult]

type loginhandler struct {
	accountRepo account.Repository
}

func NewLoginHanle(
	accountRepo account.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) LoginHandler {
	if accountRepo == nil {
		panic("accountRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[LoginParams, LoginResult](
		loginhandler{
			accountRepo: accountRepo,
		},
		logger,
		metricsClient,
	)
}

func (l loginhandler) Handle(ctx context.Context, param LoginParams) (LoginResult, error) {
	err := param.Validate()
	if err != nil {
		return LoginResult{}, err
	}
	// Handler
	acc, err := l.accountRepo.GetAccountByUsername(ctx, param.Username)
	if err != nil {
		return LoginResult{}, err
	}
	if acc == nil || !util_password.CheckPasswordHash(param.Password, acc.Password) {
		return LoginResult{}, errors.NewBadRequestError("invalid username or password")
	}

	return LoginResult{}, nil
}

func (p *LoginParams) Validate() error {
	if p.Username == "" || p.Password == "" {
		return errors.NewBadRequestError("Invalid Username or Password")
	}
	return nil
}

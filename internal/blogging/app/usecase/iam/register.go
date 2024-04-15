package iam

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/account"
	"kang-blogging/internal/blogging/domain/user"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/utils"
)

type RegisterParams struct {
	Username     string
	Password     string
	Name         string
	DisplayName  string
	Email        string
	Gender       *bool
	PhoneNumbers *string
	BirthOfDay   *int64
	Avatar       *string
}

type RegisterResult struct {
	Name string
}

type RegisterHandler decorator.UsecaseHandler[RegisterParams, RegisterResult]

type registerHandler struct {
	userRepo    user.Repository
	accountRepo account.Repository
}

func NewRegisterHandler(
	userRepo user.Repository,
	accountRepo account.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) RegisterHandler {
	if userRepo == nil {
		panic("userRepo is nil")
	}
	if accountRepo == nil {
		panic("accountRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[RegisterParams, RegisterResult](
		registerHandler{
			userRepo:    userRepo,
			accountRepo: accountRepo,
		},
		logger,
		metricsClient)
}

func (r registerHandler) Handle(ctx context.Context, param RegisterParams) (RegisterResult, error) {
	err := param.Validate()
	if err != nil {
		return RegisterResult{}, err
	}
	id := utils.GenUUID()
	_, err = r.accountRepo.InsertAccount(ctx, id, param.Username, param.Password)
	if err != nil {
		return RegisterResult{}, err
	}
	return RegisterResult{}, nil
}

func (p *RegisterParams) Validate() error {
	if p.Username == "" || p.Password == "" || p.Name == "" {
		return errors.NewBadRequestError("invalid username or password")
	}
	return nil
}

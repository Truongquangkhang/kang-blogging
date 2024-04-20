package iam

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/account"
	"kang-blogging/internal/blogging/domain/role"
	"kang-blogging/internal/blogging/domain/user"
	"kang-blogging/internal/common/constants"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/model"
	"kang-blogging/internal/common/utils"
	util_password "kang-blogging/internal/common/utils/password"
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
	roleRepo    role.Repository
}

func NewRegisterHandler(
	userRepo user.Repository,
	accountRepo account.Repository,
	roleRepo role.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) RegisterHandler {
	if userRepo == nil {
		panic("userRepo is nil")
	}
	if accountRepo == nil {
		panic("accountRepo is nil")
	}
	if roleRepo == nil {
		panic("roleRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[RegisterParams, RegisterResult](
		registerHandler{
			userRepo:    userRepo,
			accountRepo: accountRepo,
			roleRepo:    roleRepo,
		},
		logger,
		metricsClient)
}

func (r registerHandler) Handle(ctx context.Context, param RegisterParams) (RegisterResult, error) {
	err := param.Validate()
	if err != nil {
		return RegisterResult{}, err
	}
	idAccount := utils.GenUUID()
	password, err := util_password.HashPassword(param.Password)
	if err != nil {
		return RegisterResult{}, err
	}
	mapNameToRole, err := r.roleRepo.GetMapNameToRole(ctx, []string{constants.USER_ROLE})
	if err != nil || len(mapNameToRole) == 0 {
		logrus.Error("Catch an error when get roles", err)
		return RegisterResult{}, err
	}

	_, err = r.accountRepo.InsertAccount(ctx, idAccount, param.Username, password)
	if err != nil {
		return RegisterResult{}, err
	}

	user := model.User{
		ID:          idAccount,
		RoleID:      mapNameToRole[constants.USER_ROLE].ID,
		Name:        param.Name,
		Email:       param.Email,
		PhoneNumber: param.PhoneNumbers,
		DisplayName: param.DisplayName,
		Avatar:      param.Avatar,
		Gender:      param.Gender,
		BirthOfDay:  param.BirthOfDay,
	}

	_, err = r.userRepo.InsertUser(ctx, &user)
	if err != nil {
		return RegisterResult{}, err
	}
	return RegisterResult{}, nil
}

func (p *RegisterParams) Validate() error {
	if p.Username == "" || p.Password == "" || p.Name == "" ||
		p.DisplayName == "" || p.Email == "" {
		return errors.NewBadRequestError("invalid params")
	}
	return nil
}

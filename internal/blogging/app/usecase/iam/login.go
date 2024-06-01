package iam

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/account"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/jwt"
	"kang-blogging/internal/common/model"
	util_password "kang-blogging/internal/common/utils/password"
	"os"
	"strconv"
)

type LoginParams struct {
	Username string
	Password string
}

type LoginResult struct {
	AccessToken  string
	RefreshToken string
	UserInfo     model.User
}

type LoginHandler decorator.UsecaseHandler[LoginParams, LoginResult]

type loginHandler struct {
	accountRepo account.Repository
}

func NewLoginHandler(
	accountRepo account.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) LoginHandler {
	if accountRepo == nil {
		panic("accountRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[LoginParams, LoginResult](
		loginHandler{
			accountRepo: accountRepo,
		},
		logger,
		metricsClient,
	)
}

func (l loginHandler) Handle(ctx context.Context, param LoginParams) (LoginResult, error) {
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

	role, err := l.accountRepo.GetRoleUserByID(ctx, acc.ID)
	if err != nil {
		return LoginResult{}, err
	}

	secretKey := os.Getenv("JWT_SECRET_KEY")
	expireHoursAccessToken, _ := strconv.Atoi(os.Getenv("JWT_EXPIRE_HOURS_ACCESS_TOKEN"))
	expireHoursRefreshTokenm, _ := strconv.Atoi(os.Getenv("JWT_EXPIRE_HOURS_REFRESH_TOKEN"))

	accessToken, err := jwt.CreateAccessToken(acc.ID, role.Name, secretKey, expireHoursAccessToken)
	if err != nil {
		return LoginResult{}, err
	}
	refreshToken, err := jwt.CreateRefreshToken(acc.ID, role.Name, secretKey, expireHoursRefreshTokenm)
	if err != nil {
		return LoginResult{}, err
	}

	return LoginResult{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserInfo:     acc.User,
	}, nil
}

func (p *LoginParams) Validate() error {
	if p.Username == "" || p.Password == "" {
		return errors.NewBadRequestError("Invalid Username or Password")
	}
	return nil
}

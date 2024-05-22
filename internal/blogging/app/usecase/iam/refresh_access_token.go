package iam

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/account"
	"kang-blogging/internal/common/constants"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/jwt"
	"os"
	"strconv"
)

type RefreshAccessTokenParams struct {
	UserID string
	Role   string
}

type RefreshAccessTokenResult struct {
	AccessToken string
}

type RefreshAccessTokenHandler decorator.UsecaseHandler[RefreshAccessTokenParams, RefreshAccessTokenResult]

type refreshAccessTokenHandler struct {
	accountRepo account.Repository
}

func NewRefreshAccessTokenHandler(
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) RefreshAccessTokenHandler {
	return decorator.ApplyUsecaseDecorators[RefreshAccessTokenParams, RefreshAccessTokenResult](
		refreshAccessTokenHandler{},
		logger,
		metricsClient,
	)
}

func (l refreshAccessTokenHandler) Handle(ctx context.Context, param RefreshAccessTokenParams) (RefreshAccessTokenResult, error) {
	err := param.Validate()
	if err != nil {
		return RefreshAccessTokenResult{}, err
	}

	secretKey := os.Getenv("JWT_SECRET_KEY")
	expireHoursAccessToken, _ := strconv.Atoi(os.Getenv("JWT_EXPIRE_HOURS_ACCESS_TOKEN"))

	accessToken, err := jwt.CreateAccessToken(param.UserID, param.Role, secretKey, expireHoursAccessToken)
	if err != nil {
		return RefreshAccessTokenResult{}, err
	}

	return RefreshAccessTokenResult{
		AccessToken: accessToken,
	}, nil
}

func (p *RefreshAccessTokenParams) Validate() error {
	if p.UserID == "" || (p.Role != constants.USER_ROLE && p.Role != constants.ADMIN_ROLE) {
		return errors.NewBadRequestError("Invalid params")
	}
	return nil
}

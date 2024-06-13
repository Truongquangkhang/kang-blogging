package user

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/domain/follow"
	"kang-blogging/internal/blogging/domain/user"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/model"
)

type UnfollowUserDetailParams struct {
	FollowedUserID string
	FollowerUserID string
}

type UnfollowUserDetailResult struct {
}

type UnfollowUserDetailHandler decorator.UsecaseHandler[UnfollowUserDetailParams, UnfollowUserDetailResult]

type unfollowUserDetailHandler struct {
	userRepo   user.Repository
	followRepo follow.Repository
}

func NewUnfollowUserDetailHandler(
	userRepo user.Repository,
	followRepo follow.Repository,
	logger *logrus.Entry,
	metrics decorator.MetricsClient,
) UnfollowUserDetailHandler {
	if userRepo == nil {
		panic("blogRepo is nil")
	}
	if followRepo == nil {
		panic("followRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[UnfollowUserDetailParams, UnfollowUserDetailResult](
		unfollowUserDetailHandler{
			userRepo:   userRepo,
			followRepo: followRepo,
		},
		logger,
		metrics,
	)
}

func (g unfollowUserDetailHandler) Handle(ctx context.Context, param UnfollowUserDetailParams) (UnfollowUserDetailResult, error) {
	err := param.Validate()
	if err != nil {
		return UnfollowUserDetailResult{}, err
	}

	_, err = g.followRepo.DeleteFollow(ctx, &model.Follow{
		FollowerID: param.FollowerUserID,
		FollowedID: param.FollowedUserID,
	})
	if err != nil {
		return UnfollowUserDetailResult{}, err
	}

	return UnfollowUserDetailResult{}, nil
}

func (p *UnfollowUserDetailParams) Validate() error {
	if p.FollowedUserID == "" || p.FollowerUserID == "" {
		return errors.NewBadRequestDefaultError()
	}
	return nil
}

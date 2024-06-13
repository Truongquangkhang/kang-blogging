package user

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/app/usecase/user"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/jwt"
)

func (g GrpcService) FollowUserDetail(
	ctx context.Context,
	request *blogging.FollowUserDetailRequest,
) (*blogging.FollowUserDetailResponse, error) {
	auth, err := jwt.GetAuthenticationFromRequest(ctx)
	if err != nil || auth == nil {
		return nil, infra.ParseGrpcError(err)
	}

	params := user.FollowUserDetailParams{
		FollowerUserID: auth.UserID,
		FollowedUserID: request.UserId,
	}
	_, err = g.usecase.FollowUserDetail.Handle(ctx, params)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}

	return &blogging.FollowUserDetailResponse{
		Code:    0,
		Message: "Success",
	}, nil
}

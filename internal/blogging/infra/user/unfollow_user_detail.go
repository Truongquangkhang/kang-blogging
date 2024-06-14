package user

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/app/usecase/user"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/jwt"
)

func (g GrpcService) UnfollowUserDetail(
	ctx context.Context,
	request *blogging.UnfollowUserDetailRequest,
) (*blogging.UnfollowUserDetailResponse, error) {
	auth, err := jwt.GetAuthenticationFromRequest(ctx)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}
	params := user.UnfollowUserDetailParams{
		FollowerUserID: auth.UserID,
		FollowedUserID: request.UserId,
	}
	_, err = g.usecase.UnfollowUserDetail.Handle(ctx, params)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}

	return &blogging.UnfollowUserDetailResponse{
		Code:    0,
		Message: "Success",
	}, nil
}

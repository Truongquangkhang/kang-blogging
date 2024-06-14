package user

import (
	"context"
	"kang-blogging/internal/blogging/app/usecase/user"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/common"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/constants"
	"kang-blogging/internal/common/jwt"
)

func (g GrpcService) GetUserDetail(
	ctx context.Context,
	request *blogging.GetUserDetailRequest,
) (*blogging.GetUserDetailResponse, error) {
	params := user.GetUserDetailParams{
		ID:          request.UserId,
		HasFullData: false,
	}

	// check user request
	auth, err := jwt.GetAuthenticationFromRequest(ctx)
	if err == nil && auth != nil {
		if auth.Role == constants.ADMIN_ROLE || auth.UserID == request.UserId {
			params.HasFullData = true
		}
		params.CurrentUserID = &auth.UserID
	}

	rs, err := g.usecase.GetUserDetail.Handle(ctx, params)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}

	return &blogging.GetUserDetailResponse{
		Code:    0,
		Message: "Success",
		Data: &blogging.GetUserDetailResponse_Data{
			User:     common.MapToUserInfoResponse(rs.User),
			Blogs:    common.MapToListBlogMetadataResponse(rs.User.Blogs),
			Comments: common.MapToCommentMetadatasResponse(rs.User.Comments),
		},
	}, nil
}

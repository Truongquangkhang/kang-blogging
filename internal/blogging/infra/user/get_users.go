package user

import (
	"context"
	"kang-blogging/internal/blogging/app/usecase/user"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/common"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/jwt"
	"kang-blogging/internal/common/utils"
)

func (g GrpcService) GetUsers(
	ctx context.Context,
	request *blogging.GetUsersRequest,
) (*blogging.GetUsersResponse, error) {
	params := user.GetUsersParams{
		Page:       request.Page,
		PageSize:   request.PageSize,
		SearchName: utils.WrapperValueString(request.SearchName),
		SearchBy:   utils.WrapperValueString(request.SearchBy),
		IsActive:   utils.WrapperValueBool(request.IsActive),
		SortBy:     utils.WrapperValueString(request.SortBy),
		Follower:   utils.WrapperValueBool(request.Follower),
		Followed:   utils.WrapperValueBool(request.Followed),
	}
	// ignore error when get auth from request
	auth, err := jwt.GetAuthenticationFromRequest(ctx)
	if err == nil && auth != nil {
		params.CurrentUserID = utils.ToStringPointerValue(auth.UserID)
	}

	rs, err := g.usecase.GetUsers.Handle(ctx, params)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}

	return &blogging.GetUsersResponse{
		Code:    0,
		Message: "Success",
		Data: &blogging.GetUsersResponse_Data{
			Users:      common.MapToUsersInfoResponse(rs.Users),
			Pagination: common.MapToPaginationResponse(rs.Pagination),
		},
	}, nil
}

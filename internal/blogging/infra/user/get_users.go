package user

import (
	"context"
	"kang-blogging/internal/blogging/app/usecase/user"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/model"
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
	}

	rs, err := g.usecase.GetUsers.Handle(ctx, params)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}

	var usersMetadata []*blogging.UserInfoMetadata
	for _, u := range rs.Users {
		usersMetadata = append(usersMetadata, mapUserToUserInfoMetadataResponse(u))
	}

	return &blogging.GetUsersResponse{
		Code:    0,
		Message: "Success",
		Data: &blogging.GetUsersResponse_Data{
			Users: usersMetadata,
			Pagination: &blogging.Pagination{
				Page:     rs.Pagination.Page,
				PageSize: rs.Pagination.PageSize,
				Total:    rs.Pagination.Total,
			},
		},
	}, nil
}

func mapUserToUserInfoMetadataResponse(u model.User) *blogging.UserInfoMetadata {
	return &blogging.UserInfoMetadata{
		Id:          u.ID,
		Name:        u.Name,
		DisplayName: u.DisplayName,
		Email:       u.Email,
		TotalBlogs:  0,
		Avatar:      utils.WrapperStringFromString(u.Avatar),
		Gender:      utils.WrapperBoolFromBool(u.Gender),
	}
}

package common

import (
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/model"
	"kang-blogging/internal/common/utils"
)

func MapUserToUserInfoMetadataResponse(u model.User) *blogging.UserInfoMetadata {
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

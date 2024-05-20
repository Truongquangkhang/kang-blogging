package common

import (
	category2 "kang-blogging/internal/blogging/domain/category"
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

func MapToCategoriesMetadata(categories []category2.ResultGetCategories) []*blogging.CategoryMetadata {
	var response []*blogging.CategoryMetadata
	for _, cate := range categories {
		response = append(response, &blogging.CategoryMetadata{
			Id:        cate.ID,
			Name:      cate.Name,
			BlogCount: cate.BlogCount,
		})
	}
	return response
}

func MapModelCommentToResponse(comment model.Comment) *blogging.Comment {
	return &blogging.Comment{
		Id:         comment.ID,
		Content:    comment.Content,
		IsToxicity: comment.IsToxicity,
		CreatedAt:  comment.CreatedAt.Unix(),
		UpdatedAt:  comment.UpdatedAt.Unix(),
		User:       MapUserToUserInfoMetadataResponse(comment.User),
	}
}

func MapToPaginationResponse(pagination model.Pagination) *blogging.Pagination {
	return &blogging.Pagination{
		Page:     pagination.Page,
		PageSize: pagination.PageSize,
		Total:    pagination.Total,
	}
}

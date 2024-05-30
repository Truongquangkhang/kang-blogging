package common

import (
	category2 "kang-blogging/internal/blogging/domain/category"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/model"
	"kang-blogging/internal/common/utils"
)

func MapUserToUserInfoMetadataResponse(u model.User) *blogging.UserInfoMetadata {
	return &blogging.UserInfoMetadata{
		Id:            u.ID,
		Name:          u.Name,
		DisplayName:   u.DisplayName,
		TotalBlogs:    u.TotalBlogs,
		Avatar:        utils.WrapperStringFromString(u.Avatar),
		TotalComments: utils.WrapperInt32FromInt32(u.TotalComments),
		Description:   utils.WrapperStringFromString(u.Description),
	}
}

func MapToUserInfoResponse(u model.User) *blogging.UserInfo {
	return &blogging.UserInfo{
		UserInfo:    MapUserToUserInfoMetadataResponse(u),
		Email:       u.Email,
		Gender:      utils.WrapperBoolFromBool(u.Gender),
		DateOfBirth: utils.WrapperInt64FromInt64(u.BirthOfDay),
		CreatedAt:   u.CreatedAt.Unix(),
		Blogs:       MapToListBlogMetadataResponse(u.Blogs),
	}
}

func MapToListBlogMetadataResponse(blogs []model.Blog) []*blogging.BlogMetadata {
	var result []*blogging.BlogMetadata
	for _, b := range blogs {
		result = append(result, MapBlogModelToBlogMetadataResponse(&b))
	}
	return result
}

func MapBlogModelToBlogMetadataResponse(b *model.Blog) *blogging.BlogMetadata {
	var categories []*blogging.Category
	for _, cat := range b.Categories {
		categories = append(categories, &blogging.Category{
			Id:   cat.ID,
			Name: cat.Name,
		})
	}
	return &blogging.BlogMetadata{
		Id:                b.ID,
		Name:              b.Title,
		Description:       *b.Summary,
		Thumbnail:         utils.WrapperStringFromString(b.Thumbnail),
		CreatedAt:         b.CreatedAt.Unix(),
		UpdatedAt:         b.UpdatedAt.Unix(),
		Categories:        categories,
		Author:            MapUserToUserInfoMetadataResponse(*b.User),
		TotalBlogComments: utils.ToInt32Value(b.TotalBlogComments),
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

func MapListModelCommentToResponse(comments []model.Comment) []*blogging.Comment {
	var result []*blogging.Comment
	for _, comment := range comments {
		result = append(result, MapModelCommentToResponse(comment))
	}
	return result
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

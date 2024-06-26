package common

import (
	"fmt"
	category2 "kang-blogging/internal/blogging/domain/category"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/model"
	"kang-blogging/internal/common/utils"
	"time"
)

func MapToUsersInfoResponse(users []model.User) []*blogging.UserInfo {
	var response []*blogging.UserInfo
	for _, user := range users {
		response = append(response, MapToUserInfoResponse(user))
	}
	return response
}

func MapToUserInfoResponse(u model.User) *blogging.UserInfo {
	return &blogging.UserInfo{
		UserInfo:        MapUserToUserInfoMetadataResponse(u),
		Email:           u.Email,
		Gender:          utils.WrapperBoolFromBool(u.Gender),
		DateOfBirth:     utils.WrapperInt64FromInt64(u.BirthOfDay),
		CreatedAt:       u.CreatedAt.Unix(),
		TotalComments:   u.TotalComments,
		TotalBlogs:      u.TotalBlogs,
		TotalFollowers:  u.TotalFollowers,
		TotalFolloweds:  u.TotalFolloweds,
		TotalViolations: u.TotalViolation,
		IsFollower:      u.IsFollower,
		IsFollowed:      u.IsFollowed,
	}
}

func MapUserToUserInfoMetadataResponse(u model.User) *blogging.UserInfoMetadata {
	var expire_warning_time *int64
	if u.ExpireWarningTime != nil {
		if utils.PointerInt64ToValue(u.ExpireWarningTime) > time.Now().Unix() {
			expire_warning_time = u.ExpireWarningTime
		}
	}
	return &blogging.UserInfoMetadata{
		Id:                u.ID,
		Name:              u.Name,
		DisplayName:       u.DisplayName,
		Avatar:            utils.WrapperStringFromString(u.Avatar),
		Description:       utils.WrapperStringFromString(u.Description),
		IsActive:          u.IsActive,
		ExpireWarningTime: utils.WrapperInt64FromInt64(expire_warning_time),
	}
}

func MapToCommentMetadatasResponse(comments []model.Comment) []*blogging.CommentMetadata {
	var results []*blogging.CommentMetadata
	for _, comment := range comments {
		results = append(results, MapToCommentMetadataResponse(&comment))
	}
	return results
}

func MapToCommentMetadataResponse(comment *model.Comment) *blogging.CommentMetadata {
	return &blogging.CommentMetadata{
		Id:             comment.ID,
		Content:        comment.Content,
		IsToxicity:     comment.IsToxicity,
		CreatedAt:      comment.CreatedAt.Unix(),
		UpdatedAt:      comment.UpdatedAt.Unix(),
		BlogId:         comment.BlogID,
		ReplyCommentId: utils.WrapperStringFromString(comment.ReplyCommentID),
	}
}

func MapToBlogInfoResponse(blog *model.Blog) *blogging.BlogInfo {
	return &blogging.BlogInfo{
		BlogInfo: MapBlogModelToBlogMetadataResponse(blog),
		Content:  utils.WrapperStringFromString(blog.Content),
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
		IsDeprecated:      b.IsDeprecated,
		Published:         b.Published,
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

func MapToCategoryResponse(category *model.Category) *blogging.Category {
	return &blogging.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: utils.WrapperStringFromString(category.Description),
	}
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
		Id:           comment.ID,
		Content:      comment.Content,
		IsToxicity:   comment.IsToxicity,
		CreatedAt:    comment.CreatedAt.Unix(),
		UpdatedAt:    comment.UpdatedAt.Unix(),
		User:         MapUserToUserInfoMetadataResponse(comment.User),
		IsDeprecated: comment.IsDeprecated,
	}
}

func MapToPaginationResponse(pagination model.Pagination) *blogging.Pagination {
	return &blogging.Pagination{
		Page:     pagination.Page,
		PageSize: pagination.PageSize,
		Total:    pagination.Total,
	}
}

func MapToPolicyResponse(policy *model.Policy) *blogging.Policy {
	return &blogging.Policy{
		Name:  utils.ToStringValue(policy.Type),
		Value: fmt.Sprint(utils.PointerInt64ToValue(policy.Value)),
	}
}

func MapToPoliciesResponse(policies []model.Policy) []*blogging.Policy {
	var result []*blogging.Policy
	for _, policy := range policies {
		result = append(result, MapToPolicyResponse(&policy))
	}
	return result
}

func MapViolationResponse(violation *model.Violation) *blogging.Violation {
	return &blogging.Violation{
		Id:          violation.ID,
		Type:        violation.ViolationType,
		TargetId:    violation.ViolationTargetID,
		Description: utils.WrapperStringFromString(violation.Description),
		CreatedAt:   violation.CreatedAt.Unix(),
		User:        MapUserToUserInfoMetadataResponse(*violation.User),
	}
}

func MapListViolationsResponse(violations []model.Violation) []*blogging.Violation {
	var result []*blogging.Violation
	for _, violation := range violations {
		result = append(result, MapViolationResponse(&violation))
	}
	return result
}

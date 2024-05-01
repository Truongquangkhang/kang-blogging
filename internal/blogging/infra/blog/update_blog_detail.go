package blog

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/app/usecase/blog"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/utils"
)

func (g GrpcService) UpdateBlogDetail(
	ctx context.Context,
	request *blogging.UpdateBlogDetailRequest,
) (*blogging.UpdateBlogDetailResponse, error) {
	authorId, _, err := infra.GetIDAndRoleFromJwtToken(ctx)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}

	param := blog.UpdateBlogDetailParams{
		AuthorID:    *authorId,
		BlogID:      request.BlogId,
		Name:        utils.WrapperValueString(request.Name),
		Description: utils.WrapperValueString(request.Description),
		Thumbnail:   utils.WrapperValueString(request.Thumbnail),
		Content:     utils.WrapperValueString(request.Content),
		CategoryIDs: request.CategoryIds,
	}

	rs, err := g.usecase.UpdateBlogDetail.Handle(ctx, param)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}

	return &blogging.UpdateBlogDetailResponse{
		Code:    0,
		Message: "Success",
		Data:    buildUpdateBlogDetailResponseData(rs),
	}, nil
}

func buildUpdateBlogDetailResponseData(rs blog.UpdateBlogDetailResult) *blogging.UpdateBlogDetailResponse_Data {
	b := rs.Blog
	var categories []*blogging.Category
	for _, category := range b.Categories {
		categories = append(categories, &blogging.Category{
			Id:   category.ID,
			Name: category.Name,
		})
	}
	return &blogging.UpdateBlogDetailResponse_Data{
		Blog: &blogging.BlogInfo{
			BlogInfo: &blogging.BlogMetadata{
				Id:          b.ID,
				Name:        b.Title,
				Description: *b.Summary,
				Thumbnail:   utils.WrapperStringFromString(b.Thumbnail),
				CreatedAt:   b.CreatedAt.Unix(),
				Categories:  categories,
				Author: &blogging.UserInfoMetadata{
					Id:          b.User.ID,
					Name:        b.User.Name,
					DisplayName: b.User.DisplayName,
					Email:       b.User.Email,
					Avatar:      utils.WrapperStringFromString(b.User.Avatar),
					Gender:      utils.WrapperBoolFromBool(b.User.Gender),
				},
			},
			Content: utils.WrapperStringFromString(b.Content),
		},
	}
}
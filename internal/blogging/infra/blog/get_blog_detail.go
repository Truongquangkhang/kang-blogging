package blog

import (
	"context"
	"kang-blogging/internal/blogging/app/usecase/blog"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/utils"
	"time"
)

func (g GrpcService) GetBlogDetail(
	ctx context.Context,
	request *blogging.GetBlogDetailRequest,
) (*blogging.GetBlogDetailResponse, error) {
	param := blog.GetBlogDetailParams{
		BlogID: request.BlogId,
	}
	rs, err := g.usecase.GetBlogDetail.Handle(ctx, param)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}

	return &blogging.GetBlogDetailResponse{
		Code:    0,
		Message: "Success",
		Data:    buildGetBlogDetailResponseData(rs),
	}, nil
}

func buildGetBlogDetailResponseData(rs blog.GetBlogDetailResult) *blogging.GetBlogDetailResponse_Data {
	b := rs.Blog
	var categories []*blogging.Category
	for _, category := range b.Categories {
		categories = append(categories, &blogging.Category{
			Id:   category.ID,
			Name: category.Name,
		})
	}
	return &blogging.GetBlogDetailResponse_Data{
		Blog: &blogging.BlogInfo{
			BlogInfo: &blogging.BlogMetadata{
				Id:          b.ID,
				Name:        b.Title,
				Description: *b.Summary,
				Thumbnail:   utils.WrapperStringFromString(b.Thumbnail),
				CreatedAt:   time.Now().Unix(),
				Categories:  categories,
				Author: &blogging.UserInfoMetadata{
					Id:          b.User.ID,
					Name:        b.User.Name,
					DisplayName: b.User.DisplayName,
					Email:       b.User.Email,
					Avatar:      utils.WrapperStringFromString(b.User.Avatar),
					Gender:      utils.WrapperBoolFromBool(b.User.Gender),
				},
				TotalBlogComments: b.TotalBlogComments,
			},
			Content: utils.WrapperStringFromString(b.Content),
		},
	}
}

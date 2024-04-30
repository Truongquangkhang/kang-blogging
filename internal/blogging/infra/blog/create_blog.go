package blog

import (
	"context"
	"kang-blogging/internal/blogging/app/usecase/blog"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/utils"
	"time"
)

func (g GrpcService) CreateBlog(
	ctx context.Context,
	request *blogging.CreateBlogRequest,
) (*blogging.CreateBlogResponse, error) {
	// validate auth token and get data
	authorId, _, err := infra.GetIDAndRoleFromJwtToken(ctx)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}

	param := blog.CreateBlogParams{
		Name:        request.Name,
		Description: request.Description,
		CategoryIds: request.CategoryIds,
		Thumbnail:   utils.WrapperValueString(request.Thumbnail),
		Content:     utils.WrapperValueString(request.Content),
		AuthorId:    *authorId,
	}
	rs, err := g.usecase.CreateBlog.Handle(ctx, param)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}

	return &blogging.CreateBlogResponse{
		Code:    0,
		Message: "Success",
		Data: &blogging.CreateBlogResponse_Data{
			Blog: &blogging.BlogInfo{
				BlogInfo: buildCreateBlogResponse(&rs),
				Content:  utils.WrapperStringFromString(rs.Blog.Content),
			},
		},
	}, nil
}

func buildCreateBlogResponse(rs *blog.CreateBlogResult) *blogging.BlogMetadata {
	b := rs.Blog
	u := rs.Author
	var categoriesResponse []*blogging.Category
	for _, category := range rs.BlogCategories {
		categoriesResponse = append(categoriesResponse, &blogging.Category{
			Id:   category.ID,
			Name: category.Name,
		})
	}

	return &blogging.BlogMetadata{
		Id:          b.ID,
		Name:        b.Title,
		Description: *b.Summary,
		Thumbnail:   utils.WrapperStringFromString(b.Thumbnail),
		Categories:  categoriesResponse,
		CreatedAt:   time.Now().Unix(),
		Author: &blogging.UserInfoMetadata{
			Id:          u.ID,
			Name:        u.Name,
			DisplayName: u.DisplayName,
			Email:       u.Email,
			Avatar:      utils.WrapperStringFromString(u.Avatar),
			Gender:      utils.WrapperBoolFromBool(u.Gender),
		},
	}
}

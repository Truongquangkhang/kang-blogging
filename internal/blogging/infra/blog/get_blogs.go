package blog

import (
	"context"
	"kang-blogging/internal/blogging/app/usecase/blog"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/utils"
	"time"
)

func (g GrpcService) GetBlogs(
	ctx context.Context,
	request *blogging.GetBlogsRequest,
) (*blogging.GetBlogsResponse, error) {
	param := blog.GetBlogsParams{
		Page:        request.Page,
		PageSize:    request.PageSize,
		SearchBy:    utils.WrapperValueString(request.SearchBy),
		SearchName:  utils.WrapperValueString(request.SearchName),
		AuthorIds:   utils.WrapperValueString(request.AuthorIds),
		CategoryIds: utils.WrapperValueString(request.CategoryIds),
	}
	rs, err := g.usecase.GetBlogs.Handle(ctx, param)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}
	return &blogging.GetBlogsResponse{
		Code:    0,
		Message: "Success",
		Data:    buildGetBlogsResponseData(rs),
	}, nil
}

func buildGetBlogsResponseData(rs blog.GetBlogsResult) *blogging.GetBlogsResponse_Data {
	var blogMetadatas []*blogging.BlogMetadata
	for _, b := range rs.Blogs {
		blogMetadatas = append(blogMetadatas, &blogging.BlogMetadata{
			Id:          b.ID,
			Name:        b.Title,
			Description: *b.Summary,
			Thumbnail:   utils.WrapperStringFromString(b.Thumbnail),
			CreatedAt:   time.Now().Unix(),
		})
	}
	pagination := &blogging.Pagination{
		Page:     rs.Pagination.Page,
		PageSize: rs.Pagination.PageSize,
		Total:    rs.Pagination.Total,
	}
	return &blogging.GetBlogsResponse_Data{
		Blogs:      blogMetadatas,
		Pagination: pagination,
	}
}

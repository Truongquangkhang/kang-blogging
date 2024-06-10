package blog

import (
	"context"
	"kang-blogging/internal/blogging/app/usecase/blog"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/common"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/jwt"
	"kang-blogging/internal/common/utils"
)

func (g GrpcService) CreateBlog(
	ctx context.Context,
	request *blogging.CreateBlogRequest,
) (*blogging.CreateBlogResponse, error) {
	// validate auth token and get data
	auth, err := jwt.GetAuthenticationFromRequest(ctx)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}

	param := blog.CreateBlogParams{
		Name:        request.Name,
		Description: request.Description,
		CategoryIds: request.CategoryIds,
		Thumbnail:   utils.WrapperValueString(request.Thumbnail),
		Content:     utils.WrapperValueString(request.Content),
		AuthorId:    auth.UserID,
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
				BlogInfo: common.MapBlogModelToBlogMetadataResponse(&rs.Blog),
				Content:  utils.WrapperStringFromString(rs.Blog.Content),
			},
		},
	}, nil
}

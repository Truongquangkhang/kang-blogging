package blog

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/app/usecase/blog"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
)

func (g GrpcService) DeleteBlogDetail(
	ctx context.Context,
	request *blogging.DeleteBlogDetailRequest,
) (*blogging.DeleteBlogDetailResponse, error) {
	authorId, _, err := infra.GetIDAndRoleFromJwtToken(ctx)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}
	param := blog.DeleteBlogDetailParams{
		DeletedById: *authorId,
		BlogID:      request.BlogId,
	}
	_, err = g.usecase.DeleteBlogDetail.Handle(ctx, param)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}
	return &blogging.DeleteBlogDetailResponse{
		Code:    0,
		Message: "Success",
	}, nil
}

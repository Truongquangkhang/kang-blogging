package blog

import (
	"context"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
)

func (g GrpcService) CreateBlog(
	ctx context.Context,
	request *blogging.CreateBlogRequest,
) (*blogging.CreateBlogResponse, error) {
	return &blogging.CreateBlogResponse{
		Code:    0,
		Message: "Success",
	}, nil
}

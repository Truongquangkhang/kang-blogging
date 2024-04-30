package blog

import (
	"context"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
)

func (g GrpcService) GetBlogs(
	ctx context.Context,
	request *blogging.GetBlogsRequest,
) (*blogging.GetBlogsResponse, error) {
	return &blogging.GetBlogsResponse{
		Code:    0,
		Message: "Success",
	}, nil
}

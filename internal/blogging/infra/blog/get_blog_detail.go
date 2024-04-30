package blog

import (
	"context"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
)

func (g GrpcService) GetBlogDetail(
	ctx context.Context,
	request *blogging.GetBlogDetailRequest,
) (*blogging.GetBlogDetailResponse, error) {
	return &blogging.GetBlogDetailResponse{
		Code:    0,
		Message: "Success",
	}, nil
}

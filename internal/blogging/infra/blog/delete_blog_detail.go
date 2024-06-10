package blog

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/app/usecase/blog"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/constants"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/jwt"
)

func (g GrpcService) DeleteBlogDetail(
	ctx context.Context,
	request *blogging.DeleteBlogDetailRequest,
) (*blogging.DeleteBlogDetailResponse, error) {
	auth, err := jwt.GetAuthenticationFromRequest(ctx)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}
	if auth.Role != constants.ADMIN_ROLE {
		return nil, infra.ParseGrpcError(errors.NewForbiddenDefaultError())
	}
	param := blog.DeleteBlogDetailParams{
		DeletedById: auth.UserID,
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

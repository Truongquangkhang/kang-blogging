package management

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/app/usecase/management"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/common"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/constants"
	"kang-blogging/internal/common/jwt"
)

func (g GrpcService) GetDashboard(
	ctx context.Context,
	request *blogging.GetDashboardRequest,
) (*blogging.GetDashboardResponse, error) {
	auth, err := jwt.GetAuthenticationFromRequest(ctx)
	if err != nil || auth == nil || auth.Role != constants.ADMIN_ROLE {
		return nil, infra.ParseGrpcError(err)
	}

	param := management.GetDashboardParams{}
	result, err := g.usecase.GetDashboard.Handle(ctx, param)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}

	return &blogging.GetDashboardResponse{
		Code:    0,
		Message: "Success",
		Data: &blogging.GetDashboardResponse_Data{
			TotalCategories:       result.SystemInfo.TotalCategories,
			TotalComments:         result.SystemInfo.TotalComments,
			TotalBlogs:            result.SystemInfo.TotalBlogs,
			TotalUsers:            result.SystemInfo.TotalUsers,
			CommentsIncreaseInDay: result.SystemInfo.CommentIncreaseInDay,
			BlogsIncreaseInDay:    result.SystemInfo.BlogIncreaseInDay,
			UsersIncreaseInDay:    result.SystemInfo.UserIncreaseInDay,
			LatestBlogs:           common.MapToListBlogMetadataResponse(result.LatestBlogs),
			LatestComments:        common.MapListModelCommentToResponse(result.LatestComments),
		},
	}, nil
}

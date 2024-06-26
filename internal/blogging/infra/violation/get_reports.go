package violation

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/app/usecase/violation"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/common"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/constants"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/jwt"
	"kang-blogging/internal/common/utils"
)

func (g GrpcService) GetReports(
	ctx context.Context,
	request *blogging.GetReportsRequest,
) (*blogging.GetReportsResponse, error) {
	auth, err := jwt.GetAuthenticationFromRequest(ctx)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}
	if auth.Role != constants.ADMIN_ROLE {
		return nil, infra.ParseGrpcError(errors.NewForbiddenDefaultError())
	}

	params := violation.GetReportsParams{
		Page:     request.Page,
		PageSize: request.PageSize,
		Type:     utils.WrapperValueString(request.Type),
		IsClosed: utils.WrapperValueBool(request.IsClosed),
		UserIDs:  utils.WrapperValueString(request.UserIds),
	}

	rs, err := g.usecase.GetReports.Handle(ctx, params)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}

	return &blogging.GetReportsResponse{
		Code:    0,
		Message: "Success",
		Data: &blogging.GetReportsResponse_Data{
			Reports:    common.MapListReportsResponse(rs.Reports),
			Pagination: common.MapToPaginationResponse(rs.Pagination),
		},
	}, nil
}

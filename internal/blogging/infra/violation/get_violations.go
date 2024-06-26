package violation

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/app/usecase/violation"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/common"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/utils"
)

func (g GrpcService) GetViolations(
	ctx context.Context,
	request *blogging.GetViolationsRequest,
) (*blogging.GetViolationsResponse, error) {
	params := violation.GetViolationParams{
		Page:     request.Page,
		PageSize: request.PageSize,
		Type:     utils.WrapperValueString(request.Type),
		UserIDs:  utils.WrapperValueString(request.UserIds),
	}
	rs, err := g.usecase.GetViolations.Handle(ctx, params)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}

	return &blogging.GetViolationsResponse{
		Code:    0,
		Message: "Success",
		Data: &blogging.GetViolationsResponse_Data{
			Violations: common.MapListViolationsResponse(rs.Violations),
			Pagination: common.MapToPaginationResponse(rs.Pagination),
		},
	}, nil
}

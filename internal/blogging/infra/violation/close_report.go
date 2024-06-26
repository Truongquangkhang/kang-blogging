package violation

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/app/usecase/violation"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/constants"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/jwt"
)

func (g GrpcService) CloseReport(
	ctx context.Context,
	request *blogging.CloseReportRequest,
) (*blogging.CloseReportResponse, error) {
	auth, err := jwt.GetAuthenticationFromRequest(ctx)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}
	if auth.Role != constants.ADMIN_ROLE {
		return nil, infra.ParseGrpcError(errors.NewForbiddenDefaultError())
	}

	params := violation.CloseReportParams{
		ReportID: request.ReportId,
	}
	_, err = g.usecase.CloseReport.Handle(ctx, params)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}

	return &blogging.CloseReportResponse{
		Code:    0,
		Message: "Success",
	}, nil
}

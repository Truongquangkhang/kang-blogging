package violation

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/app/usecase/violation"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/jwt"
	"kang-blogging/internal/common/utils"
)

func (g GrpcService) CreateReport(
	ctx context.Context,
	request *blogging.CreateReportRequest,
) (*blogging.CreateReportResponse, error) {
	auth, err := jwt.GetAuthenticationFromRequest(ctx)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}
	params := violation.CreateReportParams{
		Type:        request.Type,
		TargetID:    request.TargetId,
		Reason:      request.Reason,
		ReporterID:  auth.UserID,
		Description: utils.WrapperValueString(request.Description),
	}
	_, err = g.usecase.CreateReport.Handle(ctx, params)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}

	return &blogging.CreateReportResponse{
		Code:    0,
		Message: "Success",
	}, nil
}

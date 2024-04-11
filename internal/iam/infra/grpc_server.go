package infra

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GrpcServer struct {
	app app.Application
}

func NewGrpcServer(app app.Application) GrpcServer {
	return GrpcServer{app: app}
}

func (g GrpcServer) GetApplicableVouchers(
	ctx context.Context,
	_ *applicable_vouchers.GetApplicableVouchersRequest,
) (*applicable_vouchers.GetApplicableVouchersResponse, error) {
	// Spans are created automatically for GRPC operations
	voucherEntities, err := g.app.Queries.AllApplicableVouchers.Handle(
		ctx,
		query.AllApplicableVouchers{},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &applicable_vouchers.GetApplicableVouchersResponse{
		Code:    0,
		Message: "Success",
		Data: &applicable_vouchers.GetApplicableVouchersResponseData{
			ApplicableVouchers: voucherEntitiesToResponse(voucherEntities),
		},
	}, nil
}

func (g GrpcServer) GetApplicableVoucher(
	ctx context.Context,
	request *applicable_vouchers.GetApplicableVoucherRequest,
) (*applicable_vouchers.GetApplicableVoucherResponse, error) {
	// Spans are created automatically for GRPC operations
	voucherEntity, err := g.app.Queries.ApplicableVoucherByCode.Handle(
		ctx,
		query.ApplicableVoucherByCode{
			VoucherCode: request.VoucherCode,
		},
	)

	if err == nil {
		return &applicable_vouchers.GetApplicableVoucherResponse{
			Code:    0,
			Message: "Success",
			Data: &applicable_vouchers.GetApplicableVoucherResponseData{
				ApplicableVoucher: voucherEntityToResponse(voucherEntity),
			},
		}, nil
	}

	if baseError, ok := err.(errors.BaseError); ok {
		if baseError.ErrorCode() == errors.ERRCODE_GET_APPLICABLE_VOUCHER_BY_VOUCHER_CODE_BAD_REQUEST {
			return nil, status.Error(codes.InvalidArgument, baseError.ErrorMessage())
		} else if baseError.ErrorCode() == errors.ERRCODE_GET_APPLICABLE_VOUCHER_BY_VOUCHER_CODE_NOT_FOUND {
			return nil, status.Error(codes.NotFound, baseError.ErrorMessage())
		}
	}
	return nil, status.Error(codes.Internal, err.Error())
}

func voucherEntitiesToResponse(entities []voucher.Voucher) []*applicable_vouchers.ApplicableVoucher {
	var applicableVouchers []*applicable_vouchers.ApplicableVoucher
	for i := range entities {
		if applicableVoucher := voucherEntityToResponse(&entities[i]); applicableVoucher != nil {
			applicableVouchers = append(applicableVouchers, applicableVoucher)
		}
	}

	return applicableVouchers
}

func voucherEntityToResponse(entity *voucher.Voucher) *applicable_vouchers.ApplicableVoucher {
	if entity == nil {
		return nil
	}

	var startTime *timestamppb.Timestamp = nil
	if entity.StartTime() != nil {
		startTime = timestamppb.New(*entity.StartTime())
	}
	var endTime *timestamppb.Timestamp = nil
	if entity.EndTime() != nil {
		endTime = timestamppb.New(*entity.EndTime())
	}

	return &applicable_vouchers.ApplicableVoucher{
		VoucherCode:   entity.VoucherCode(),
		VoucherSource: entity.VoucherSource(),
		VoucherType:   entity.VoucherType(),
		Summary:       entity.Summary(),
		StartTime:     startTime,
		EndTime:       endTime,
	}
}

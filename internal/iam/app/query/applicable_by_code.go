package query

import (
	"context"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/logs"
	voucher "kang-blogging/internal/iam/domain"

	"github.com/sirupsen/logrus"
)

type ApplicableVoucherByCode struct {
	VoucherCode string
}

type ApplicableVoucherByCodeHandler decorator.QueryHandler[ApplicableVoucherByCode, *voucher.Voucher]

type applicableVoucherByCodeHandler struct {
	voucherRepo voucher.Repository
}

func NewApplicableVoucherByCodeHandler(
	voucherRepo voucher.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) ApplicableVoucherByCodeHandler {
	if voucherRepo == nil {
		panic("nil voucherRepo")
	}

	return decorator.ApplyQueryDecorators[ApplicableVoucherByCode, *voucher.Voucher](
		applicableVoucherByCodeHandler{
			voucherRepo: voucherRepo,
		},
		logger,
		metricsClient,
	)
}

func (h applicableVoucherByCodeHandler) Handle(
	ctx context.Context,
	query ApplicableVoucherByCode,
) (r *voucher.Voucher, err error) {
	defer func() {
		logs.LogQueryExecution("ApplicableVoucherByCodeHandler", query, err)
	}()

	return h.voucherRepo.GetActiveVoucherByVoucherCode(ctx, query.VoucherCode)
}

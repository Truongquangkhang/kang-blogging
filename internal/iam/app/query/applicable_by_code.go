package query

import (
	"context"
	"kang-edu/common/decorator"
	"kang-edu/common/logs"
	voucher "kang-edu/iam/domain"

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

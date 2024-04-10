package query

import (
	"context"
	"kang-edu/common/decorator"
	"kang-edu/common/logs"
	voucher "kang-edu/iam/domain"

	"github.com/sirupsen/logrus"
)

type AllApplicableVouchers struct {
}

type AllApplicableVouchersHandler decorator.QueryHandler[AllApplicableVouchers, []voucher.Voucher]

type allApplicableVouchersHandler struct {
	voucherRepo voucher.Repository
}

func NewAllApplicableVouchersHandler(
	voucherRepo voucher.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) AllApplicableVouchersHandler {
	if voucherRepo == nil {
		panic("nil voucherRepo")
	}

	return decorator.ApplyQueryDecorators[AllApplicableVouchers, []voucher.Voucher](
		allApplicableVouchersHandler{
			voucherRepo: voucherRepo,
		},
		logger,
		metricsClient,
	)
}

func (h allApplicableVouchersHandler) Handle(
	ctx context.Context,
	query AllApplicableVouchers,
) (r []voucher.Voucher, err error) {
	defer func() {
		logs.LogQueryExecution("AllApplicableVouchersHandler", query, err)
	}()

	return h.voucherRepo.GetAllActiveVouchers(ctx)
}

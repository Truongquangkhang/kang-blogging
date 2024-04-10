package voucher

import (
	"context"
)

type Repository interface {
	GetAllActiveVouchers(
		ctx context.Context,
	) ([]Voucher, error)

	GetActiveVoucherByVoucherCode(
		ctx context.Context,
		voucherCode string,
	) (*Voucher, error)
}

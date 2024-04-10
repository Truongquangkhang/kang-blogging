package voucher

import (
	"time"

	"github.com/pkg/errors"
)

type Voucher struct {
	voucherCode   string
	voucherSource string
	voucherType   string
	summary       string
	startTime     *time.Time
	endTime       *time.Time
}

func (t Voucher) VoucherCode() string {
	return t.voucherCode
}

func (t Voucher) VoucherSource() string {
	return t.voucherSource
}

func (t Voucher) VoucherType() string {
	return t.voucherType
}

func (t Voucher) Summary() string {
	return t.summary
}

func (t Voucher) StartTime() *time.Time {
	return t.startTime
}

func (t Voucher) EndTime() *time.Time {
	return t.endTime
}

func NewVoucher(
	voucherCode string,
	voucherSource string,
	voucherType string,
	summary string,
	startTime *time.Time,
	endTime *time.Time,
) (*Voucher, error) {
	if voucherCode == "" {
		return nil, errors.New("empty voucher code")
	}
	if voucherSource == "" {
		return nil, errors.New("empty voucher source")
	}
	if voucherType == "" {
		return nil, errors.New("empty voucher type")
	}
	if summary == "" {
		return nil, errors.New("empty voucher summary")
	}

	return &Voucher{
		voucherCode:   voucherCode,
		voucherSource: voucherSource,
		voucherType:   voucherType,
		summary:       summary,
		startTime:     startTime,
		endTime:       endTime,
	}, nil
}

// UnmarshalVoucherFromDatabase unmarshals Voucher from the database.
//
// It should be used only for unmarshalling from the database!
// You can't use UnmarshalTrainingFromDatabase as constructor - It may put domain into the invalid state!
func UnmarshalVoucherFromDatabase(
	voucherCode string,
	voucherSource string,
	voucherType string,
	summary string,
	startTime *time.Time,
	endTime *time.Time,
) (*Voucher, error) {
	voucher, err := NewVoucher(
		voucherCode,
		voucherSource,
		voucherType,
		summary,
		startTime,
		endTime,
	)
	if err != nil {
		return nil, err
	}

	return voucher, nil
}

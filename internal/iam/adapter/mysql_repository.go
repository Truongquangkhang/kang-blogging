package adapters

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
)

type mysqlVoucher struct {
	ID            int64      `db:"id"`
	VoucherCode   string     `db:"voucher_code"`
	VoucherSource string     `db:"voucher_source"`
	VoucherType   string     `db:"voucher_type"`
	Summary       string     `db:"summary"`
	StartTime     *time.Time `db:"start_time"`
	EndTime       *time.Time `db:"end_time"`
}

const voucherMysqlName = "VoucherMysqlRepository"

var mysql_voucher_repository_logger *logrus.Entry = logrus.StandardLogger().WithField("logger", voucherMysqlName)

type MySQLVoucherRepository struct {
	db *sqlx.DB
}

func NewMySQLVoucherRepository(db *sqlx.DB) *MySQLVoucherRepository {
	if db == nil {
		panic("missing db")
	}

	return &MySQLVoucherRepository{db: db}
}

func (r MySQLVoucherRepository) GetAllActiveVouchers(
	ctx context.Context,
) ([]voucher.Voucher, error) {
	_, span := tracing.StartSpan(ctx)
	defer span.End()

	results := []voucher.Voucher{}

	vouchersFromDb := []mysqlVoucher{}
	if err := r.db.Unsafe().Select(
		&vouchersFromDb,
		`SELECT *
		FROM vouchers
		WHERE NOW() BETWEEN start_time AND end_time AND deleted = false;`,
	); err != nil {
		mysql_voucher_repository_logger.WithField("err", err).
			Error("error while query vouchers from db and scan into structs")
		tracing.UpdateSpanError(span, err)
		return nil, err
	}

	for _, voucherFromDb := range vouchersFromDb {
		newVoucher, err := voucher.UnmarshalVoucherFromDatabase(
			voucherFromDb.VoucherCode,
			voucherFromDb.VoucherSource,
			voucherFromDb.VoucherType,
			voucherFromDb.Summary,
			voucherFromDb.StartTime,
			voucherFromDb.EndTime,
		)
		if err != nil {
			mysql_voucher_repository_logger.WithField("err", err).Error("error while creating new voucher from db data")
			continue
		}
		results = append(results, *newVoucher)
	}

	return results, nil
}

func (r MySQLVoucherRepository) GetActiveVoucherByVoucherCode(
	ctx context.Context,
	voucherCode string,
) (*voucher.Voucher, error) {
	_, span := tracing.StartSpan(ctx)
	defer span.End()
	span.SetAttributes(attribute.String("voucherCode", voucherCode))

	var voucherFromDb mysqlVoucher
	// There might be many vouchers with the same voucher_code
	// These are considered same value, we will prioritize voucher
	// with expire date as late as possible
	if err := r.db.Unsafe().Get(
		&voucherFromDb,
		`SELECT *
		FROM vouchers
		WHERE deleted = false AND voucher_code=(?)
		ORDER BY end_time DESC LIMIT 1;`,
		voucherCode,
	); err != nil {
		switch {
		case err == sql.ErrNoRows:
			mysql_voucher_repository_logger.WithField("err", err).
				Info(fmt.Sprintf("not found active voucher with code '%s'", voucherCode))
			tracing.UpdateSpanError(span, err)
			return nil, errors.NewGetApplicableVoucherByVoucherCodeNotFoundDefaultError()
		default:
			mysql_voucher_repository_logger.WithField("err", err).
				Error("error while querying db and scan into structs")
			tracing.UpdateSpanError(span, err)
			return nil, err
		}
	}

	svNow := utils.GetServerNow()
	span.SetAttributes(attribute.String("Server Now Timestamp", svNow.String()))
	if svNow.Before(*voucherFromDb.StartTime) || svNow.After(*voucherFromDb.EndTime) {
		return nil, errors.NewGetApplicableVoucherByVoucherCodeBadRequestDefaultError()
	}

	voucher, err := voucher.UnmarshalVoucherFromDatabase(
		voucherFromDb.VoucherCode,
		voucherFromDb.VoucherSource,
		voucherFromDb.VoucherType,
		voucherFromDb.Summary,
		voucherFromDb.StartTime,
		voucherFromDb.EndTime,
	)
	if err != nil {
		mysql_voucher_repository_logger.WithField("err", err).Error("error while creating new voucher from db data")
		tracing.UpdateSpanError(span, err)
		return nil, err
	}
	return voucher, nil
}

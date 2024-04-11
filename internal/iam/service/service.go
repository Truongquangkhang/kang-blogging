package service

import (
	"context"
	"kang-blogging/common/db"
	metrics "kang-blogging/common/metric"
	adapters "kang-blogging/iam/adapter"
	"kang-blogging/iam/app"
	"kang-blogging/iam/app/command"
	"kang-blogging/iam/app/query"
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

func NewApplication(ctx context.Context) (app.Application, func()) {
	return newApplication(ctx),
		func() {}
}

func newApplication(ctx context.Context) app.Application {
	logrus.Info(ctx)

	connCount, _ := strconv.Atoi(os.Getenv("DB_CONN_COUNT"))
	connIdleTimeSec, _ := strconv.Atoi(os.Getenv("DB_CONN_IDLE_TIME_SEC"))
	connLifeTimeSec, _ := strconv.Atoi(os.Getenv("DB_CONN_LIFE_TIME_SEC"))

	dbConfig := db.MysqlConfig{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
		Name: os.Getenv("DB_NAME"),

		MaxOpenConns:       connCount,
		MaxIdleConns:       connCount,
		ConnMaxIdleTimeSec: time.Duration(connIdleTimeSec),
		ConnMaxLifetimeSec: time.Duration(connLifeTimeSec),
	}

	mysqlDb, err := db.NewMySQLConnection(dbConfig)
	if err != nil {
		panic(err)
	}
	voucherRepository := adapters.NewMySQLVoucherRepository(mysqlDb)

	logger := logrus.NewEntry(logrus.StandardLogger())

	metricsClient := metrics.NoOp{}

	return app.Application{
		Commands: app.Commands{
			DoSomething: command.NewDoSomethingHandler(voucherRepository, logger, metricsClient),
		},
		Queries: app.Queries{
			AllApplicableVouchers:   query.NewAllApplicableVouchersHandler(voucherRepository, logger, metricsClient),
			ApplicableVoucherByCode: query.NewApplicableVoucherByCodeHandler(voucherRepository, logger, metricsClient),
		},
	}
}

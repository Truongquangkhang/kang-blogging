package service

import (
	"context"
	"github.com/sirupsen/logrus"
	adapters "kang-blogging/internal/blogging/adapter"
	"kang-blogging/internal/blogging/app"
	"kang-blogging/internal/blogging/app/command"
	"kang-blogging/internal/common/db"
	metrics "kang-blogging/internal/common/metric"
	"os"
	"strconv"
	"time"
)

func NewApplication(ctx context.Context) (app.Application, func()) {
	return newService(ctx), func() {

	}
}

func newService(ctx context.Context) app.Application {
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

	repository := adapters.NewMySQLVoucherRepository(mysqlDb)
	logger := logrus.NewEntry(logrus.StandardLogger())
	metricsClient := metrics.NoOp{}

	return app.Application{
		Command: app.Command{
			DoSomething: command.NewDoSomethingHandler(repository, logger, metricsClient),
		},
		//Queries: app.Queries{
		//	AllApplicableVouchers:   query.NewAllApplicableVouchersHandler(repository, logger, metricsClient),
		//	ApplicableVoucherByCode: query.NewApplicableVoucherByCodeHandler(repository, logger, metricsClient),
		//},
	}
}

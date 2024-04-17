package service

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/adapter/account"
	"kang-blogging/internal/blogging/adapter/role"
	"kang-blogging/internal/blogging/adapter/user"
	"kang-blogging/internal/blogging/app"
	"kang-blogging/internal/blogging/app/usecase/iam"
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
	var gdb = db.GetDBInstance()
	if err = gdb.Open(dbConfig); err != nil {
		panic(err)
	}
	fmt.Printf("db: %v\n", mysqlDb)

	userRepository := user.NewRepository()
	accountRepository := account.NewRepository()
	roleRepository := role.NewRepository()
	logger := logrus.NewEntry(logrus.StandardLogger())
	metricsClient := metrics.NoOp{}

	return app.Application{
		IAMUsecases: app.IAMUsecases{
			Login: iam.NewLoginHanle(
				accountRepository, logger, metricsClient,
			),
			Register: iam.NewRegisterHandler(
				userRepository, accountRepository, roleRepository, logger, metricsClient,
			),
			CheckExistUsername: iam.NewCheckExistUsernameHandler(accountRepository, logger, metricsClient),
		},
	}
}

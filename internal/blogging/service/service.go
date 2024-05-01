package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/adapter/account"
	blog2 "kang-blogging/internal/blogging/adapter/blog"
	"kang-blogging/internal/blogging/adapter/blog_categories"
	"kang-blogging/internal/blogging/adapter/category"
	"kang-blogging/internal/blogging/adapter/role"
	"kang-blogging/internal/blogging/adapter/user"
	"kang-blogging/internal/blogging/app"
	"kang-blogging/internal/blogging/app/usecase/blog"
	"kang-blogging/internal/blogging/app/usecase/iam"
	user2 "kang-blogging/internal/blogging/app/usecase/user"
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

	var gdb = db.GetDBInstance()
	if err := gdb.Open(dbConfig); err != nil {
		panic(err)
	}

	userRepository := user.NewRepository()
	accountRepository := account.NewRepository()
	roleRepository := role.NewRepository()
	blogRepository := blog2.NewRepository()
	categoryRepository := category.NewRepository()
	blogCategoriesRepository := blog_categories.NewRepository()
	logger := logrus.NewEntry(logrus.StandardLogger())
	metricsClient := metrics.NoOp{}

	return app.Application{
		IAMUsecases: app.IAMUsecases{
			Login: iam.NewLoginHandler(
				accountRepository, logger, metricsClient,
			),
			Register: iam.NewRegisterHandler(
				userRepository, accountRepository, roleRepository, logger, metricsClient,
			),
			CheckExistUsername: iam.NewCheckExistUsernameHandler(
				accountRepository, logger, metricsClient,
			),
		},
		UserUsecase: app.UserUsecase{
			GetUsers: user2.NewGetUserHandler(
				userRepository, logger, metricsClient,
			),
			GetUserDetail: user2.NewGetUserDetailHandler(
				userRepository, logger, metricsClient,
			),
			UpdateUser: user2.NewUpdateUserHandler(
				userRepository, logger, metricsClient,
			),
		},
		BlogUsecase: app.BlogUsecase{
			GetBlogs: blog.NewGetBlogsHandler(
				userRepository, blogRepository, logger, metricsClient,
			),
			CreateBlog: blog.NewCreateBlogHandler(
				blogRepository, categoryRepository, blogCategoriesRepository,
				userRepository, logger, metricsClient,
			),
			GetBlogDetail: blog.NewGetBlogDetailHandler(
				blogRepository, logger, metricsClient,
			),
			UpdateBlogDetail: blog.NewUpdateBlogDetailHandler(
				blogRepository, logger, metricsClient,
			),
		},
	}
}

package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/adapter/account"
	blog2 "kang-blogging/internal/blogging/adapter/blog"
	"kang-blogging/internal/blogging/adapter/blog_categories"
	"kang-blogging/internal/blogging/adapter/category"
	"kang-blogging/internal/blogging/adapter/comment"
	"kang-blogging/internal/blogging/adapter/follow"
	"kang-blogging/internal/blogging/adapter/policy"
	"kang-blogging/internal/blogging/adapter/report"
	"kang-blogging/internal/blogging/adapter/role"
	"kang-blogging/internal/blogging/adapter/toxicity_detection_client"
	"kang-blogging/internal/blogging/adapter/user"
	"kang-blogging/internal/blogging/adapter/violation"
	"kang-blogging/internal/blogging/app"
	"kang-blogging/internal/blogging/app/usecase/blog"
	category2 "kang-blogging/internal/blogging/app/usecase/category"
	comment2 "kang-blogging/internal/blogging/app/usecase/comment"
	"kang-blogging/internal/blogging/app/usecase/iam"
	"kang-blogging/internal/blogging/app/usecase/image"
	"kang-blogging/internal/blogging/app/usecase/management"
	user2 "kang-blogging/internal/blogging/app/usecase/user"
	violation2 "kang-blogging/internal/blogging/app/usecase/violation"
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

	// repo
	userRepository := user.NewRepository()
	accountRepository := account.NewRepository()
	roleRepository := role.NewRepository()
	blogRepository := blog2.NewRepository()
	categoryRepository := category.NewRepository()
	blogCategoriesRepository := blog_categories.NewRepository()
	commentRepository := comment.NewRepository()
	policyRepository := policy.NewRepository()
	followRepository := follow.NewRepository()
	violationRepository := violation.NewRepository()
	reportRepository := report.NewRepository()

	logger := logrus.NewEntry(logrus.StandardLogger())
	metricsClient := metrics.NoOp{}

	// adapter client
	detectionClient := toxicity_detection_client.NewToxicityDetectionClient()

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
			RefreshAccessToken: iam.NewRefreshAccessTokenHandler(
				accountRepository, logger, metricsClient,
			),
			ChangePassword: iam.NewChangePasswordHandler(
				roleRepository, accountRepository, logger, metricsClient,
			),
		},
		UserUsecase: app.UserUsecase{
			GetUsers: user2.NewGetUserHandler(
				userRepository, logger, metricsClient,
			),
			GetUserDetail: user2.NewGetUserDetailHandler(
				userRepository, blogRepository, commentRepository, logger, metricsClient,
			),
			UpdateUser: user2.NewUpdateUserHandler(
				userRepository, logger, metricsClient,
			),
			DeleteUserDetail: user2.NewDeleteUserDetailHandler(
				userRepository, roleRepository, logger, metricsClient,
			),
			FollowUserDetail: user2.NewFollowUserDetailHandler(
				userRepository, followRepository, logger, metricsClient,
			),
			UnfollowUserDetail: user2.NewUnfollowUserDetailHandler(
				userRepository, followRepository, logger, metricsClient,
			),
		},
		BlogUsecase: app.BlogUsecase{
			GetBlogs: blog.NewGetBlogsHandler(
				userRepository, blogRepository, followRepository, logger, metricsClient,
			),
			CreateBlog: blog.NewCreateBlogHandler(
				blogRepository, categoryRepository, blogCategoriesRepository,
				userRepository, logger, metricsClient,
			),
			GetBlogDetail: blog.NewGetBlogDetailHandler(
				blogRepository, logger, metricsClient,
			),
			UpdateBlogDetail: blog.NewUpdateBlogDetailHandler(
				blogRepository, roleRepository, logger, metricsClient,
			),
			DeleteBlogDetail: blog.NewDeleteBlogDetailHandler(
				blogRepository, roleRepository, logger, metricsClient,
			),
		},
		CategoryUsecase: app.CategoryUsecase{
			GetCategories: category2.NewGetCategoriesHandler(
				categoryRepository, logger, metricsClient,
			),
			CreateCategory: category2.NewCreateCategoryHandler(
				categoryRepository, logger, metricsClient,
			),
			UpdateCategory: category2.NewUpdateCategoryHandler(
				categoryRepository, logger, metricsClient,
			),
		},
		CommentUsecase: app.CommentUsecase{
			GetBlogComments: comment2.NewGetBlogCommentsHandler(
				commentRepository, logger, metricsClient,
			),
			CreateBlogComment: comment2.NewCreateBlogCommentHandler(
				commentRepository, userRepository, detectionClient, logger, metricsClient,
			),
			GetComments: comment2.NewGetCommentsHandler(
				commentRepository, logger, metricsClient,
			),
			GetComment: comment2.NewGetCommentHandler(
				commentRepository, userRepository, logger, metricsClient,
			),
			UpdateComment: comment2.NewUpdateCommentHandler(
				commentRepository, userRepository, roleRepository, detectionClient, logger, metricsClient,
			),
			DeleteComment: comment2.NewDeleteCommentHandler(
				commentRepository, userRepository, roleRepository, logger, metricsClient,
			),
			SetCommentAsToxic: comment2.NewSetCommentAsToxicHandler(
				commentRepository, logger, metricsClient,
			),
		},
		ImageUsecase: app.ImageUsecase{
			UploadImage: image.NewUploadImageHandler(
				logger, metricsClient,
			),
		},
		ManagementUsecase: app.ManagementUsecase{
			GetDashboard: management.NewGetDashboardHandler(
				commentRepository, blogRepository, categoryRepository, userRepository, logger, metricsClient,
			),
			GetPolicies: management.NewGetPoliciesHandler(
				policyRepository, logger, metricsClient,
			),
			UpdatePolicies: management.NewUpdatePoliciesHandler(
				policyRepository, logger, metricsClient,
			),
		},
		ViolationUsecase: app.ViolationUsecase{
			GetViolations: violation2.NewGetViolationHandler(
				violationRepository, logger, metricsClient,
			),
			CreateReport: violation2.NewCreateReportHandler(
				reportRepository, commentRepository, logger, metricsClient,
			),
			GetReports: violation2.NewGetReportsHandler(
				reportRepository, logger, metricsClient,
			),
			CloseReport: violation2.NewCloseReportHandler(
				reportRepository, logger, metricsClient,
			),
		},
	}
}

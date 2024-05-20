package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/infra/blog"
	"kang-blogging/internal/blogging/infra/category"
	"kang-blogging/internal/blogging/infra/comment"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/blogging/infra/iam"
	"kang-blogging/internal/blogging/infra/user"
	"kang-blogging/internal/blogging/service"
	"kang-blogging/internal/common/logs"
	"kang-blogging/internal/common/server"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	logs.Init()
	//tracingCleanUp := tracing.Init()
	//defer tracingCleanUp()

	ctx := context.Background()

	application, appCleanUp := service.NewApplication(ctx)
	defer appCleanUp()

	server.RunGRPCServer(
		ctx,
		func(server *grpc.Server) {
			svcIAM := iam.NewGrpcService(application.IAMUsecases)
			blogging.RegisterIAMServiceServer(server, svcIAM)
			svcUser := user.NewGrpcService(application.UserUsecase)
			blogging.RegisterUserServiceServer(server, svcUser)
			svcBlog := blog.NewGrpcService(application.BlogUsecase)
			blogging.RegisterBlogServiceServer(server, svcBlog)
			svcCategory := category.NewGrpcService(application.CategoryUsecase)
			blogging.RegisterCategoryServiceServer(server, svcCategory)
			svcComment := comment.NewGrpcService(application.CommentUsecase)
			blogging.RegisterCommentServiceServer(server, svcComment)
		},
		func(mux *runtime.ServeMux, conn *grpc.ClientConn) {
			err := blogging.RegisterIAMServiceHandler(ctx, mux, conn)
			if err != nil {
				logrus.Fatal(err)
			}
			err = blogging.RegisterUserServiceHandler(ctx, mux, conn)
			if err != nil {
				logrus.Fatal(err)
			}
			err = blogging.RegisterBlogServiceHandler(ctx, mux, conn)
			if err != nil {
				logrus.Fatal(err)
			}
			err = blogging.RegisterCategoryServiceHandler(ctx, mux, conn)
			if err != nil {
				logrus.Fatal(err)
			}
			err = blogging.RegisterCommentServiceHandler(ctx, mux, conn)
			if err != nil {
				logrus.Fatal(err)
			}
		},
	)
}

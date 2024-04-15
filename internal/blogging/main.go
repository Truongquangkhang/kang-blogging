package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/blogging/infra/iam"
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
			blogging.RegisterBloggingServiceServer(server, svcIAM)
		},
		func(mux *runtime.ServeMux, conn *grpc.ClientConn) {
			err := blogging.RegisterBloggingServiceHandler(ctx, mux, conn)
			if err != nil {
				logrus.Fatal(err)
			}
		},
	)
}

package main

import (
	"context"
	"kang-edu/common/logs"
	"kang-edu/common/server"
	"kang-edu/common/tracing"
	"kang-edu/iam/infra"
	"kang-edu/iam/infra/genproto/applicable_vouchers"
	"kang-edu/iam/service"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	logs.Init()
	tracingCleanUp := tracing.Init()
	defer tracingCleanUp()

	ctx := context.Background()

	application, appCleanUp := service.NewApplication(ctx)
	defer appCleanUp()

	server.RunGRPCServer(
		ctx,
		func(server *grpc.Server) {
			svc := infra.NewGrpcServer(application)
			applicable_vouchers.RegisterApplicableVouchersServiceServer(server, svc)
		},
		func(mux *runtime.ServeMux, conn *grpc.ClientConn) {
			err := applicable_vouchers.RegisterApplicableVouchersServiceHandler(ctx, mux, conn)
			if err != nil {
				logrus.Fatal(err)
			}
		},
	)
}

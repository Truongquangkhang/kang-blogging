package main

import (
	"context"
	"fmt"
	"kang-blogging/internal/blogging/infra"
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
			svc := infra.NewGrpcServer(application)
			fmt.Println(svc)
		},
		func(mux *runtime.ServeMux, conn *grpc.ClientConn) {
			//err := applicable_vouchers.RegisterApplicableVouchersServiceHandler(ctx, mux, conn)
			//if err != nil {
			//	logrus.Fatal(err)
			//}
		},
	)
}

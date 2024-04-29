package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/blogging/infra/iam"
	"kang-blogging/internal/blogging/infra/user"
	"kang-blogging/internal/blogging/service"
	"kang-blogging/internal/common/logs"
	"kang-blogging/internal/common/server"
)

func main() {
	logs.Init()
	//tracingCleanUp := tracing.Init()
	//defer tracingCleanUp()

	ctx := context.Background()

	application, appCleanUp := service.NewApplication(ctx)
	defer appCleanUp()

	//server.RunHTTPServer(func(router chi.Router) http.Handler {
	//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//		body, _ := ioutil.ReadFile("api/openapi/blogging/blogging.yaml")
	//		fmt.Fprint(w, string(body))
	//	})
	//})

	server.RunGRPCServer(
		ctx,
		func(server *grpc.Server) {
			svcIAM := iam.NewGrpcService(application.IAMUsecases)
			blogging.RegisterIAMServiceServer(server, svcIAM)
			svcUser := user.NewGrpcService(application.UserUsecase)
			blogging.RegisterUserServiceServer(server, svcUser)
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
		},
	)
}

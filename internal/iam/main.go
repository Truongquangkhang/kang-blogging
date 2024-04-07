package main

import (
	"context"
	"fmt"
	"kang-edu/common/logs"
	"kang-edu/common/server"
	"kang-edu/common/tracing"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	logs.Init()
	tracingCleanup := tracing.Init()
	defer tracingCleanup()

	ctx := context.Background()

	// application, appCleanUp := service.NewApplication(ctx)
	// defer appCleanUp()

	server.RunHTTPServer(
		func(router chi.Router) http.Handler {
			fmt.Printf("%v", ctx)
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Println("Hello World")
			})
		},
	)

	// server.RunHTTPServer(
	// 	func(router chi.Router) http.Handler {
	// 		return infra.HandlerWithOptionsWrapper(
	// 			infra.NewHttpServer(application),
	// 			voucherhub.ChiServerOptions{
	// 				BaseRouter:       router,
	// 				ErrorHandlerFunc: httperr.ChiErrorHandler(),
	// 			},
	// 		)
	// 	},
	// )
}

package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpclogrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpcctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func init() {
	logger := logrus.New()
	//logs.SetFormatter(logger)
	logger.SetLevel(logrus.WarnLevel)

	grpclogrus.ReplaceGrpcLogger(logrus.NewEntry(logger))
	grpcprometheus.EnableHandlingTimeHistogram()
}

func RunGRPCServer(
	ctx context.Context,
	registerServer func(server *grpc.Server),
	registerHandler func(mux *runtime.ServeMux, conn *grpc.ClientConn),
) {
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "10443"
	}
	RunGRPCServerOnAddr(
		fmt.Sprintf(":%s", grpcPort),
		registerServer,
	)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "10080"
	}
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		ctx,
		fmt.Sprintf("localhost:%s", grpcPort),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logrus.Fatal(err)
	}
	ServingGRPCGatewayOnAddr(
		conn,
		fmt.Sprintf(":%s", httpPort),
		registerHandler,
	)
}

func RunGRPCServerOnAddr(
	addr string,
	registerServer func(server *grpc.Server),
) {
	logrusEntry := logrus.NewEntry(logrus.StandardLogger())

	grpcServer := grpc.NewServer(
		grpcmiddleware.WithUnaryServerChain(
			grpcctxtags.UnaryServerInterceptor(grpcctxtags.WithFieldExtractor(grpcctxtags.CodeGenRequestFieldExtractor)),
			grpclogrus.UnaryServerInterceptor(logrusEntry),
			grpcprometheus.UnaryServerInterceptor,
			otelgrpc.UnaryServerInterceptor(),
			//grpc.UnaryServerInterceptor(schemalog.UnaryServerInterceptor),
		),
		grpcmiddleware.WithStreamServerChain(
			grpcctxtags.StreamServerInterceptor(grpcctxtags.WithFieldExtractor(grpcctxtags.CodeGenRequestFieldExtractor)),
			grpclogrus.StreamServerInterceptor(logrusEntry),
			grpcprometheus.StreamServerInterceptor,
			otelgrpc.StreamServerInterceptor(),
			//grpc.StreamServerInterceptor(schemalog.StreamServerInterceptor),
		),
	)

	//grpc_health_v1.RegisterHealthServer(grpcServer, health.NewGrpcHealthChecker())

	registerServer(grpcServer)

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.WithField("grpcEndpoint", addr).Info("Starting: gRPC Listener")
	go func() {
		logrus.Fatal(grpcServer.Serve(listen))
	}()
}

func ServingGRPCGatewayOnAddr(
	conn *grpc.ClientConn,
	addr string,
	registerHandler func(mux *runtime.ServeMux, conn *grpc.ClientConn),
) {
	gwmux := runtime.NewServeMux(runtime.WithHealthzEndpoint(grpc_health_v1.NewHealthClient(conn)))
	registerHandler(gwmux, conn)
	registerPrometheusMetricsHandler(gwmux)

	server := &http.Server{
		Addr:              addr,
		Handler:           AddTracingMiddleware(AddSchemaLogMiddleware(gwmux)),
		ReadHeaderTimeout: 20 * time.Second,
		ReadTimeout:       1 * time.Minute,
		WriteTimeout:      2 * time.Minute,
	}
	logrus.WithField("grpcGatewayEndpoint", addr).Info("Starting: gRPC Gateway")
	logrus.Fatal(server.ListenAndServe())
}

func registerPrometheusMetricsHandler(gwmux *runtime.ServeMux) {
	err := gwmux.HandlePath(
		"GET", "/metrics",
		func(w http.ResponseWriter, req *http.Request, _ map[string]string) {
			promhttp.Handler().ServeHTTP(w, req)
		},
	)
	if err != nil {
		panic(fmt.Errorf("register prometheus handler for chi failed. Err=%v", err))
	}
}

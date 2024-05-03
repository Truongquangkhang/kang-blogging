package health

import (
	"context"

	"google.golang.org/grpc/health/grpc_health_v1"
)

type GrpcHealthChecker struct{}

func NewGrpcHealthChecker() *GrpcHealthChecker {
	return &GrpcHealthChecker{}
}

func (s *GrpcHealthChecker) Check(
	ctx context.Context,
	req *grpc_health_v1.HealthCheckRequest,
) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

func (s *GrpcHealthChecker) Watch(
	req *grpc_health_v1.HealthCheckRequest,
	server grpc_health_v1.Health_WatchServer,
) error {
	return server.Send(&grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	})
}

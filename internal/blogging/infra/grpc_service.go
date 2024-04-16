package infra

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"kang-blogging/internal/blogging/app"
	"kang-blogging/internal/common/errors"
	"net/http"
)

type GrpcServer struct {
	app app.Application
}

func NewGrpcServer(app app.Application) GrpcServer {
	return GrpcServer{app: app}
}

func ParseGrpcError(err error) error {
	if err == nil {
		return nil
	}
	if baseError, ok := err.(errors.BaseError); ok {
		msg := baseError.ErrorMessage()
		switch baseError.BaseErrorCode() {
		case http.StatusBadRequest:
			return status.Error(codes.InvalidArgument, msg)
		case http.StatusForbidden:
			return status.Error(codes.PermissionDenied, msg)
		case http.StatusNotFound:
			return status.Error(codes.NotFound, msg)
		case http.StatusConflict:
			return status.Error(codes.AlreadyExists, msg)
		default:
			return status.Error(codes.Internal, msg)
		}
	} else {
		return status.Error(codes.Internal, err.Error())
	}
}

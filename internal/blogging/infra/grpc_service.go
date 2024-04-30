package infra

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"kang-blogging/internal/blogging/app"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/jwt"
	"net/http"
	"os"
	"strings"
)

type GrpcServer struct {
	app app.Application
}

func NewGrpcServer(app app.Application) GrpcServer {
	return GrpcServer{app: app}
}

func GetIDAndRoleFromJwtToken(ctx context.Context) (*string, *string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, nil, fmt.Errorf("catch an error when read header")
	}
	authorizations := md.Get("authorization")
	if len(authorizations) < 1 {
		return nil, nil, fmt.Errorf("bearer token not found")
	}
	bearerToken := strings.TrimPrefix(authorizations[0], "Bearer ")

	if err := jwt.VerifyToken(bearerToken, secretKey); err != nil {
		return nil, nil, err
	}
	id, role, err := jwt.GetIDAndRoleFromJwtToken(bearerToken, secretKey)
	if err != nil {
		return nil, nil, err
	}
	return &id, &role, err
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

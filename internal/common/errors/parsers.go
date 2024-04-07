package errors

import (
	"github.com/sirupsen/logrus"
	grpcCodes "google.golang.org/grpc/codes"
	grpcStatus "google.golang.org/grpc/status"
)

func ParseGrpcError(err error) BaseError {
	// check if err is grpc error
	if st, ok := grpcStatus.FromError(err); ok {
		if st.Code() == grpcCodes.NotFound {
			return NewNotFoundError(st.Message())
		}
		if st.Code() == grpcCodes.InvalidArgument {
			return NewBadRequestError(st.Message())
		}
		// network error
		return NewInternalErrorDefaultError()
	} else {
		return NewInternalErrorDefaultError()
	}
}

func ParseGrpcCallError(logger *logrus.Entry, err error) BaseError {
	// check if err is grpc error
	if st, ok := grpcStatus.FromError(err); ok {
		logger := logger.WithField("err", err).
			WithField("grpc-code", st.Code()).
			WithField("grpc-msg", st.Message())

		if st.Code() == grpcCodes.NotFound {
			logger.Info("GRPC call failed - request not found")
			return NewNotFoundError(st.Message())
		}
		if st.Code() == grpcCodes.InvalidArgument {
			logger.Info("GRPC call failed - bad request")
			return NewBadRequestError(st.Message())
		}
		// network error
		logger.Error("GRPC call failed - unexpected")
		return NewInternalErrorDefaultError()
	} else {
		logger.WithField("err", err).
			Error("err is not type grpc status.Error()")
		return NewInternalErrorDefaultError()
	}
}

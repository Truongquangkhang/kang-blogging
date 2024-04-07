package httperr

import (
	"kang-edu/common/errors"
	"kang-edu/common/logs"
	"kang-edu/common/server/httpheader"
	"net/http"

	"github.com/go-chi/render"
)

func RespondWithError(err error, w http.ResponseWriter, r *http.Request) {
	logger := logs.GetLogEntry(r).WithError(err)
	w.Header().Set(httpheader.CONTENT_TYPE, httpheader.CONTENT_TYPE_APPLICATION_JSON)
	baseError, ok := err.(errors.BaseError)
	if !ok {
		logger.WithField("error-msg", err.Error()).Error(errors.ERRCODE_INTERNAL_ERROR)
		httpRespondWithError(
			errors.NewInternalErrorDefaultError(),
			w, r,
		)
		return
	}

	logger.WithField("error-msg", baseError.ErrorMessage()).Info(baseError.ErrorCode())
	httpRespondWithError(baseError, w, r)
}

func httpRespondWithError(
	err errors.BaseError,
	w http.ResponseWriter,
	r *http.Request,
) {
	resp := ErrorResponse{
		Code:    err.ErrorCode(),
		Message: err.ErrorMessage(),
	}
	if err := render.Render(w, r, resp); err != nil {
		panic(err)
	}
}

type ErrorResponse struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func (e ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(int(e.Code / 1000))
	return nil
}

func ChiErrorHandler() func(w http.ResponseWriter, r *http.Request, err error) {
	// err: pure Go error type
	// 		This handler is called only when passed params go against API definition
	return func(w http.ResponseWriter, r *http.Request, err error) {
		RespondWithError(
			errors.NewBadRequestDefaultError(),
			w, r,
		)
	}
}

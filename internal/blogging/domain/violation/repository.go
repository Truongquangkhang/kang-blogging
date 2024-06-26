package violation

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/common/model"
)

type Repository interface {
	GetViolationByParams(
		ctx context.Context,
		params ParamsGetViolations,
	) ([]model.Violation, int32, error)
}

type ParamsGetViolations struct {
	Page     int32
	PageSize int32
	Type     *string
	UserIDs  []string
}

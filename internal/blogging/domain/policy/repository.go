package policy

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/common/model"
)

type Repository interface {
	GetPolicies(
		ctx context.Context,
	) ([]model.Policy, error)

	UpdatePolicies(
		ctx context.Context,
		policies []model.Policy,
	) ([]model.Policy, error)
}

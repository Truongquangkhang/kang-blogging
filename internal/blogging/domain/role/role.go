package role

import (
	"context"
	"kang-blogging/internal/common/model"
)

type Repository interface {
	GetMapNameToRole(
		ctx context.Context,
		roleNames []string,
	) (map[string]model.Role, error)
}

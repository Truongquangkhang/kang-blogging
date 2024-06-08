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

	GetRoleById(
		ctx context.Context,
		roleId string,
	) (*model.Role, error)

	GetRoleByUserId(
		ctx context.Context,
		userId string,
	) (*model.Role, error)
}

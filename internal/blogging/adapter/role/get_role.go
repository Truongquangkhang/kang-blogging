package role

import (
	"context"
	"kang-blogging/internal/common/model"
)

func (u RoleRepository) GetMapNameToRole(
	ctx context.Context,
	roleNames []string,
) (map[string]model.Role, error) {
	var rs []model.Role
	err := u.gdb.DB().WithContext(ctx).Where("name in (?)", roleNames).Find(&rs).Error
	if err != nil {
		return nil, err
	}
	mapRole := map[string]model.Role{}
	for _, r := range rs {
		mapRole[r.Name] = r
	}
	return mapRole, nil
}

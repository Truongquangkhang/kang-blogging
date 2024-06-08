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

func (u RoleRepository) GetRoleById(
	ctx context.Context,
	roleId string,
) (*model.Role, error) {
	var rs model.Role
	err := u.gdb.DB().WithContext(ctx).Model(&model.Role{}).
		First(&rs, "id = ?", roleId).Error
	return &rs, err
}

func (u RoleRepository) GetRoleByUserId(
	ctx context.Context,
	userId string,
) (*model.Role, error) {
	var rs model.Role
	err := u.gdb.DB().WithContext(ctx).Model(&model.Role{}).
		Select("roles.*").
		Joins("inner join users on users.role_id = roles.id").
		Where("users.id = ?", userId).First(&rs).Error
	return &rs, err
}

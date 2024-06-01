package account

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	gormAdapter "kang-blogging/internal/common/db"
	"kang-blogging/internal/common/model"
)

type AccountRepository struct {
	gdb gormAdapter.DBAdapter
}

const repoName = "AccountRepository"

var logger *logrus.Entry = logrus.StandardLogger().
	WithField("logger", repoName)

func NewRepository() *AccountRepository {
	return &AccountRepository{
		gdb: gormAdapter.GetDBInstance(),
	}
}

func (u AccountRepository) GetRoleUserByID(
	ctx context.Context,
	id string,
) (*model.Role, error) {
	var role *model.Role
	err := u.gdb.DB().WithContext(ctx).Model(&model.Role{}).
		Joins("join users on users.role_id = roles.id").
		Where("users.id = ?", id).
		First(&role).Error
	if err != nil {
		return nil, err
	}
	return role, nil
}

package user

import (
	"errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/domain/user"
	gormAdapter "kang-blogging/internal/common/db"
	"time"
)

type UserRepository struct {
	gdb gormAdapter.DBAdapter
}

const repoName = "UserRepository"

var logger *logrus.Entry = logrus.StandardLogger().
	WithField("logger", repoName)

func NewRepository() *UserRepository {
	return &UserRepository{
		gdb: gormAdapter.GetDBInstance(),
	}
}

func (u UserRepository) GetInfoFromMultiTable(
	ctx context.Context,
) (*user.SystemInfo, error) {
	timeADayBefore := time.Now().Unix() - 86400

	rawQuery := `
		(SELECT COUNT(1) AS tt FROM blogs)
		union all
		(SELECT COUNT(1) AS tt FROM comments)
		union all
		(SELECT COUNT(1) AS tt FROM users)
		union all
		(SELECT COUNT(1) AS tt FROM categories)
		union all
		(SELECT COUNT(1) AS tt FROM blogs WHERE UNIX_TIMESTAMP(created_at) >= ?)
	 	union all
		(SELECT COUNT(1) AS tt FROM comments WHERE UNIX_TIMESTAMP(created_at) >= ?)
		union all
		(SELECT COUNT(1) AS tt FROM users WHERE UNIX_TIMESTAMP(created_at) >= ?)`

	var tt []int32
	err := u.gdb.DB().WithContext(ctx).
		Raw(rawQuery, timeADayBefore, timeADayBefore, timeADayBefore).
		Scan(&tt).Error
	if err != nil {
		return nil, err
	}
	if len(tt) != 7 {
		return nil, errors.New("not fix fields")
	}
	result := &user.SystemInfo{
		TotalBlogs:           tt[0],
		TotalComments:        tt[1],
		TotalUsers:           tt[2],
		TotalCategories:      tt[3],
		BlogIncreaseInDay:    tt[4],
		CommentIncreaseInDay: tt[5],
		UserIncreaseInDay:    tt[6],
	}
	return result, nil
}

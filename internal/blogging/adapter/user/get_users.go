package user

import (
	"context"
	"fmt"
	"kang-blogging/internal/blogging/domain/user"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/model"
	"kang-blogging/internal/common/utils"
)

func (u UserRepository) GetUsers(
	ctx context.Context,
	params user.UserParams,
) ([]model.User, int32, error) {
	var users []model.User
	var total int64
	query := u.gdb.DB().WithContext(ctx).Model(&model.User{})
	limit, offset := utils.PagePageSizeToLimitOffset(params.Page, params.PageSize)
	if params.SearchName != nil && params.SearchBy != nil {
		switch *params.SearchBy {
		case "name":
			query = query.Where("name LIKE ?", fmt.Sprintf("%%%s%%", *params.SearchName))
		case "email":
			query = query.Where("email LIKE ?", fmt.Sprintf("%%%s%%", *params.SearchName))
		case "display_name":
			query = query.Where("display_name LIKE ?", fmt.Sprintf("%%%s%%", *params.SearchName))
		case "phone_number":
			query = query.Where("phone_number LIKE ?", fmt.Sprintf("%%%s%%", *params.SearchName))
		default:
			return nil, 0, errors.NewBadRequestError("invalid search name")
		}
	}
	err := query.
		Select("users.*, count(blogs.id) as total_blogs, count(comments.id) as total_comments").
		Joins("left join blogs on users.id = blogs.author_id").
		Joins("left join comments on users.id = comments.user_id").
		Group("users.id").Count(&total).
		Offset(int(offset)).Limit(int(limit)).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}
	return users, int32(total), nil
}

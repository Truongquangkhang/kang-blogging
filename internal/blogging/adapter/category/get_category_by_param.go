package category

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/domain/category"
	"kang-blogging/internal/common/model"
	"kang-blogging/internal/common/utils"
)

func (r *CategoryRepository) GetCategoriesByParam(
	ctx context.Context,
	param category.ParamGetCategories,
) ([]category.ResultGetCategories, int32, error) {
	var categories []category.ResultGetCategories
	var count int64
	limit, offset := utils.PagePageSizeToLimitOffset(param.Page, param.PageSize)
	query := r.gdb.DB().WithContext(ctx).Model(model.Category{}).
		Select("categories.*, count(blog_categories.id) as blog_count").
		Joins("LEFT JOIN blog_categories ON categories.id = blog_categories.category_id").
		Group("categories.id")

	if param.SearchByName != nil {
		query = query.Where("name like ?", "%"+*param.SearchByName+"%")
	}
	errCount := query.Count(&count).Error
	if errCount != nil {
		return nil, 0, errCount
	}

	if param.SortBy != nil {
		if *param.SortBy == "created_at" {
			query = query.Order("created_at desc")
		}
		if *param.SortBy == "total_blog" {
			query = query.Order("blog_count desc")
		}
	}
	err := query.Limit(int(limit)).Offset(int(offset)).
		Find(&categories).Error
	if err != nil {
		return nil, 0, err
	}
	return categories, int32(count), nil
}

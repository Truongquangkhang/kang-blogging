package category

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
)

func GetCategories(
	ctx context.Context,
	request *blogging.GetCategoriesRequest,
) (*blogging.GetCategoriesResponse, error) {
	return &blogging.GetCategoriesResponse{
		Code:    0,
		Message: "Success",
	}, nil
}

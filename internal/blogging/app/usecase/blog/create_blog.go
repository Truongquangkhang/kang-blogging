package blog

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/blog"
	"kang-blogging/internal/blogging/domain/blog_categories"
	"kang-blogging/internal/blogging/domain/category"
	"kang-blogging/internal/blogging/domain/user"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/model"
	"kang-blogging/internal/common/utils"
)

type CreateBlogParams struct {
	Name        string
	Description string
	Thumbnail   *string
	CategoryIds []string
	Content     *string
	AuthorId    string
}

type CreateBlogResult struct {
	Blog           model.Blog
	BlogCategories []model.Category
	Author         model.User
}

type CreateBlogHandler decorator.UsecaseHandler[CreateBlogParams, CreateBlogResult]

type createBlogHandler struct {
	blogRepo           blog.Repository
	categoryRepo       category.Repository
	blogCategoriesRepo blog_categories.Repository
	userRepo           user.Repository
}

func NewCreateBlogHandler(
	blogRepo blog.Repository,
	categoryRepo category.Repository,
	blogCategoriesRepo blog_categories.Repository,
	userRepo user.Repository,
	logger *logrus.Entry,
	metrics decorator.MetricsClient,
) CreateBlogHandler {
	if blogRepo == nil {
		panic("blogRepo is nil")
	}
	if categoryRepo == nil {
		panic("categoryRepo is nil")
	}
	if blogCategoriesRepo == nil {
		panic("blogCategoriesRepo is nil")
	}
	if userRepo == nil {
		panic("userRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[CreateBlogParams, CreateBlogResult](
		createBlogHandler{
			blogRepo:           blogRepo,
			categoryRepo:       categoryRepo,
			blogCategoriesRepo: blogCategoriesRepo,
			userRepo:           userRepo,
		},
		logger,
		metrics,
	)
}

func (c createBlogHandler) Handle(ctx context.Context, param CreateBlogParams) (CreateBlogResult, error) {
	err := param.Validate()
	if err != nil {
		return CreateBlogResult{}, err
	}
	// get author info
	author, err := c.userRepo.GetUserByID(ctx, param.AuthorId)
	if err != nil || author == nil {
		return CreateBlogResult{}, err
	}
	// get categories
	categories, err := c.categoryRepo.GetCategories(ctx, param.CategoryIds)
	if err != nil || len(categories) == 0 {
		return CreateBlogResult{}, fmt.Errorf("catch an error while get categories: %w", err)
	}
	// insert blog
	blogId := utils.GenUUID()
	blog, err := c.blogRepo.InsertBlog(ctx, &model.Blog{
		ID:        blogId,
		AuthorID:  param.AuthorId,
		Title:     param.Name,
		Summary:   &param.Description,
		Thumbnail: param.Thumbnail,
		Content:   param.Content,
	})
	if err != nil || blog == nil {
		return CreateBlogResult{}, err
	}

	// insert blog categories
	var blogCategories []model.BlogCategories
	for _, c := range categories {
		blogCategories = append(blogCategories, model.BlogCategories{
			ID:         utils.GenUUID(),
			BlogID:     blogId,
			CategoryID: c.ID,
		})
	}
	blogCategories, err = c.blogCategoriesRepo.InsertBlogCategories(ctx, blogCategories)
	if err != nil {
		return CreateBlogResult{}, err
	}

	return CreateBlogResult{
		Blog:           *blog,
		BlogCategories: categories,
		Author:         *author,
	}, nil
}

func (p *CreateBlogParams) Validate() error {
	if p.Name == "" || p.Description == "" || len(p.CategoryIds) == 0 {
		return errors.NewBadRequestError("Invalid params")
	}
	return nil
}

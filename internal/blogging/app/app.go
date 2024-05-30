package app

import (
	"kang-blogging/internal/blogging/app/usecase/blog"
	"kang-blogging/internal/blogging/app/usecase/category"
	"kang-blogging/internal/blogging/app/usecase/comment"
	"kang-blogging/internal/blogging/app/usecase/iam"
	"kang-blogging/internal/blogging/app/usecase/image"
	"kang-blogging/internal/blogging/app/usecase/management"
	"kang-blogging/internal/blogging/app/usecase/user"
)

type Application struct {
	IAMUsecases       IAMUsecases
	UserUsecase       UserUsecase
	BlogUsecase       BlogUsecase
	CategoryUsecase   CategoryUsecase
	CommentUsecase    CommentUsecase
	ImageUsecase      ImageUsecase
	ManagementUsecase ManagementUsecase
}

type IAMUsecases struct {
	Login              iam.LoginHandler
	Register           iam.RegisterHandler
	CheckExistUsername iam.CheckExistUsernameHandler
	RefreshAccessToken iam.RefreshAccessTokenHandler
}

type UserUsecase struct {
	GetUsers      user.GetUsersHandler
	GetUserDetail user.GetUserDetailHandler
	UpdateUser    user.UpdateUserHandler
}

type BlogUsecase struct {
	GetBlogs         blog.GetBlogsHandler
	CreateBlog       blog.CreateBlogHandler
	GetBlogDetail    blog.GetBlogDetailHandler
	UpdateBlogDetail blog.UpdateBlogDetailHandler
	DeleteBlogDetail blog.DeleteBlogDetailHandler
}

type CategoryUsecase struct {
	GetCategories category.GetCategoriesHandler
}

type CommentUsecase struct {
	GetBlogComments   comment.GetBlogCommentsHandler
	CreateBlogComment comment.CreateBlogCommentHandler
	GetComments       comment.GetCommentsHandler
}

type ImageUsecase struct {
	UploadImage image.UploadImageHandler
}

type ManagementUsecase struct {
	GetDashboard management.GetDashboardHandler
}

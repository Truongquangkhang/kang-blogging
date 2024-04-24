package blog

import "context"

type Repository interface {
	InsertBlog(
		ctx context.Context)
}

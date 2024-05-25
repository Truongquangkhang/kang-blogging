package blog

type BlogsParams struct {
	Page        int32
	PageSize    int32
	SearchName  *string
	SearchBy    *string
	CategoryIds []string
	AuthorIds   []string
	SortBy      *string
}

package utils

func PagePageSizeToLimitOffset(page int32, pageSize int32) (int32, int32) {
	offset := (page - 1) * pageSize
	limit := pageSize
	return limit, offset
}

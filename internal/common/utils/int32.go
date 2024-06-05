package utils

func ToInt32Value(v *int32) int32 {
	if v == nil {
		return 0
	}
	return *v
}

func ToInt32Pointer(v int32) *int32 {
	return &v
}

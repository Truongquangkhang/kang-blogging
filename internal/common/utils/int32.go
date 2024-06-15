package utils

import "strconv"

func ToInt32Value(v *int32) int32 {
	if v == nil {
		return 0
	}
	return *v
}

func ToInt32Pointer(v int32) *int32 {
	return &v
}

func ConvertStringToInt32(s string) (int32, error) {
	value, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return int32(value), nil
}

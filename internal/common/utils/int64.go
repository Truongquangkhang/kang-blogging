package utils

func PointerInt64ToValue(value *int64) int64 {
	if value == nil {
		return 0
	}
	return *value
}

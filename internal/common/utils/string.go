package utils

func ToStringValue(str *string) string {
	if str == nil {
		return ""
	}
	return *str
}

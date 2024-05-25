package utils

func ToStringValue(str *string) string {
	if str == nil {
		return ""
	}
	return *str
}

func ToStringPointerValue(str string) *string {
	if str == "" {
		return nil
	}
	return &str
}

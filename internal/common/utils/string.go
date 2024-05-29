package utils

import "strings"

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

func SplitStringSeparateCommaToSlice(str *string) []string {
	if str == nil {
		return []string{}
	}
	return strings.Split(*str, ",")
}

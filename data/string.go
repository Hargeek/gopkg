package data

// GetStringValue 用于获取字符串指针的值,避免空指针异常
func GetStringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

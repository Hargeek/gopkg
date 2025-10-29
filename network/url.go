package network

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
)

// GetBasePathFromURL 从URL获取最后一个路径
func GetBasePathFromURL(URL string) (string, error) {
	u, err := url.Parse(URL)
	if err != nil {
		return "", err
	}
	parts := strings.Split(u.Path, "/")
	if len(parts) > 0 {
		return parts[len(parts)-1], nil
	}
	return "", errors.New("invalid URL")
}

// NormalizePathFor 将带参数的路径转换为通配符格式
// 例如：/api/v1/user/123 -> /api/v1/user/*
func NormalizePathFor(path string) string {
	// 处理空路径和根路径
	if path == "" || path == "/" {
		return path
	}

	// 标准化路径：移除开头和结尾的斜杠，处理连续斜杠
	normalizedPath := strings.Trim(path, "/")
	if normalizedPath == "" {
		return "/"
	}

	// 分割路径并过滤空段
	var parts []string
	for _, part := range strings.Split(normalizedPath, "/") {
		if part != "" {
			parts = append(parts, part)
		}
	}

	if len(parts) == 0 {
		return "/"
	}

	// 检查最后一段是否为数字（路径参数）
	lastPart := parts[len(parts)-1]
	isNumeric := func(s string) bool {
		_, err := strconv.Atoi(s)
		return err == nil
	}
	if isNumeric(lastPart) {
		// 将最后一段替换为通配符
		parts[len(parts)-1] = "*"
		return "/" + strings.Join(parts, "/")
	}

	return "/" + strings.Join(parts, "/")
}

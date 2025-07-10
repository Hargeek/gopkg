package gonic

import (
	"github.com/gin-gonic/gin"
)

// GetClientIP 获取客户端真实IP
func GetClientIP(c *gin.Context) string {
	// 1. X-Forwarded-For
	ip := c.GetHeader("X-Forwarded-For")
	if ip != "" {
		return ip
	}

	// 2. X-Real-Ip
	ip = c.GetHeader("X-Real-Ip")
	if ip != "" {
		return ip
	}

	// 3. Remoteip（自定义header）
	ip = c.GetHeader("Remoteip")
	if ip != "" {
		return ip
	}

	// 4. Gin默认
	return c.ClientIP()
}

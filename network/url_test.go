package network

import (
	"testing"
)

// TestNormalizePathForAuth 测试路径标准化功能
func TestNormalizePathForAuth(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "数字参数路径",
			input:    "/api/v1/user/123",
			expected: "/api/v1/user/*",
		},
		{
			name:     "多个数字参数",
			input:    "/api/v1/user/123/profile/456",
			expected: "/api/v1/user/123/profile/*",
		},
		{
			name:     "非数字参数路径",
			input:    "/api/v1/user/admin",
			expected: "/api/v1/user/admin",
		},
		{
			name:     "混合参数路径",
			input:    "/api/v1/user/admin/profile/789",
			expected: "/api/v1/user/admin/profile/*",
		},
		{
			name:     "根路径",
			input:    "/",
			expected: "/",
		},
		{
			name:     "空路径",
			input:    "",
			expected: "",
		},
		{
			name:     "单段数字路径",
			input:    "/123",
			expected: "/*",
		},
		{
			name:     "单段非数字路径",
			input:    "/admin",
			expected: "/admin",
		},
		{
			name:     "ORN策略路径",
			input:    "/api/v1/authorization/orn/policy/1",
			expected: "/api/v1/authorization/orn/policy/*",
		},
		{
			name:     "ORN策略列表路径",
			input:    "/api/v1/authorization/orn/policy",
			expected: "/api/v1/authorization/orn/policy",
		},
		{
			name:     "负数参数",
			input:    "/api/v1/user/-123",
			expected: "/api/v1/user/*",
		},
		{
			name:     "零参数",
			input:    "/api/v1/user/0",
			expected: "/api/v1/user/*",
		},
		{
			name:     "大数字参数",
			input:    "/api/v1/user/999999999",
			expected: "/api/v1/user/*",
		},
		{
			name:     "包含特殊字符的非数字",
			input:    "/api/v1/user/user-123",
			expected: "/api/v1/user/user-123",
		},
		{
			name:     "UUID格式参数",
			input:    "/api/v1/user/550e8400-e29b-41d4-a716-446655440000",
			expected: "/api/v1/user/550e8400-e29b-41d4-a716-446655440000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NormalizePathForAuth(tt.input)
			if result != tt.expected {
				t.Errorf("NormalizePathForAuth(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

// TestNormalizePathForAuthEdgeCases 测试边界情况
func TestNormalizePathForAuthEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "只有斜杠",
			input:    "/",
			expected: "/",
		},
		{
			name:     "多个斜杠",
			input:    "///",
			expected: "/",
		},
		{
			name:     "末尾斜杠",
			input:    "/api/v1/user/123/",
			expected: "/api/v1/user/*",
		},
		{
			name:     "开头和末尾斜杠",
			input:    "/api/v1/user/123/",
			expected: "/api/v1/user/*",
		},
		{
			name:     "连续斜杠",
			input:    "/api//v1///user/123",
			expected: "/api/v1/user/*",
		},
		{
			name:     "空段",
			input:    "/api//v1/user/123",
			expected: "/api/v1/user/*",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NormalizePathForAuth(tt.input)
			if result != tt.expected {
				t.Errorf("NormalizePathForAuth(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

// TestNormalizePathForAuthRealWorld 测试真实世界场景
func TestNormalizePathForAuthRealWorld(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "用户详情API",
			input:    "/api/v1/users/123",
			expected: "/api/v1/users/*",
		},
		{
			name:     "订单详情API",
			input:    "/api/v1/orders/456",
			expected: "/api/v1/orders/*",
		},
		{
			name:     "产品详情API",
			input:    "/api/v1/products/789",
			expected: "/api/v1/products/*",
		},
		{
			name:     "嵌套资源API",
			input:    "/api/v1/users/123/orders/456",
			expected: "/api/v1/users/123/orders/*",
		},
		{
			name:     "管理API",
			input:    "/api/v1/admin/users/123",
			expected: "/api/v1/admin/users/*",
		},
		{
			name:     "RESTful资源列表",
			input:    "/api/v1/users",
			expected: "/api/v1/users",
		},
		{
			name:     "RESTful资源创建",
			input:    "/api/v1/users",
			expected: "/api/v1/users",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NormalizePathForAuth(tt.input)
			if result != tt.expected {
				t.Errorf("NormalizePathForAuth(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

// BenchmarkNormalizePathForAuth 性能测试
func BenchmarkNormalizePathForAuth(b *testing.B) {
	testPaths := []string{
		"/api/v1/user/123",
		"/api/v1/user/admin",
		"/api/v1/authorization/orn/policy/1",
		"/api/v1/users/123/orders/456",
		"/api/v1/products/789",
	}

	for i := 0; i < b.N; i++ {
		for _, path := range testPaths {
			NormalizePathForAuth(path)
		}
	}
}

// TestNormalizePathForAuthConsistency 测试一致性
func TestNormalizePathForAuthConsistency(t *testing.T) {
	// 测试多次调用结果一致
	path := "/api/v1/user/123"
	expected := "/api/v1/user/*"

	for i := 0; i < 100; i++ {
		result := NormalizePathForAuth(path)
		if result != expected {
			t.Errorf("第 %d 次调用结果不一致: got %q, expected %q", i+1, result, expected)
		}
	}
}

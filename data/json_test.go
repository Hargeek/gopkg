package data

import (
	"fmt"
	"testing"
)

func TestJsonArrayToCSV(t *testing.T) {
	tests := []struct {
		name     string
		input    []map[string]interface{}
		expected string
	}{
		{
			name: "基本测试用例 - 用户提供的示例",
			input: []map[string]interface{}{
				{
					"PType": "p",
					"V0":    "p1",
					"V1":    "p2",
					"V2":    "p3",
					"V3":    "p4",
				},
				{
					"PType": "g",
					"V0":    "g1",
					"V1":    "g2",
					"V2":    "g3",
					"V3":    "g4",
				},
			},
			expected: "p,p1,p2,p3,p4\ng,g1,g2,g3,g4",
		},
		{
			name: "单个对象",
			input: []map[string]interface{}{
				{
					"name": "张三",
					"age":  "25",
					"city": "北京",
				},
			},
			expected: "25,北京,张三",
		},
		{
			name: "不同字段的对象",
			input: []map[string]interface{}{
				{
					"name": "张三",
					"age":  "25",
				},
				{
					"name": "李四",
					"city": "上海",
				},
			},
			expected: "25,,张三\n,上海,李四",
		},
		{
			name:     "空数组",
			input:    []map[string]interface{}{},
			expected: "",
		},
		{
			name: "包含null值",
			input: []map[string]interface{}{
				{
					"name": "张三",
					"age":  nil,
					"city": "北京",
				},
			},
			expected: ",北京,张三",
		},
		{
			name: "包含数字和布尔值",
			input: []map[string]interface{}{
				{
					"name":    "张三",
					"age":     25,
					"married": true,
					"score":   85.5,
				},
			},
			expected: "25,true,张三,85.5",
		},
		{
			name: "空值混合测试 - 您提到的示例",
			input: []map[string]interface{}{
				{
					"PType": "p",
					"V0":    "",
					"V1":    "p2",
					"V2":    "p3",
					"V3":    "p4",
				},
			},
			expected: "p,,p2,p3,p4",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := JsonArrayToCSV(tt.input)

			if result != tt.expected {
				t.Errorf("结果不匹配.\n期望:\n%s\n实际:\n%s", tt.expected, result)
			}
			fmt.Printf("转换结果:\n%s\n", result)
		})
	}
}

func TestJsonStringToCSV(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		hasError bool
	}{
		{
			name: "基本测试用例 - 用户提供的示例",
			input: `[
				{
					"PType": "p",
					"V0": "p1",
					"V1": "p2",
					"V2": "p3",
					"V3": "p4"
				},
				{
					"PType": "g",
					"V0": "g1",
					"V1": "g2",
					"V2": "g3",
					"V3": "g4"
				}
			]`,
			expected: "p,p1,p2,p3,p4\ng,g1,g2,g3,g4",
			hasError: false,
		},
		{
			name:     "空数组",
			input:    `[]`,
			expected: "",
			hasError: false,
		},
		{
			name:     "无效JSON",
			input:    `[{"name": "张三"`,
			expected: "",
			hasError: true,
		},
		{
			name:     "不是数组而是对象",
			input:    `{"name": "张三"}`,
			hasError: true,
		},
		{
			name:     "空字符串",
			input:    ``,
			hasError: true,
		},
		{
			name:     "只有空白字符",
			input:    `   `,
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := JsonStringToCSV(tt.input)

			if tt.hasError {
				if err == nil {
					t.Errorf("期望有错误，但没有返回错误")
				}
				return
			}

			if err != nil {
				t.Errorf("不期望有错误，但返回了错误: %v", err)
				return
			}

			if result != tt.expected {
				t.Errorf("结果不匹配.\n期望:\n%s\n实际:\n%s", tt.expected, result)
			}
			fmt.Printf("转换结果:\n%s\n", result)
		})
	}
}

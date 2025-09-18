package data

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

// JsonArrayToCSV 将JSON数组转换为CSV格式
func JsonArrayToCSV(jsonArray []map[string]interface{}) string {
	// 空数组返回空字符串
	if len(jsonArray) == 0 {
		return ""
	}

	// 收集所有唯一的字段名并排序
	fieldSet := make(map[string]bool)
	for _, obj := range jsonArray {
		for key := range obj {
			fieldSet[key] = true
		}
	}

	var fields []string
	for field := range fieldSet {
		fields = append(fields, field)
	}
	sort.Strings(fields) // 排序确保输出顺序一致

	// 构建CSV行
	var csvLines []string
	for _, obj := range jsonArray {
		var values []string
		for _, field := range fields {
			if value, exists := obj[field]; exists && value != nil {
				values = append(values, fmt.Sprintf("%v", value))
			} else {
				values = append(values, "") // 空值
			}
		}
		csvLines = append(csvLines, strings.Join(values, ","))
	}

	return strings.Join(csvLines, "\n")
}

// JsonStringToCSV 从JSON字符串解析并转换为CSV格式
func JsonStringToCSV(jsonArrayStr string) (string, error) {
	var jsonArray []map[string]interface{}
	if err := json.Unmarshal([]byte(jsonArrayStr), &jsonArray); err != nil {
		return "", fmt.Errorf("JSON解析失败: %v", err)
	}
	return JsonArrayToCSV(jsonArray), nil
}

package network

import (
	"fmt"
	"net/url"
	"reflect"
)

// ConvertToQueryParams 将结构体转换为查询参数，参数名为form tag
func ConvertToQueryParams(params interface{}) (url.Values, error) {
	v := url.Values{}

	// 检查是否为指针
	val := reflect.ValueOf(params)
	if val.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("参数必须是结构体指针")
	}

	// 检查指针是否为 nil
	if val.IsNil() {
		return nil, fmt.Errorf("参数不能为 nil")
	}

	// 获取指针指向的实际值
	elem := val.Elem()

	// 检查是否为结构体
	if elem.Kind() != reflect.Struct {
		return nil, fmt.Errorf("参数必须是结构体指针")
	}

	// 遍历结构体字段
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Type().Field(i)
		fieldValue := elem.Field(i)
		tag := field.Tag.Get("form")

		if tag != "" {
			// 如果是结构体类型，返回错误
			if fieldValue.Kind() == reflect.Struct {
				return nil, fmt.Errorf("不支持嵌套结构体字段: %s", field.Name)
			}
			v.Set(tag, fmt.Sprintf("%v", fieldValue.Interface()))
		}
	}

	return v, nil
}

package data

import (
	"fmt"
	"reflect"
	"strings"
)

// FormatStructFields 格式化结构体字段
func FormatStructFields(v interface{}) string {
	if v == nil {
		return ""
	}

	val := reflect.ValueOf(v)
	// 如果是指针，获取其底层值
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return ""
		}
		val = val.Elem()
	}

	// 确保是结构体
	if val.Kind() != reflect.Struct {
		return fmt.Sprintf("%v", v)
	}

	var result []string
	t := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := t.Field(i).Name

		var fieldValue string
		switch {
		case !field.IsValid():
			fieldValue = "invalid"
		case field.Kind() == reflect.Ptr:
			if field.IsNil() {
				fieldValue = "nil"
			} else {
				// 处理不同类型的指针
				switch field.Interface().(type) {
				case *string:
					fieldValue = *field.Interface().(*string)
				case *float64:
					fieldValue = fmt.Sprintf("%f", *field.Interface().(*float64))
				case *int32, *int64, *int:
					fieldValue = fmt.Sprintf("%d", field.Elem().Int())
				case *bool:
					fieldValue = fmt.Sprintf("%v", *field.Interface().(*bool))
				default:
					fieldValue = fmt.Sprintf("%v", field.Elem().Interface())
				}
			}
		case field.Kind() == reflect.Struct:
			// 递归处理嵌套结构体
			fieldValue = FormatStructFields(field.Interface())
		case field.Kind() == reflect.Slice || field.Kind() == reflect.Array:
			// 处理切片和数组
			if field.Len() == 0 {
				fieldValue = "[]"
			} else {
				var elements []string
				for j := 0; j < field.Len(); j++ {
					elem := field.Index(j)
					elements = append(elements, fmt.Sprintf("%v", elem.Interface()))
				}
				fieldValue = fmt.Sprintf("[%s]", strings.Join(elements, ", "))
			}
		default:
			// 处理其他类型
			fieldValue = fmt.Sprintf("%v", field.Interface())
		}

		result = append(result, fmt.Sprintf("%s: %s", fieldName, fieldValue))
	}

	return strings.Join(result, ", ")
}

// PrintStructFieldsAndValues 打印结构体字段名和值
func PrintStructFieldsAndValues(v interface{}) {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem() // 如果是指针，则获取其所指向的元素
	}

	if val.Kind() != reflect.Struct {
		fmt.Println("Expected a struct")
		return
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)
		fmt.Printf("Field Name: '%s', Field Value: '%v'\n", field.Name, value.Interface())
	}
}

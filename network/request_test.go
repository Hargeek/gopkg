package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 定义用于测试的结构体
type TestStruct struct {
	Name   string `form:"name"`
	Age    int    `form:"age"`
	Gender string // 无 form 标签
	Active bool   `form:"active"`
}

func TestConvertToQueryParams(t *testing.T) {
	t.Run("TC01: 正常结构体", func(t *testing.T) {
		input := &TestStruct{
			Name:   "Alice",
			Age:    30,
			Gender: "female",
			Active: true,
		}
		values, err := ConvertToQueryParams(input)
		assert.NoError(t, err)
		assert.Equal(t, "Alice", values.Get("name"))
		assert.Equal(t, "30", values.Get("age"))
		assert.Equal(t, "true", values.Get("active"))
		assert.Empty(t, values.Get("Gender")) // 无 form 标签
	})

	t.Run("TC02: 无 form 标签字段", func(t *testing.T) {
		type NoFormTagStruct struct {
			ID   int
			Info string
		}
		input := &NoFormTagStruct{ID: 1, Info: "test"}
		values, err := ConvertToQueryParams(input)
		assert.NoError(t, err)
		assert.Empty(t, values.Encode()) // 应为空
	})

	t.Run("TC03: 非结构体指针", func(t *testing.T) {
		input := "not a struct"
		_, err := ConvertToQueryParams(input)
		assert.Error(t, err)
	})

	t.Run("TC04: 多种字段类型", func(t *testing.T) {
		type MixedTypeStruct struct {
			IntVal   int     `form:"int"`
			FloatVal float64 `form:"float"`
			BoolVal  bool    `form:"bool"`
		}
		input := &MixedTypeStruct{
			IntVal:   42,
			FloatVal: 3.14,
			BoolVal:  false,
		}
		values, err := ConvertToQueryParams(input)
		assert.NoError(t, err)
		assert.Equal(t, "42", values.Get("int"))
		assert.Equal(t, "3.14", values.Get("float"))
		assert.Equal(t, "false", values.Get("bool"))
	})

	t.Run("TC05: 嵌套结构体（应返回错误）", func(t *testing.T) {
		type NestedStruct struct {
			User struct {
				Name string `form:"name"`
			} `form:"user"` // 嵌套结构体字段应该返回错误
		}
		input := &NestedStruct{}
		_, err := ConvertToQueryParams(input)
		assert.Error(t, err) // 期望返回错误
		assert.Contains(t, err.Error(), "不支持嵌套结构体字段")
	})

	t.Run("TC06: 无form标签的嵌套结构体（应被忽略）", func(t *testing.T) {
		type NestedStructNoTag struct {
			Name string `form:"name"`
			User struct {
				Age int `form:"age"`
			} // 无 form 标签，应该被忽略
		}
		input := &NestedStructNoTag{Name: "test"}
		values, err := ConvertToQueryParams(input)
		assert.NoError(t, err) // 不应该返回错误
		assert.Equal(t, "test", values.Get("name"))
		assert.Empty(t, values.Get("age")) // 嵌套字段不会被处理
	})
}

package data

import (
	"fmt"
	"reflect"
	"unsafe"
)

// PrintMemSize 打印任意值的内存占用
func PrintMemSize(name string, v any) string {
	sh, dp := MemSize(v)
	return fmt.Sprintf("%s: shallow=%s (%dB), deep=%s (%dB)", name, FormatBytes(sh), sh, FormatBytes(dp), dp)
}

// FormatBytes 将字节数格式化为可读字符串
func FormatBytes(b int64) string {
	const (
		KB = 1024
		MB = 1024 * KB
		GB = 1024 * MB
	)
	if b >= GB {
		return fmt.Sprintf("%.2fGB", float64(b)/float64(GB))
	}
	if b >= MB {
		return fmt.Sprintf("%.2fMB", float64(b)/float64(MB))
	}
	if b >= KB {
		return fmt.Sprintf("%.2fKB", float64(b)/float64(KB))
	}
	return fmt.Sprintf("%dB", b)
}

// MemSize 计算任意值的浅（shallow）与深（deep）内存占用（近似值）
// 说明：
// - 浅占用：unsafe.Sizeof(v) 或其等价头部开销
// - 深占用：包含其引用/子元素（切片底层数组、map 键/值、字符串数据、结构体字段等）的总和
// 注意：
// - 该计算为近似值：未考虑运行时对齐、map 内部桶开销、共享底层数组的重复引用等复杂情况
// - 通过 visited 去重，避免重复统计相同地址
func MemSize(v any) (shallow int64, deep int64) {
	visited := make(map[unsafe.Pointer]struct{})
	val := reflect.ValueOf(v)
	shallow = int64(shallowSizeOf(val))
	deep = int64(deepSizeOf(val, visited))
	return
}

// shallowSizeOf 仅统计值本身（头部）大小
func shallowSizeOf(v reflect.Value) uintptr {
	if !v.IsValid() {
		return 0
	}
	// 对接口，取其动态值的头部开销
	if v.Kind() == reflect.Interface && !v.IsNil() {
		return unsafe.Sizeof(v.Interface())
	}
	// 对切片/字符串/映射/指针等，Sizeof 返回其头部大小
	return v.Type().Size()
}

// deepSizeOf 递归统计包含的全部数据大小（近似）
func deepSizeOf(v reflect.Value, visited map[unsafe.Pointer]struct{}) uintptr {
	if !v.IsValid() {
		return 0
	}
	// 处理接口，解包其动态值
	if v.Kind() == reflect.Interface && !v.IsNil() {
		return deepSizeOf(v.Elem(), visited) + shallowSizeOf(v)
	}

	switch v.Kind() {
	case reflect.Ptr:
		if v.IsNil() {
			return shallowSizeOf(v)
		}
		ptr := unsafe.Pointer(v.Pointer())
		if _, ok := visited[ptr]; ok {
			return shallowSizeOf(v)
		}
		visited[ptr] = struct{}{}
		return shallowSizeOf(v) + deepSizeOf(v.Elem(), visited)

	case reflect.String:
		// 头部 + 数据
		return shallowSizeOf(v) + uintptr(v.Len())

	case reflect.Slice:
		// 头部 + 底层数组数据（按 len 统计）
		base := shallowSizeOf(v)
		if v.IsNil() {
			return base
		}
		total := base
		for i := 0; i < v.Len(); i++ {
			total += deepSizeOf(v.Index(i), visited)
		}
		return total

	case reflect.Array:
		total := uintptr(0)
		for i := 0; i < v.Len(); i++ {
			total += deepSizeOf(v.Index(i), visited)
		}
		return total

	case reflect.Map:
		// 头部 + 键/值（不含桶等运行时开销）
		base := shallowSizeOf(v)
		if v.IsNil() {
			return base
		}
		total := base
		for _, key := range v.MapKeys() {
			val := v.MapIndex(key)
			total += deepSizeOf(key, visited)
			total += deepSizeOf(val, visited)
		}
		return total

	case reflect.Struct:
		total := shallowSizeOf(v)
		for i := 0; i < v.NumField(); i++ {
			total += deepSizeOf(v.Field(i), visited)
		}
		return total

	case reflect.Chan, reflect.Func, reflect.UnsafePointer:
		// 仅计头部
		return shallowSizeOf(v)

	default:
		// 基本类型：bool, 数值等
		return shallowSizeOf(v)
	}
}

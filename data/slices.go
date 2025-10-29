package data

import (
	"reflect"
	"sort"
)

// StrSlicesContains 字符串切片中是否包含特定元素
func StrSlicesContains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

// UintSlicesContains 无符号整数切片中是否包含特定元素
func UintSlicesContains(slice []uint, item uint) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

// StrSlicesContainsOneElement 字符串切片中是否包含另一个字符串切片中的元素
func StrSlicesContainsOneElement(sliceOrigin []string, sliceCollection []string) bool {
	for _, a := range sliceOrigin {
		for _, b := range sliceCollection {
			if a == b {
				return true
			}
		}
	}
	return false
}

// StrSlicesEqualWithLoop 两个字符串切片是否相等(使用循环)
func StrSlicesEqualWithLoop(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	// compare the slices
	if (slice1 == nil) != (slice2 == nil) {
		return false
	}
	for i, a := range slice1 {
		if a != slice2[i] {
			return false
		}
	}
	return true
}

// IntDisorderSlicesEqualWithLoop 两个整数切片是否相等(使用循环且不考虑顺序)
func IntDisorderSlicesEqualWithLoop(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	// sort the slices
	sort.Slice(slice1, func(i, j int) bool { return slice1[i] < slice1[j] })
	sort.Slice(slice2, func(i, j int) bool { return slice2[i] < slice2[j] })
	// compare the slices
	if (slice1 == nil) != (slice2 == nil) {
		return false
	}
	for i, a := range slice1 {
		if a != slice2[i] {
			return false
		}
	}
	return true
}

// IntDisorderSlicesEqualWithReflect 两个整数切片是否相等(使用反射且不考虑顺序)
func IntDisorderSlicesEqualWithReflect(slice1, slice2 []int) bool {
	// sort the slices
	sort.Slice(slice1, func(i, j int) bool { return slice1[i] < slice1[j] })
	sort.Slice(slice2, func(i, j int) bool { return slice2[i] < slice2[j] })
	return reflect.DeepEqual(slice1, slice2)
}

// UintDisorderSlicesEqualWithLoop 两个无符号整数切片是否相等(使用循环且不考虑顺序)
func UintDisorderSlicesEqualWithLoop(slice1, slice2 []uint) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	// sort the slices
	sort.Slice(slice1, func(i, j int) bool { return slice1[i] < slice1[j] })
	sort.Slice(slice2, func(i, j int) bool { return slice2[i] < slice2[j] })
	// compare the slices
	if (slice1 == nil) != (slice2 == nil) {
		return false
	}
	for i, a := range slice1 {
		if a != slice2[i] {
			return false
		}
	}
	return true
}

// UintDisorderSlicesEqualWithReflect 两个无符号整数切片是否相等(使用反射且不考虑顺序)
func UintDisorderSlicesEqualWithReflect(slice1, slice2 []uint) bool {
	// sort the slices
	sort.Slice(slice1, func(i, j int) bool { return slice1[i] < slice1[j] })
	sort.Slice(slice2, func(i, j int) bool { return slice2[i] < slice2[j] })
	return reflect.DeepEqual(slice1, slice2)
}

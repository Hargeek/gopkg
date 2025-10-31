package data

// 现在已经可以用github.com/samber/lo的lo.Uniq函数代替以下去重函数

// RemoveDuplicates 数组去重
func RemoveDuplicates(nums []uint) []uint {
	seen := make(map[uint]struct{}, len(nums))
	res := make([]uint, 0, len(nums))
	for _, n := range nums {
		if _, ok := seen[n]; ok {
			continue
		}
		seen[n] = struct{}{}
		res = append(res, n)
	}
	return res
}

// RemoveDuplicateStrings 字符串数组去重
func RemoveDuplicateStrings(items []string) []string {
	seen := make(map[string]struct{}, len(items))
	res := make([]string, 0, len(items))
	for _, s := range items {
		if _, ok := seen[s]; ok {
			continue
		}
		seen[s] = struct{}{}
		res = append(res, s)
	}
	return res
}

package time

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// ParseDuration 解析时间字符串
func ParseDuration(d string) (time.Duration, error) {
	d = strings.TrimSpace(d)
	dr, err := time.ParseDuration(d)
	if err == nil {
		return dr, nil
	}
	if strings.Contains(d, "d") {
		index := strings.Index(d, "d")

		hour, _ := strconv.Atoi(d[:index])
		dr = time.Hour * 24 * time.Duration(hour)
		ndr, err := time.ParseDuration(d[index+1:])
		if err != nil {
			return dr, nil
		}
		return dr + ndr, nil
	}

	dv, err := strconv.ParseInt(d, 10, 64)
	return time.Duration(dv), err
}

// ParseHumanDurationMillis 解析毫秒时间为人类可读的时间区间
func ParseHumanDurationMillis(millis float64) string {
	totalMillis := int64(millis)       // 转换成整数毫秒
	milliseconds := totalMillis % 1000 // 剩余的毫秒
	seconds := totalMillis / 1000      // 总秒数
	minutes := seconds / 60            // 总分钟数
	hours := minutes / 60              // 总小时数
	days := hours / 24                 // 总天数

	seconds %= 60 // 剩余秒数
	minutes %= 60 // 剩余分钟数
	hours %= 24   // 剩余小时数

	var result []string

	if days > 0 {
		result = append(result, fmt.Sprintf("%d day", days))
	}
	if hours > 0 || (days > 0 && (minutes > 0 || seconds > 0 || milliseconds > 0)) {
		result = append(result, fmt.Sprintf("%d hour", hours))
	}
	if minutes > 0 || (hours > 0 && (seconds > 0 || milliseconds > 0)) {
		result = append(result, fmt.Sprintf("%d min", minutes))
	}
	if seconds > 0 || (minutes > 0 && milliseconds > 0) {
		result = append(result, fmt.Sprintf("%d sec", seconds))
	}
	if milliseconds > 0 {
		result = append(result, fmt.Sprintf("%d msec", milliseconds))
	}

	return strings.Join(result, " ")
}

// ParseHumanTimeCost 解析时间消耗为人类可读的时间区间
func ParseHumanTimeCost(start, end time.Time) string {
	//if end.Before(start) {
	//	return "", fmt.Errorf("end time must be after start time")
	//}
	return ParseHumanDurationMillis(float64(end.Sub(start).Milliseconds()))
}

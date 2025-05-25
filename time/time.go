package time

import "time"

const ctLayout = "2006-01-02 15:04:05"

type TimeString struct {
	time.Time
}

func (ct *TimeString) UnmarshalText(b []byte) (err error) {
	ct.Time, err = time.Parse(ctLayout, string(b))
	return
}

// UTC2CST UTC时间转换为CST时间
func UTC2CST(t time.Time) (timeStr string, err error) {
	utcTime, err := time.Parse(ctLayout, t.Format(ctLayout))
	if err != nil {
		return "", err
	}
	loc, _ := time.LoadLocation("Asia/Shanghai")
	cstTime := utcTime.In(loc)
	return cstTime.Format(ctLayout), nil
}

// UTC2CSTStr UTC时间转换为CST时间
func UTC2CSTStr(t string) (timeStr string, err error) {
	utcTime, err := time.Parse(ctLayout, t)
	if err != nil {
		return "", err
	}
	loc, _ := time.LoadLocation("Asia/Shanghai")
	cstTime := utcTime.In(loc)
	return cstTime.Format(ctLayout), nil
	//return utcTime.In(time.FixedZone("UTC+8", 8*3600)).Format(ctLayout), nil
}

// TimeStrToCompact 将时间字符串"2006-01-02 15:04:05"转换为"20060102150405"格式
func TimeStrToCompact(t string) (string, error) {
	parsed, err := time.Parse(ctLayout, t)
	if err != nil {
		return "", err
	}
	return parsed.Format("20060102150405"), nil
}

// TimeStrToTimestamp 将时间字符串"2006-01-02 15:04:05"转换为秒级时间戳
func TimeStrToTimestamp(t string) (int64, error) {
	parsed, err := time.Parse(ctLayout, t)
	if err != nil {
		return 0, err
	}
	return parsed.Unix(), nil
}

// TimeStrToMilliTimestamp 将时间字符串"2006-01-02 15:04:05"转换为毫秒级时间戳
func TimeStrToMilliTimestamp(t string) (int64, error) {
	parsed, err := time.Parse(ctLayout, t)
	if err != nil {
		return 0, err
	}
	return parsed.UnixNano() / 1e6, nil
}

// TimeToCompact 将 time.Time 转换为 "20060102150405" 格式
func TimeToCompact(t time.Time) string {
	return t.Format("20060102150405")
}

// TimeToTimestamp 将 time.Time 转换为秒级时间戳
func TimeToTimestamp(t time.Time) int64 {
	return t.Unix()
}

// TimeToMilliTimestamp 将 time.Time 转换为毫秒级时间戳
func TimeToMilliTimestamp(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

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

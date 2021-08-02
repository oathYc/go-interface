package utils

import (
	"github.com/pkg/errors"
	"time"
)

func StartOfWeek(inTime int64) int64 {
	now := time.Unix(inTime, 0)
	offset := int(now.Weekday()) - int(time.Monday)
	if offset < 0 {
		offset = 6
	}
	weekStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, -offset)
	return weekStart.Unix()
}

func StartOfDay(in int64) int64 {
	t := time.Unix(in, 0)
	outTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return outTime.Unix()
}

func EndOfDay(in int64) int64 {
	t := time.Unix(in, 0)
	outTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).AddDate(0, 0, 1)
	return outTime.Unix() - 1
}

func DatetimeStringToUnix(s string) (int64, error) {
	var (
		err error
		ret int64
	)

	sCount := len(s)

	if sCount != 10 && sCount != 19 {
		return ret, errors.Wrap(err, "时间格式错误eg:2006-01-02 00:00:00 | 2006-01-02")
	}
	if sCount == 10 {
		s = s + " 00:00:00"
	}
	timeL, err := time.ParseInLocation("2006-01-02 15:04:05", s, time.Local)
	if nil != err {
		return ret, err
	}
	ret = timeL.Unix()
	return ret, err

}

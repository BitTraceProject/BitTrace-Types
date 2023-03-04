package common

import (
	"strconv"
	"time"
)

// Timestamp 是 19 位长，精确到纳秒的时间戳
type (
	Timestamp int64
)

func FromTime(t time.Time) Timestamp {
	return Timestamp(t.UnixNano())
}

func FromNow() Timestamp {
	return FromTime(time.Now())
}

func FromInt64(t int64) Timestamp {
	return Timestamp(t)
}

func (t Timestamp) String() string {
	return strconv.FormatInt(int64(t), 10) // 如：1665807442207974500
}

func (t Timestamp) Format(layout string) string {
	return t.FormatTime().Format(layout)
}

func (t Timestamp) FormatString() string {
	return t.FormatTime().String()
}

func (t Timestamp) FormatTime() time.Time {
	return time.Unix(0, int64(t))
}

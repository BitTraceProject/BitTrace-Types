package structure

import "time"

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

// TODO add more method about Timestamp

func (t Timestamp) Format(layout string) string {
	return t.FormatTime().Format(layout)
}

func (t Timestamp) FormatString() string {
	return t.FormatTime().String()
}

func (t Timestamp) FormatTime() time.Time {
	return time.Unix(0, int64(t))
}

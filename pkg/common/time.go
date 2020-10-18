package common

import "time"

const (
	Day time.Duration = time.Hour * 24
	Year time.Duration = Day * 365
)

func GetMillisecond(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func GetTimeStampMillisecond() int64 {
	return GetMillisecond(time.Now())
}


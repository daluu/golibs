package time

import (
	"fmt"
	"time"
)

func Iso8601TimeParser(iso8601ts string) (time.Time, error) {
	ts, err := time.Parse("2006-01-02T15:04:05Z07:00", iso8601ts)
	if err != nil {
		return time.Time{}, fmt.Errorf("Error parsing ISO8601 timestamp: %s, err: %v\n", iso8601ts, err)
	}
	return ts, err
}

func UnixTimeParser(unixTime int64) time.Time {
	var ts time.Time
	if unixTime > 9999999999 {
		//ms since epoch
		sec := unixTime / 1000
		ns := (unixTime % 1000) * 1000 * 1000
		ts = time.Unix(sec, ns)
	} else {
		//secs since epoch
		ts = time.Unix(unixTime, 0)
	}
	return ts
}

func UnixMillisec(ts time.Time) int64 {
	return ts.UnixNano() / 1000 / 1000
}

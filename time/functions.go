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

func formatDuration(duration time.Duration, decimalPrecision int) string {
	formatString := "%." + fmt.Sprintf("%d", decimalPrecision) + "f"
	h := duration.Hours()
	m := duration.Minutes()
	s := duration.Seconds()
	ms := float64(s * 1000)
	//although, often no need go past ms for timings in real world?
	us := float64(s * 1000 * 1000)
	//ns := duration.Nanoseconds()
	formattedDuration := ""
	if h > 1 {
		//highest possible duration we can use, following go convention
		//don't deal with days, weeks, months, years
		formattedDuration = fmt.Sprintf(formatString, h) + "h"
	} else if m > 1 {
		formattedDuration = fmt.Sprintf(formatString, m) + "m"
	} else if s > 1 {
		formattedDuration = fmt.Sprintf(formatString, s) + "s"
	} else if ms > 1 {
		//future TODO: regex parse out digits for better precision than this method
		formattedDuration = fmt.Sprintf(formatString, ms) + "ms"
	} else if us > 1 {
		formattedDuration = fmt.Sprintf(formatString, us) + "µs"
	} else {
		//lowest possible duration we can use, & can't format decimal, int only
		formattedDuration = duration.String()
	}
	return formattedDuration
}

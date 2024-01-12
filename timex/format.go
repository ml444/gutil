package timex

import "time"

func DateTimeFormat(ts int64) string {
	if ts == 0 {
		return ""
	}
	return time.Unix(ts, 0).Format(time.DateTime)
}

func DateFormat(ts int64) string {
	if ts == 0 {
		return ""
	}
	return time.Unix(ts, 0).Format(time.DateOnly)
}

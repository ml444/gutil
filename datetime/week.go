package datetime

import "time"

// CurrentWeekdayTS Obtain the 0:00 a.m. timestamp of the specified day of the week
func CurrentWeekdayTS(wd time.Weekday) uint32 {
	now := time.Now()
	monday := now.AddDate(0, 0, -int(now.Weekday()-wd))
	mondayMidnight := time.Date(
		monday.Year(), monday.Month(), monday.Day(),
		0, 0, 0, 0, time.Local)
	return uint32(mondayMidnight.Unix())
}

func LastWeekdayTS(wd time.Weekday) uint32 {
	return CurrentWeekdayTS(wd) - (86400 * 7)
}

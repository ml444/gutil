package timex

import "time"

const defaultTimeLayout = time.DateTime // "2006-01-02 15:04:05"

type DT struct {
	t              time.Time
	loc            *time.Location
	layout         string
	startAtWeekday time.Weekday
}

func New() *DT {
	return &DT{
		t:              time.Now(),
		loc:            time.Local,
		layout:         defaultTimeLayout,
		startAtWeekday: time.Sunday,
	}
}

func (d *DT) SetLoc(loc *time.Location) *DT {
	d.loc = loc
	return d
}

func (d *DT) SetTime(t time.Time) *DT {
	d.t = t
	return d
}

func (d *DT) SetLayout(layout string) *DT {
	d.layout = layout
	return d
}

func (d *DT) SetStartAsWeekday(wd time.Weekday) *DT {
	d.startAtWeekday = wd
	return d
}

func (d *DT) AddDays(n int) *DT {
	d.t = d.t.AddDate(0, 0, n)
	return d
}

func (d *DT) AddMonths(n int) *DT {
	d.t = d.t.AddDate(0, n, 0)
	return d
}

func (d *DT) AddYears(n int) *DT {
	d.t = d.t.AddDate(n, 0, 0)
	return d
}

func (d *DT) Yesterday() *DT {
	d.t = d.t.Add(time.Hour * 24)
	return d
}

func (d *DT) Tomorrow() *DT {
	d.t = d.t.Add(-time.Hour * 24)
	return d
}

func (d *DT) OnWeekday(wd time.Weekday) *DT {
	d.t = d.t.AddDate(0, 0, -int(d.t.Weekday()-wd))
	return d
}

func (d *DT) ZeroClock() int64 {
	return d.OnClock(0)
}

// OnClock Obtain the 0:00 a.m. timestamp
func (d *DT) OnClock(n int) int64 {
	year, month, day := d.t.Date()
	return time.Date(year, month, day, n, 0, 0, 0, d.t.Location()).Unix()
}

// WeekRanges Get the weekly start and end times for the current month
// (from the 1st to the end of the month)
// (default starting on Sunday)
func (d *DT) WeekRanges() [][2]int64 {
	var weekRanges [][2]int64
	year, month, _ := d.t.Date()
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, d.loc)
	lastDay := firstDay.AddDate(0, 1, -1).Add(86399 * time.Second)

	start := firstDay
	end := firstDay.AddDate(0, 0, 6-int(firstDay.Weekday())).Add(86399 * time.Second)

	for start.Before(lastDay) {
		weekRanges = append(weekRanges, [2]int64{start.Unix(), end.Unix()})
		start = start.AddDate(0, 0, 7-int(start.Weekday()))
		end = end.AddDate(0, 0, 7)
		if end.After(lastDay) {
			end = lastDay
		}
	}
	return weekRanges
}

func (d *DT) Format(ts ...int64) string {
	if len(ts) == 0 {
		return ""
	}
	sec := ts[0]
	nsec := int64(0)
	if len(ts) > 1 {
		nsec = ts[1]
	}
	return time.Unix(sec, nsec).Format(d.layout)
}

func (d *DT) ParseTime(st string) error {
	t, err := time.ParseInLocation(d.layout, st, d.loc)
	if err != nil {
		return err
	}
	d.t = t
	return nil
}

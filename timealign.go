package timealign

import (
	"time"
)

func AlignToMinute(t time.Time) time.Time {
	return time.Unix(t.Unix()-int64(t.Second()), 0)
}

func AlignToMinutes(t time.Time, minutes int) time.Time {
	minutes = minutes % 60
	unixts := t.Unix()
	delta := t.Minute()%minutes*60 + t.Second()
	return time.Unix(unixts-int64(delta), 0)
}

func AlignToHour(t time.Time) time.Time {
	unixts := t.Unix()
	delta := t.Minute()*60 + t.Second()
	return time.Unix(unixts-int64(delta), 0)
}

func AlignToDay(t time.Time) time.Time {
	unixts := t.Unix()
	delta := t.Hour()*60*60 + t.Minute()*60 + t.Second()
	return time.Unix(unixts-int64(delta), 0)
}

func AlignToMonth(t time.Time) time.Time {
	unixts := t.Unix()
	delta := t.Day()*24*60*60 + t.Hour()*60*60 + t.Minute()*60 + t.Second()
	return time.Unix(unixts-int64(delta), 0)
}

package util

import (
	"fmt"
	"strings"
	"time"
)

var (
	ShortMonthNames   = []string{"JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"}
	ShortWeekdayNames = []string{"SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT"}
)

func CastWeekday(s string) (time.Weekday, error) {
	for i, n := range ShortWeekdayNames {
		if strings.EqualFold(n, s) {
			return time.Weekday(i), nil
		}
	}

	return -1, fmt.Errorf("cannot convert to weekday from %s", s)
}

func CastMonth(s string) (time.Month, error) {
	for i, n := range ShortMonthNames {
		if strings.EqualFold(n, s) {
			return time.Month(i + 1), nil
		}
	}

	return -1, fmt.Errorf("cannot convert to month from %s", s)
}

func LastOfMonth(t time.Time) int {
	return t.AddDate(0, 1, -t.Day()).Day()
}

func LastWdayOfMonth(t time.Time, w time.Weekday) int {
	lom := t.AddDate(0, 1, -t.Day())

	for i := lom; ; i = i.AddDate(0, 0, -1) {
		if i.Weekday() == w {
			return i.Day()
		}
	}
}

func NearestWeekday(t2 time.Time, day int) int {
	base := time.Date(t2.Year(), t2.Month(), day, 0, 0, 0, 0, t2.Location())
	lom := LastOfMonth(time.Date(t2.Year(), t2.Month(), 1, 0, 0, 0, 0, t2.Location()))

	if day == 1 {
		switch base.Weekday() {
		case time.Saturday:
			return 3
		case time.Sunday:
			return 2
		default:
			return day
		}
	} else if day >= lom {
		for i := 0; i <= 2; i++ {
			prev := base.AddDate(0, 0, -i)

			if prev.Month() == t2.Month() && prev.Weekday() != time.Saturday && prev.Weekday() != time.Sunday {
				return prev.Day()
			}
		}

		return 0
	} else if base.Weekday() == time.Saturday {
		return base.AddDate(0, 0, -1).Day()
	} else if base.Weekday() == time.Sunday {
		return base.AddDate(0, 0, 1).Day()
	} else {
		return day
	}
}

func NthDayOfWeek(t time.Time, wday time.Weekday, nth int) int {
	firstOfMonth := time.Date(t.Year(), t.Month(), 1, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
	offset := (wday + 7 - firstOfMonth.Weekday()) % 7
	nthDoW := firstOfMonth.AddDate(0, 0, 7*(nth-1)+int(offset))

	if nthDoW.Month() != t.Month() {
		return 0
	}

	return nthDoW.Day()
}

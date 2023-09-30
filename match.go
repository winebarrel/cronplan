package cronplan

import (
	"time"

	"github.com/winebarrel/cronplan/internal/util"
)

// minute =====================================================================

func (v *Minute) Match(t time.Time) bool {
	return v.Int() == t.Minute()
}

func (v *MinuteRange) Match(t time.Time) bool {
	minute := t.Minute()
	list, err := util.ListMinute(v.Start.Int(), v.End.Int())

	if err != nil {
		panic(err)
	}

	for _, i := range list {
		if i == minute {
			return true
		}
	}

	return false
}

func (e *MinuteExp) Match(t time.Time) bool {
	if e.Bottom != nil {
		minute := t.Minute()
		bottom := *e.Bottom

		if e.Range != nil {
			if bottom == 0 {
				return e.Range.Start.Match(t)
			}

			start := e.Range.Start.Int()
			end := e.Range.End.Int()

			if start > end {
				return false
			}

			list, err := util.ListMinute(start, end)

			if err != nil {
				panic(err)
			}

			for _, i := range list {
				if i == minute && i%bottom == start%bottom {
					return true
				}
			}

			return false
		} else {
			var top int

			if e.Wildcard {
				top = 0
			} else {
				top = e.Number.Int()
			}

			if bottom == 0 {
				if e.Wildcard {
					bottom = 1
				} else {
					return e.Number.Match(t)
				}
			}

			return minute >= top && minute%bottom == top%bottom
		}
	} else {
		if e.Wildcard {
			return true
		} else if e.Range != nil {
			return e.Range.Match(t)
		} else if e.Number != nil {
			return e.Number.Match(t)
		}
	}

	panic("must not happen")
}

func (v *MinuteField) Match(t time.Time) bool {
	for _, e := range v.Exps {
		if e.Match(t) {
			return true
		}
	}

	return false
}

// hour =======================================================================

func (v *Hour) Match(t time.Time) bool {
	return v.Int() == t.Hour()
}

func (v *HourRange) Match(t time.Time) bool {
	hour := t.Hour()
	list, err := util.ListHour(v.Start.Int(), v.End.Int())

	if err != nil {
		panic(err)
	}

	for _, i := range list {
		if i == hour {
			return true
		}
	}

	return false
}

func (e *HourExp) Match(t time.Time) bool {
	if e.Bottom != nil {
		hour := t.Hour()
		bottom := *e.Bottom

		if e.Range != nil {
			if bottom == 0 {
				return e.Range.Start.Match(t)
			}

			start := e.Range.Start.Int()
			end := e.Range.End.Int()

			if start > end {
				return false
			}

			list, err := util.ListMinute(start, end)

			if err != nil {
				panic(err)
			}

			for _, i := range list {
				if i == hour && i%bottom == start%bottom {
					return true
				}
			}

			return false
		} else {
			var top int

			if e.Wildcard {
				top = 0
			} else {
				top = e.Number.Int()
			}

			if bottom == 0 {
				if e.Wildcard {
					bottom = 1
				} else {
					return e.Number.Match(t)
				}
			}

			return hour >= top && hour%bottom == top%bottom
		}
	} else {
		if e.Wildcard {
			return true
		} else if e.Range != nil {
			return e.Range.Match(t)
		} else if e.Number != nil {
			return e.Number.Match(t)
		}
	}

	panic("must not happen")
}

func (v *HourField) Match(t time.Time) bool {
	for _, e := range v.Exps {
		if e.Match(t) {
			return true
		}
	}

	return false
}

// day_of_month ===============================================================

func (v *DayOfMonth) Match(t time.Time) bool {
	return v.Int() == t.Day()
}

func (v *DayOfMonthRange) Match(t time.Time) bool {
	day := t.Day()
	list, err := util.ListDayOfMonth(t, v.Start.Int(), v.End.Int())

	if err != nil {
		panic(err)
	}

	for _, i := range list {
		if i == day {
			return true
		}
	}

	return false
}

func (v *NearestWeekday) Match(t time.Time) bool {
	return util.NearestWeekday(t, v.Int()) == t.Day()
}

func (v *LastDayOfMonth) Match(t time.Time) bool {
	return util.LastOfMonth(t)-v.Int() == t.Day()
}

func (e *DayOfMonthExp) Match(t time.Time) bool {
	if e.NearestWeekday != nil {
		return e.NearestWeekday.Match(t)
	} else if e.Last != nil {
		return e.Last.Match(t)
	} else if e.Bottom != nil {
		day := t.Day()
		bottom := *e.Bottom

		if e.Range != nil {
			if bottom == 0 {
				return e.Range.Start.Match(t)
			}

			start := e.Range.Start.Int()
			end := e.Range.End.Int()

			if start > end {
				return false
			}

			list, err := util.ListDayOfMonth(t, start, end)

			if err != nil {
				panic(err)
			}

			for _, i := range list {
				if i == day && i%bottom == start%bottom {
					return true
				}
			}

			return false
		} else {
			var top int

			if e.Wildcard {
				top = 1
			} else {
				top = e.Number.Int()
			}

			if bottom == 0 {
				if e.Wildcard {
					bottom = 1
				} else {
					return e.Number.Match(t)
				}
			}

			return day >= top && day%bottom == top%bottom
		}
	} else {
		if e.Wildcard {
			return true
		} else if e.Range != nil {
			return e.Range.Match(t)
		} else if e.Number != nil {
			return e.Number.Match(t)
		}
	}

	panic("must not happen")
}

func (v *DayOfMonthField) Match(t time.Time) bool {
	if v.Any {
		return true
	}

	for _, e := range v.Exps {
		if e.Match(t) {
			return true
		}
	}

	return false
}

// month ======================================================================

func (v *Month) Match(t time.Time) bool {
	return v.Month() == t.Month()
}

func (v *MonthRange) Match(t time.Time) bool {
	month := t.Month()
	list, err := util.ListMonth(v.Start.Month(), v.End.Month())

	if err != nil {
		panic(err)
	}

	for _, i := range list {
		if i == month {
			return true
		}
	}

	return false
}

func (e *MonthExp) Match(t time.Time) bool {
	if e.Bottom != nil {
		month := t.Month()
		bottom := *e.Bottom

		if e.Range != nil {
			if bottom == 0 {
				return e.Range.Start.Match(t)
			}

			start := e.Range.Start.Month()
			end := e.Range.End.Month()

			if start > end {
				return false
			}

			list, err := util.ListMonth(start, end)

			if err != nil {
				panic(err)
			}

			for _, i := range list {
				if i == month && int(i)%bottom == int(start)%bottom {
					return true
				}
			}

			return false
		} else {
			var top time.Month

			if e.Wildcard {
				top = time.January
			} else {
				top = e.Month.Month()
			}

			if bottom == 0 {
				if e.Wildcard {
					bottom = 1
				} else {
					return e.Month.Match(t)
				}
			}

			return month >= top && int(month)%bottom == int(top)%bottom
		}
	} else {
		if e.Wildcard {
			return true
		} else if e.Range != nil {
			return e.Range.Match(t)
		} else if e.Month != nil {
			return e.Month.Match(t)
		}
	}

	panic("must not happen")
}

func (v *MonthField) Match(t time.Time) bool {
	for _, e := range v.Exps {
		if e.Match(t) {
			return true
		}
	}

	return false
}

// day_of_week ================================================================

func (v *Weekday) Match(t time.Time) bool {
	return v.Weekday() == t.Weekday()
}

func (v *WeekdayRange) Match(t time.Time) bool {
	wday := t.Weekday()
	list, err := util.ListWeekday(v.Start.Weekday(), v.End.Weekday())

	if err != nil {
		panic(err)
	}

	for _, i := range list {
		if i == wday {
			return true
		}
	}

	return false
}

func (v *NthDayOfWeek) Match(t time.Time) bool {
	return util.NthDayOfWeek(t, v.Wday.Weekday(), v.Nth) == t.Day()
}

func (v *LastDayOfWeek) Match(t time.Time) bool {
	if v.Wday == nil {
		// NOTE: If the day of the week is not specified,
		//       it will be the same as when SAT is specified.
		return t.Weekday() == time.Saturday
	} else {
		return util.LastWdayOfMonth(t, v.Weekday()) == t.Day()
	}
}

func (e *DayOfWeekExp) Match(t time.Time) bool {
	if e.Nth != nil {
		return e.Nth.Match(t)
	} else if e.Last != nil {
		return e.Last.Match(t)
	} else if e.Bottom != nil {
		wday := t.Weekday()
		bottom := *e.Bottom

		if e.Range != nil {
			if bottom == 0 {
				return e.Range.Start.Match(t)
			}

			start := e.Range.Start.Weekday()
			end := e.Range.End.Weekday()

			if start > end {
				return false
			}

			list, err := util.ListWeekday(start, end)

			if err != nil {
				panic(err)
			}

			for _, i := range list {
				if i == wday && int(i)%bottom == int(start)%bottom {
					return true
				}
			}

			return false
		} else {
			var top time.Weekday

			if e.Wildcard {
				top = time.Sunday
			} else {
				top = e.Wday.Weekday()
			}

			if bottom == 0 {
				if e.Wildcard {
					bottom = 1
				} else {
					return e.Wday.Match(t)
				}
			}

			return wday >= top && int(wday)%bottom == int(top)%bottom
		}
	} else {
		if e.Wildcard {
			return true
		} else if e.Range != nil {
			return e.Range.Match(t)
		} else if e.Wday != nil {
			return e.Wday.Match(t)
		}
	}

	panic("must not happen")
}

func (v *DayOfWeekField) Match(t time.Time) bool {
	if v.Any {
		return true
	}

	for _, e := range v.Exps {
		if e.Match(t) {
			return true
		}
	}

	return false
}

// year =======================================================================

func (v *Year) Match(t time.Time) bool {
	return v.Int() == t.Year()
}

func (v *YearRange) Match(t time.Time) bool {
	year := t.Year()
	list, err := util.ListYear(v.Start.Int(), v.End.Int())

	if err != nil {
		panic(err)
	}

	for _, i := range list {
		if i == year {
			return true
		}
	}

	return false
}

func (e *YearExp) Match(t time.Time) bool {
	if e.Bottom != nil {
		year := t.Year()
		bottom := *e.Bottom

		if e.Range != nil {
			if bottom == 0 {
				return e.Range.Start.Match(t)
			}

			start := e.Range.Start.Int()
			end := e.Range.End.Int()

			if start > end {
				return false
			}

			list, err := util.ListYear(start, end)

			if err != nil {
				panic(err)
			}

			for _, i := range list {
				if i == year && i%bottom == start%bottom {
					return true
				}
			}

			return false
		} else {
			var top int

			if e.Wildcard {
				top = 1970
			} else {
				top = e.Number.Int()
			}

			if bottom == 0 {
				if e.Wildcard {
					bottom = 1
				} else {
					return e.Number.Match(t)
				}
			}

			return year >= top && year%bottom == top%bottom
		}
	} else {
		if e.Wildcard {
			return true
		} else if e.Range != nil {
			return e.Range.Match(t)
		} else if e.Number != nil {
			return e.Number.Match(t)
		}
	}

	panic("must not happen")
}

func (v *YearField) Match(t time.Time) bool {
	for _, e := range v.Exps {
		if e.Match(t) {
			return true
		}
	}

	return false
}

// expression =================================================================

func (v *Expression) Match(t time.Time) bool {
	return v.Minute.Match(t) &&
		v.Hour.Match(t) &&
		v.DayOfMonth.Match(t) &&
		v.Month.Match(t) &&
		v.DayOfWeek.Match(t) &&
		v.Year.Match(t)
}

package cronplan_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/cronplan"
)

func TestMinutToString(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("1", minute(1).String())
}

func TestHourToString(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("1", hour(1).String())
}

func TestDayOfMonthToString(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("1", day(1).String())
}

func TestYearToString(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("1", year(1).String())
}

func TestWeekNameToString(t *testing.T) {
	assert := assert.New(t)

	tt := map[time.Weekday]string{
		time.Sunday:    "SUN",
		time.Monday:    "MON",
		time.Tuesday:   "TUE",
		time.Wednesday: "WED",
		time.Thursday:  "THU",
		time.Friday:    "FRI",
		time.Saturday:  "SAT",
	}

	for k, v := range tt {
		assert.Equal(v, wday(k).String())
	}
}

func TestMonthNameToString(t *testing.T) {
	assert := assert.New(t)

	tt := map[time.Month]string{
		time.January:   "JAN",
		time.February:  "FEB",
		time.March:     "MAR",
		time.April:     "APR",
		time.May:       "MAY",
		time.June:      "JUN",
		time.July:      "JUL",
		time.August:    "AUG",
		time.September: "SEP",
		time.October:   "OCT",
		time.November:  "NOV",
		time.December:  "DEC",
	}

	for k, v := range tt {
		assert.Equal(v, month(k).String())
	}
}

func TestMinuteRangeToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronplan.MinuteRange{
		Start: minute(0),
		End:   minute(59),
	}
	assert.Equal("0-59", x.String())
}

func TestHourangeToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronplan.HourRange{
		Start: hour(0),
		End:   hour(23),
	}
	assert.Equal("0-23", x.String())
}

func TestYearRangeToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronplan.YearRange{
		Start: year(1970),
		End:   year(2199),
	}
	assert.Equal("1970-2199", x.String())
}

func TestWeekRangeToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronplan.WeekdayRange{
		Start: wday(time.Sunday),
		End:   wday(time.Saturday),
	}
	assert.Equal("SUN-SAT", x.String())
}

func TestMonthRangeToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronplan.MonthRange{
		Start: month(time.January),
		End:   month(time.December),
	}
	assert.Equal("JAN-DEC", x.String())
}

func TestMinuteAllToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronplan.MinuteExp{Wildcard: true}
	assert.Equal("*", x.String())
}

func TestMinuteIncrToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronplan.MinuteExp{Number: minute(0), Bottom: intptr(5)}
	assert.Equal("0/5", x.String())
}

func TestMinuteIncrWildcardToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronplan.MinuteExp{Wildcard: true, Bottom: intptr(5)}
	assert.Equal("*/5", x.String())
}

func TestMinuteIncrWRangeToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronplan.MinuteExp{Range: &cronplan.MinuteRange{
		Start: minute(0),
		End:   minute(59),
	}, Bottom: intptr(5)}
	assert.Equal("0-59/5", x.String())
}

func TestMinutIncrWildcardToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronplan.MinuteExp{
		Wildcard: true,
		Bottom:   intptr(5),
	}
	assert.Equal("*/5", x.String())
}

func TestHourIncrRangeToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronplan.HourExp{
		Range: &cronplan.HourRange{
			Start: hour(0),
			End:   hour(5),
		},
		Bottom: intptr(5),
	}
	assert.Equal("0-5/5", x.String())
}

func TestDayOfMonthIncrRangeToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronplan.DayOfMonthExp{
		Range: &cronplan.DayOfMonthRange{
			Start: day(0),
			End:   day(5),
		},
		Bottom: intptr(5),
	}
	assert.Equal("0-5/5", x.String())
}

func TestYearIncrRangeToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronplan.YearExp{
		Range: &cronplan.YearRange{
			Start: year(2000),
			End:   year(2020),
		},
		Bottom: intptr(5),
	}
	assert.Equal("2000-2020/5", x.String())
}

func TestDayOfMonthAnyToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronplan.DayOfMonthField{Any: true}
	assert.Equal("?", x.String())
}

func TestDayOfWeekAnyToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronplan.DayOfWeekField{Any: true}
	assert.Equal("?", x.String())
}

func TestLastOfMonthToString1(t *testing.T) {
	assert := assert.New(t)
	x := cronplan.LastDayOfMonth(3)
	assert.Equal("L-3", x.String())
}

func TestLastOfMonthToString2(t *testing.T) {
	assert := assert.New(t)
	x := cronplan.LastDayOfMonth(0)
	assert.Equal("L", x.String())
}

func TestLastOfWeekToString(t *testing.T) {
	assert := assert.New(t)
	x := cronplan.LastDayOfWeek(time.Saturday)
	assert.Equal("L", x.String())
}

func TestNearestWeekdayToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronplan.DayOfMonthExp{
		NearestWeekday: nwday(3),
	}
	assert.Equal("3W", x.String())
}

func TestInstanceToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronplan.DayOfWeekExp{
		Nth: &cronplan.NthDayOfWeek{Wday: wday(time.Friday), Nth: 3},
	}
	assert.Equal("FRI#3", x.String())
}

func TestMinuteExpToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronplan.MinuteExp{
		Number: minute(1),
	}
	assert.Equal("1", x.String())
}

func TestMinutesToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronplan.MinuteField{
		Exps: []*cronplan.MinuteExp{
			{Number: minute(1)},
			{Number: minute(2)},
		},
	}
	assert.Equal("1,2", x.String())
}

func TestHourExpToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronplan.HourExp{
		Number: hour(1),
	}
	assert.Equal("1", x.String())
}

func TestHoursToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronplan.HourField{
		Exps: []*cronplan.HourExp{
			{Number: hour(1)},
			{Number: hour(2)},
		},
	}
	assert.Equal("1,2", x.String())
}

func TestDayOfMonthExpToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronplan.DayOfMonthExp{
		Number: day(1),
	}
	assert.Equal("1", x.String())
}

func TestDayOfMonthsToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronplan.DayOfMonthField{
		Exps: []*cronplan.DayOfMonthExp{
			{Number: day(1)},
			{Number: day(2)},
		},
	}
	assert.Equal("1,2", x.String())
}

func TestMonthExpToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronplan.MonthExp{
		Month: month(1),
	}
	assert.Equal("JAN", x.String())
}

func TestMonthsToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronplan.MonthField{
		Exps: []*cronplan.MonthExp{
			{Month: month(time.January)},
			{Month: month(time.February)},
		},
	}
	assert.Equal("JAN,FEB", x.String())
}

func TestDayOfWeekExpToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronplan.DayOfWeekExp{
		Wday: wday(time.Sunday),
	}
	assert.Equal("SUN", x.String())
}

func TestDayOfWeeksToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronplan.DayOfWeekField{
		Exps: []*cronplan.DayOfWeekExp{
			{Wday: wday(time.Weekday(time.Sunday))},
			{Wday: wday(time.Weekday(time.Monday))},
		},
	}
	assert.Equal("SUN,MON", x.String())
}

func TestYearExpToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronplan.YearExp{
		Number: year(2000),
	}
	assert.Equal("2000", x.String())
}

func TestYearExpFieldToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronplan.YearField{
		Exps: []*cronplan.YearExp{
			{Number: year(2000)},
			{Number: year(2020)},
		},
	}
	assert.Equal("2000,2020", x.String())
}

func TestExpressionToString(t *testing.T) {
	assert := assert.New(t)
	x := &cronplan.Expression{
		Minute: &cronplan.MinuteField{
			Exps: []*cronplan.MinuteExp{{Number: minute(0)}},
		},
		Hour: &cronplan.HourField{
			Exps: []*cronplan.HourExp{{Number: hour(10)}},
		},
		DayOfMonth: &cronplan.DayOfMonthField{
			Exps: []*cronplan.DayOfMonthExp{{
				Wildcard: true,
			}},
		},
		Month: &cronplan.MonthField{
			Exps: []*cronplan.MonthExp{{
				Wildcard: true,
			}},
		},
		DayOfWeek: &cronplan.DayOfWeekField{
			Any: true,
		},
		Year: &cronplan.YearField{
			Exps: []*cronplan.YearExp{{
				Wildcard: true,
			}},
		},
	}

	assert.Equal("0 10 * * ? *", x.String())
}

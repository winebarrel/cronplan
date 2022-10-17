package cronplan_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/cronplan"
)

func intptr(i int) *int {
	return &i
}

func minute(i int) *cronplan.Minute {
	v := cronplan.Minute(i)
	return &v
}

func hour(i int) *cronplan.Hour {
	v := cronplan.Hour(i)
	return &v
}

func wday(i time.Weekday) *cronplan.Weekday {
	v := cronplan.Weekday(i)
	return &v
}

func day(i int) *cronplan.DayOfMonth {
	v := cronplan.DayOfMonth(i)
	return &v
}

func month(i time.Month) *cronplan.Month {
	v := cronplan.Month(i)
	return &v
}

func year(i int) *cronplan.Year {
	v := cronplan.Year(i)
	return &v
}

func last(i int) *cronplan.LastDayOfMonth {
	v := cronplan.LastDayOfMonth(i)
	return &v
}

func nwday(i int) *cronplan.NearestWeekday {
	v := cronplan.NearestWeekday(i)
	return &v
}

func lastw(i time.Weekday) *cronplan.LastDayOfWeek {
	v := cronplan.LastDayOfWeek(i)
	return &v
}

func TestIntegration(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		exp    string
		extStr string
		ast    *cronplan.Expression
	}{
		// https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html
		{
			"0 10 * * ? *",
			"0 10 * * ? *",
			&cronplan.Expression{
				Minute: &cronplan.MinuteField{
					Exps: []*cronplan.MinuteExp{{
						Number: minute(0),
					}},
				},
				Hour: &cronplan.HourField{
					Exps: []*cronplan.HourExp{{
						Number: hour(10),
					}},
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
			},
		},
		{
			"15 12 * * ? *",
			"15 12 * * ? *",
			&cronplan.Expression{
				Minute: &cronplan.MinuteField{
					Exps: []*cronplan.MinuteExp{{
						Number: minute(15),
					}},
				},
				Hour: &cronplan.HourField{
					Exps: []*cronplan.HourExp{{
						Number: hour(12),
					}},
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
			},
		},
		{
			"0 18 ? * MON-FRI *",
			"0 18 ? * MON-FRI *",
			&cronplan.Expression{
				Minute: &cronplan.MinuteField{
					Exps: []*cronplan.MinuteExp{{
						Number: minute(0),
					}},
				},
				Hour: &cronplan.HourField{
					Exps: []*cronplan.HourExp{{
						Number: hour(18),
					}},
				},
				DayOfMonth: &cronplan.DayOfMonthField{
					Any: true,
				},
				Month: &cronplan.MonthField{
					Exps: []*cronplan.MonthExp{{
						Wildcard: true,
					}},
				},
				DayOfWeek: &cronplan.DayOfWeekField{
					Exps: []*cronplan.DayOfWeekExp{{
						Range: &cronplan.WeekdayRange{
							Start: wday(time.Monday),
							End:   wday(time.Friday),
						},
					}},
				},
				Year: &cronplan.YearField{
					Exps: []*cronplan.YearExp{{
						Wildcard: true,
					}},
				},
			},
		},
		{
			"0 8 1 * ? *",
			"0 8 1 * ? *",
			&cronplan.Expression{
				Minute: &cronplan.MinuteField{
					Exps: []*cronplan.MinuteExp{{
						Number: minute(0),
					}},
				},
				Hour: &cronplan.HourField{
					Exps: []*cronplan.HourExp{{
						Number: hour(8),
					}},
				},
				DayOfMonth: &cronplan.DayOfMonthField{
					Exps: []*cronplan.DayOfMonthExp{{
						Number: day(1),
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
			},
		},
		{
			"0/15 * * * ? *",
			"0/15 * * * ? *",
			&cronplan.Expression{
				Minute: &cronplan.MinuteField{
					Exps: []*cronplan.MinuteExp{{
						Number: minute(0),
						Bottom: intptr(15),
					}},
				},
				Hour: &cronplan.HourField{
					Exps: []*cronplan.HourExp{{
						Wildcard: true,
					}},
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
			},
		},
		{
			"0/10 * ? * MON-FRI *",
			"0/10 * ? * MON-FRI *",
			&cronplan.Expression{
				Minute: &cronplan.MinuteField{
					Exps: []*cronplan.MinuteExp{{
						Number: minute(0),
						Bottom: intptr(10),
					}},
				},
				Hour: &cronplan.HourField{
					Exps: []*cronplan.HourExp{{
						Wildcard: true,
					}},
				},
				DayOfMonth: &cronplan.DayOfMonthField{
					Any: true,
				},
				Month: &cronplan.MonthField{
					Exps: []*cronplan.MonthExp{{
						Wildcard: true,
					}},
				},
				DayOfWeek: &cronplan.DayOfWeekField{
					Exps: []*cronplan.DayOfWeekExp{{
						Range: &cronplan.WeekdayRange{
							Start: wday(time.Monday),
							End:   wday(time.Friday),
						},
					}},
				},
				Year: &cronplan.YearField{
					Exps: []*cronplan.YearExp{{
						Wildcard: true,
					}},
				},
			},
		},
		{
			"0/5 8-17 ? * MON-FRI *",
			"0/5 8-17 ? * MON-FRI *",
			&cronplan.Expression{
				Minute: &cronplan.MinuteField{
					Exps: []*cronplan.MinuteExp{{
						Number: minute(0),
						Bottom: intptr(5),
					}},
				},
				Hour: &cronplan.HourField{
					Exps: []*cronplan.HourExp{{
						Range: &cronplan.HourRange{
							Start: hour(8),
							End:   hour(17),
						},
					}},
				},
				DayOfMonth: &cronplan.DayOfMonthField{
					Any: true,
				},
				Month: &cronplan.MonthField{
					Exps: []*cronplan.MonthExp{{
						Wildcard: true,
					}},
				},
				DayOfWeek: &cronplan.DayOfWeekField{
					Exps: []*cronplan.DayOfWeekExp{{
						Range: &cronplan.WeekdayRange{
							Start: wday(time.Monday),
							End:   wday(time.Friday),
						},
					}},
				},
				Year: &cronplan.YearField{
					Exps: []*cronplan.YearExp{{
						Wildcard: true,
					}},
				},
			},
		},
		// https://docs.oracle.com/cd/E12058_01/doc/doc.1014/e12030/cron_expression.htm
		{
			"0 12 * * ? *",
			"0 12 * * ? *",
			&cronplan.Expression{
				Minute: &cronplan.MinuteField{
					Exps: []*cronplan.MinuteExp{{
						Number: minute(0),
					}},
				},
				Hour: &cronplan.HourField{
					Exps: []*cronplan.HourExp{{
						Number: hour(12),
					}},
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
			},
		},
		{
			"15 10 ? * * *",
			"15 10 ? * * *",
			&cronplan.Expression{
				Minute: &cronplan.MinuteField{
					Exps: []*cronplan.MinuteExp{{
						Number: minute(15),
					}},
				},
				Hour: &cronplan.HourField{
					Exps: []*cronplan.HourExp{{
						Number: hour(10),
					}},
				},
				DayOfMonth: &cronplan.DayOfMonthField{
					Any: true,
				},
				Month: &cronplan.MonthField{
					Exps: []*cronplan.MonthExp{{
						Wildcard: true,
					}},
				},
				DayOfWeek: &cronplan.DayOfWeekField{
					Exps: []*cronplan.DayOfWeekExp{{
						Wildcard: true,
					}},
				},
				Year: &cronplan.YearField{
					Exps: []*cronplan.YearExp{{
						Wildcard: true,
					}},
				},
			},
		},
		{
			"15 10 * * ? *",
			"15 10 * * ? *",
			&cronplan.Expression{
				Minute: &cronplan.MinuteField{
					Exps: []*cronplan.MinuteExp{{
						Number: minute(15),
					}},
				},
				Hour: &cronplan.HourField{
					Exps: []*cronplan.HourExp{{
						Number: hour(10),
					}},
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
			},
		},
		{
			"15 10 * * ? 2005",
			"15 10 * * ? 2005",
			&cronplan.Expression{
				Minute: &cronplan.MinuteField{
					Exps: []*cronplan.MinuteExp{{
						Number: minute(15),
					}},
				},
				Hour: &cronplan.HourField{
					Exps: []*cronplan.HourExp{{
						Number: hour(10),
					}},
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
						Number: year(2005),
					}},
				},
			},
		},
		{
			"* 14 * * ? *",
			"* 14 * * ? *",
			&cronplan.Expression{
				Minute: &cronplan.MinuteField{
					Exps: []*cronplan.MinuteExp{{
						Wildcard: true,
					}},
				},
				Hour: &cronplan.HourField{
					Exps: []*cronplan.HourExp{{
						Number: hour(14),
					}},
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
			},
		},
		{
			"0/5 14 * * ? *",
			"0/5 14 * * ? *",
			&cronplan.Expression{
				Minute: &cronplan.MinuteField{
					Exps: []*cronplan.MinuteExp{{
						Number: minute(0),
						Bottom: intptr(5),
					}},
				},
				Hour: &cronplan.HourField{
					Exps: []*cronplan.HourExp{{
						Number: hour(14),
					}},
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
			},
		},
		{
			"0/5 14,18 * * ? *",
			"0/5 14,18 * * ? *",
			&cronplan.Expression{
				Minute: &cronplan.MinuteField{
					Exps: []*cronplan.MinuteExp{{
						Number: minute(0),
						Bottom: intptr(5),
					}},
				},
				Hour: &cronplan.HourField{
					Exps: []*cronplan.HourExp{
						{Number: hour(14)},
						{Number: hour(18)},
					},
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
			},
		},
		{
			"0-5 14 * * ? *",
			"0-5 14 * * ? *",
			&cronplan.Expression{
				Minute: &cronplan.MinuteField{
					Exps: []*cronplan.MinuteExp{{
						Range: &cronplan.MinuteRange{
							Start: minute(0),
							End:   minute(5),
						},
					}},
				},
				Hour: &cronplan.HourField{
					Exps: []*cronplan.HourExp{{
						Number: hour(14),
					}},
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
			},
		},
		{
			"10,44 14 ? 3 WED *",
			"10,44 14 ? MAR WED *",
			&cronplan.Expression{
				Minute: &cronplan.MinuteField{
					Exps: []*cronplan.MinuteExp{
						{Number: minute(10)},
						{Number: minute(44)},
					},
				},
				Hour: &cronplan.HourField{
					Exps: []*cronplan.HourExp{{
						Number: hour(14),
					}},
				},
				DayOfMonth: &cronplan.DayOfMonthField{
					Any: true,
				},
				Month: &cronplan.MonthField{
					Exps: []*cronplan.MonthExp{{
						Month: month(time.March),
					}},
				},
				DayOfWeek: &cronplan.DayOfWeekField{
					Exps: []*cronplan.DayOfWeekExp{{
						Wday: wday(time.Wednesday),
					}},
				},
				Year: &cronplan.YearField{
					Exps: []*cronplan.YearExp{{
						Wildcard: true,
					}},
				},
			},
		},
		{
			"15 10 ? * MON-FRI *",
			"15 10 ? * MON-FRI *",
			&cronplan.Expression{
				Minute: &cronplan.MinuteField{
					Exps: []*cronplan.MinuteExp{{
						Number: minute(15),
					}},
				},
				Hour: &cronplan.HourField{
					Exps: []*cronplan.HourExp{{
						Number: hour(10),
					}},
				},
				DayOfMonth: &cronplan.DayOfMonthField{
					Any: true,
				},
				Month: &cronplan.MonthField{
					Exps: []*cronplan.MonthExp{{
						Wildcard: true,
					}},
				},
				DayOfWeek: &cronplan.DayOfWeekField{
					Exps: []*cronplan.DayOfWeekExp{{
						Range: &cronplan.WeekdayRange{
							Start: wday(time.Monday),
							End:   wday(time.Friday),
						},
					}},
				},
				Year: &cronplan.YearField{
					Exps: []*cronplan.YearExp{{
						Wildcard: true,
					}},
				},
			},
		},
		{
			"15 10 15 * ? *",
			"15 10 15 * ? *",
			&cronplan.Expression{
				Minute: &cronplan.MinuteField{
					Exps: []*cronplan.MinuteExp{{
						Number: minute(15),
					}},
				},
				Hour: &cronplan.HourField{
					Exps: []*cronplan.HourExp{{
						Number: hour(10),
					}},
				},
				DayOfMonth: &cronplan.DayOfMonthField{
					Exps: []*cronplan.DayOfMonthExp{{
						Number: day(15),
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
			},
		},
		{
			"15 10 L * ? *",
			"15 10 L * ? *",
			&cronplan.Expression{
				Minute: &cronplan.MinuteField{
					Exps: []*cronplan.MinuteExp{{
						Number: minute(15),
					}},
				},
				Hour: &cronplan.HourField{
					Exps: []*cronplan.HourExp{{
						Number: hour(10),
					}},
				},
				DayOfMonth: &cronplan.DayOfMonthField{
					Exps: []*cronplan.DayOfMonthExp{{
						Last: last(0),
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
			},
		},
		{
			"15 10 ? * 6#3 *",
			"15 10 ? * FRI#3 *",
			&cronplan.Expression{
				Minute: &cronplan.MinuteField{
					Exps: []*cronplan.MinuteExp{{
						Number: minute(15),
					}},
				},
				Hour: &cronplan.HourField{
					Exps: []*cronplan.HourExp{{
						Number: hour(10),
					}},
				},
				DayOfMonth: &cronplan.DayOfMonthField{
					Any: true,
				},
				Month: &cronplan.MonthField{
					Exps: []*cronplan.MonthExp{{
						Wildcard: true,
					}},
				},
				DayOfWeek: &cronplan.DayOfWeekField{
					Exps: []*cronplan.DayOfWeekExp{{
						Nth: &cronplan.NthDayOfWeek{
							Wday: wday(time.Friday),
							Nth:  3,
						},
					}},
				},
				Year: &cronplan.YearField{
					Exps: []*cronplan.YearExp{{
						Wildcard: true,
					}},
				},
			},
		},
		{
			"0 12 1/5 * ? *",
			"0 12 1/5 * ? *",
			&cronplan.Expression{
				Minute: &cronplan.MinuteField{
					Exps: []*cronplan.MinuteExp{{
						Number: minute(0),
					}},
				},
				Hour: &cronplan.HourField{
					Exps: []*cronplan.HourExp{{
						Number: hour(12),
					}},
				},
				DayOfMonth: &cronplan.DayOfMonthField{
					Exps: []*cronplan.DayOfMonthExp{{
						Number: day(1),
						Bottom: intptr(5),
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
			},
		},
		{
			"11 11 11 11 ? *",
			"11 11 11 NOV ? *",
			&cronplan.Expression{
				Minute: &cronplan.MinuteField{
					Exps: []*cronplan.MinuteExp{{
						Number: minute(11),
					}},
				},
				Hour: &cronplan.HourField{
					Exps: []*cronplan.HourExp{{
						Number: hour(11),
					}},
				},
				DayOfMonth: &cronplan.DayOfMonthField{
					Exps: []*cronplan.DayOfMonthExp{{
						Number: day(11),
					}},
				},
				Month: &cronplan.MonthField{
					Exps: []*cronplan.MonthExp{{
						Month: month(time.November),
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
			},
		},
	}

	for _, t := range tt {
		cron, err := cronplan.Parse(t.exp)
		assert.NoError(err)
		assert.Equal(cron, t.ast, t.exp)
		assert.Equal(t.extStr, t.ast.String(), t.exp)
	}
}

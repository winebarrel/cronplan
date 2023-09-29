package cronplan_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/cronplan"
)

func TestDayOfWeekAll(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * ? * * *")
	assert.NoError(err)
	assert.True(cron.DayOfWeek.Exps[0].Wildcard)
}

func TestDayOfWeekNumber(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * ? * 1 *")
	assert.NoError(err)
	assert.Equal(wday(time.Sunday), cron.DayOfWeek.Exps[0].Wday)
}

func TestDayOfWeekNumberWithZero(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * ? * 01 *")
	assert.NoError(err)
	assert.Equal(wday(time.Sunday), cron.DayOfWeek.Exps[0].Wday)
}

func TestDayOfWeekNumberRange(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * ? * 1-7 *")
	assert.NoError(err)
	assert.Equal(&cronplan.WeekdayRange{
		Start: wday(time.Sunday),
		End:   wday(time.Saturday),
	}, cron.DayOfWeek.Exps[0].Range)
}

func TestDayOfWeekIncr(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * ? * 1/5 *")
	assert.NoError(err)
	assert.Equal(&cronplan.DayOfWeekExp{
		Wday:   wday(time.Sunday),
		Bottom: intptr(5),
	}, cron.DayOfWeek.Exps[0])
}

func TestDayOfWeekIncrWildcard(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * ? * */5 *")
	assert.NoError(err)
	assert.Equal(&cronplan.DayOfWeekExp{
		Wildcard: true,
		Bottom:   intptr(5),
	}, cron.DayOfWeek.Exps[0])
}

func TestDayOfWeekIncrRange(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * ? * 1-5/3 *")
	assert.NoError(err)
	assert.Equal(&cronplan.DayOfWeekExp{
		Range: &cronplan.WeekdayRange{
			Start: wday(time.Sunday),
			End:   wday(time.Thursday),
		},
		Bottom: intptr(3),
	}, cron.DayOfWeek.Exps[0])
}

func TestDayOfWeekAny(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * * * ? *")
	assert.NoError(err)
	assert.True(cron.DayOfWeek.Any)
}

func TestDayOfWeekLast(t *testing.T) {
	tt := []time.Weekday{
		time.Sunday,
		time.Monday,
		time.Tuesday,
		time.Wednesday,
		time.Thursday,
		time.Friday,
		time.Saturday,
	}

	assert := assert.New(t)

	for i, w := range tt {
		cron, err := cronplan.Parse(fmt.Sprintf("* * ? * %dL *", i+1))
		assert.NoError(err)
		assert.Equal(lastw(w), cron.DayOfWeek.Exps[0].Last)
	}
}

func TestDayOfWeekLastWithoutWday(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * ? * L *")
	assert.NoError(err)
	assert.Equal(&cronplan.LastDayOfWeek{}, cron.DayOfWeek.Exps[0].Last)
}

func TestDayOfWeekName(t *testing.T) {
	assert := assert.New(t)
	tt := map[string]time.Weekday{
		"SUN": time.Sunday,
		"MON": time.Monday,
		"TUE": time.Tuesday,
		"WED": time.Wednesday,
		"THU": time.Thursday,
		"FRI": time.Friday,
		"SAT": time.Saturday,
	}

	for k, v := range tt {
		cron, err := cronplan.Parse(fmt.Sprintf("* * ? * %s *", k))
		assert.NoError(err)
		assert.Equal(wday(v), cron.DayOfWeek.Exps[0].Wday, k)

		cron, err = cronplan.Parse(fmt.Sprintf("* * ? * %s *", strings.ToLower(k)))
		assert.NoError(err)
		assert.Equal(wday(v), cron.DayOfWeek.Exps[0].Wday, k)
	}
}

func TestDayOfWeekNameRange(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * ? * SUN-SAT *")
	assert.NoError(err)
	assert.Equal(&cronplan.WeekdayRange{
		Start: wday(time.Sunday),
		End:   wday(time.Saturday),
	}, cron.DayOfWeek.Exps[0].Range)
}

func TestDayOfWeekNameIncr(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * ? * SUN/3 *")
	assert.NoError(err)
	assert.Equal(
		&cronplan.DayOfWeekExp{
			Wday:   wday(time.Sunday),
			Bottom: intptr(3),
		}, cron.DayOfWeek.Exps[0])
}

func TestDayOfWeekNameIncrRange(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * ? * SUN-SAT/3 *")
	assert.NoError(err)
	assert.Equal(
		&cronplan.DayOfWeekExp{
			Range: &cronplan.WeekdayRange{
				Start: wday(time.Sunday),
				End:   wday(time.Saturday),
			},
			Bottom: intptr(3),
		}, cron.DayOfWeek.Exps[0])
}

func TestDayOfWeekComplex(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * ? * *,1,1-7,1/5,*/5,L,5L,MONL,SUN,SUN-SAT *")
	assert.NoError(err)
	assert.True(cron.DayOfWeek.Exps[0].Wildcard)
	assert.Equal(wday(time.Sunday), cron.DayOfWeek.Exps[1].Wday)
	assert.Equal(&cronplan.WeekdayRange{
		Start: wday(time.Sunday),
		End:   wday(time.Saturday),
	}, cron.DayOfWeek.Exps[2].Range)
	assert.Equal(&cronplan.DayOfWeekExp{
		Wday:   wday(time.Sunday),
		Bottom: intptr(5),
	}, cron.DayOfWeek.Exps[3])
	assert.Equal(&cronplan.DayOfWeekExp{
		Wildcard: true,
		Bottom:   intptr(5),
	}, cron.DayOfWeek.Exps[4])
	assert.Equal(&cronplan.LastDayOfWeek{}, cron.DayOfWeek.Exps[5].Last)
	assert.Equal(lastw(time.Thursday), cron.DayOfWeek.Exps[6].Last)
	assert.Equal(lastw(time.Monday), cron.DayOfWeek.Exps[7].Last)
	assert.Equal(wday(time.Sunday), cron.DayOfWeek.Exps[8].Wday)
	assert.Equal(&cronplan.WeekdayRange{
		Start: wday(time.Sunday),
		End:   wday(time.Saturday),
	}, cron.DayOfWeek.Exps[9].Range)
}

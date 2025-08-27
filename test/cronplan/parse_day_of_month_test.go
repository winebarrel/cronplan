package cronplan_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/cronplan/v2"
)

func TestDayOfMonthAll(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * * * ? *")
	assert.NoError(err)
	assert.True(cron.DayOfMonth.Exps[0].Wildcard)
}

func TestDayOfMonthNumber(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * 1 * ? *")
	assert.NoError(err)
	assert.Equal(day(1), cron.DayOfMonth.Exps[0].Number)
}

func TestDayOfMonthNumberWithZero(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * 01 * ? *")
	assert.NoError(err)
	assert.Equal(day(1), cron.DayOfMonth.Exps[0].Number)
}

func TestDayOfMonthNumberRange(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * 1-30 * ? *")
	assert.NoError(err)
	assert.Equal(&cronplan.DayOfMonthRange{
		Start: day(1),
		End:   day(30),
	}, cron.DayOfMonth.Exps[0].Range)
}

func TestDayOfMonthIncr(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * 1/5 * ? *")
	assert.NoError(err)
	assert.Equal(&cronplan.DayOfMonthExp{
		Number: day(1),
		Bottom: intptr(5),
	}, cron.DayOfMonth.Exps[0])
}

func TestDayOfMonthIncrWildcard(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * */5 * ? *")
	assert.NoError(err)
	assert.Equal(&cronplan.DayOfMonthExp{
		Wildcard: true,
		Bottom:   intptr(5),
	}, cron.DayOfMonth.Exps[0])
}

func TestDayOfMonthIncrRange(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * 1-10/5 * ? *")
	assert.NoError(err)
	assert.Equal(&cronplan.DayOfMonthExp{
		Range: &cronplan.DayOfMonthRange{
			Start: day(1),
			End:   day(10),
		},
		Bottom: intptr(5),
	}, cron.DayOfMonth.Exps[0])
}

func TestDayOfMonthAny(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * ? * * *")
	assert.NoError(err)
	assert.True(cron.DayOfMonth.Any)
}

func TestDayOfMonthLast(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * L * ? *")
	assert.NoError(err)
	assert.Equal(last(0), cron.DayOfMonth.Exps[0].Last)
}

func TestDayOfMonthLastOffset(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * L-3 * ? *")
	assert.NoError(err)
	assert.Equal(last(3), cron.DayOfMonth.Exps[0].Last)
}

func TestDayOfMonthLastOffsetWithZero(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * L-03 * ? *")
	assert.NoError(err)
	assert.Equal(last(3), cron.DayOfMonth.Exps[0].Last)
}

func TestDayOfMonthWeekday(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * 3W * ? *")
	assert.NoError(err)
	assert.Equal(nwday(3), cron.DayOfMonth.Exps[0].NearestWeekday)
}

func TestDayOfMonthLastWeekday(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * LW * ? *")
	assert.NoError(err)
	assert.Equal(&cronplan.LastWeekdayOfMonth{}, cron.DayOfMonth.Exps[0].LastWeekday)
}

func TestDayOfMonthComplex(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * *,1,1-30,1/5,*/5,L,3W,LW * ? *")
	assert.NoError(err)
	assert.True(cron.DayOfMonth.Exps[0].Wildcard)
	assert.Equal(day(1), cron.DayOfMonth.Exps[1].Number)
	assert.Equal(&cronplan.DayOfMonthRange{
		Start: day(1),
		End:   day(30),
	}, cron.DayOfMonth.Exps[2].Range)
	assert.Equal(&cronplan.DayOfMonthExp{
		Number: day(1),
		Bottom: intptr(5),
	}, cron.DayOfMonth.Exps[3])
	assert.Equal(&cronplan.DayOfMonthExp{
		Wildcard: true,
		Bottom:   intptr(5),
	}, cron.DayOfMonth.Exps[4])
	assert.Equal(last(0), cron.DayOfMonth.Exps[5].Last)
	assert.Equal(nwday(3), cron.DayOfMonth.Exps[6].NearestWeekday)
	assert.Equal(&cronplan.LastWeekdayOfMonth{}, cron.DayOfMonth.Exps[7].LastWeekday)
}

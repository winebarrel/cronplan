package cronplan_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/cronplan"
)

func TestMinutesAll(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * * * ? *")
	assert.NoError(err)
	assert.True(cron.Minute.Exps[0].Wildcard)
}

func TestMinutesNumber(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("0 * * * ? *")
	assert.NoError(err)
	assert.Equal(minute(0), cron.Minute.Exps[0].Number)
}

func TestMinutesNumberWithZero(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("01 * * * ? *")
	assert.NoError(err)
	assert.Equal(minute(1), cron.Minute.Exps[0].Number)
}

func TestMinutesNumberRange(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("0-59 * * * ? *")
	assert.NoError(err)
	assert.Equal(&cronplan.MinuteRange{
		Start: minute(0),
		End:   minute(59),
	}, cron.Minute.Exps[0].Range)
}

func TestMinutesIncr(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("0/5 * * * ? *")
	assert.NoError(err)
	assert.Equal(&cronplan.MinuteExp{
		Number: minute(0),
		Bottom: intptr(5),
	}, cron.Minute.Exps[0])
}

func TestMinutesIncrWildcard(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("*/5 * * * ? *")
	assert.NoError(err)
	assert.Equal(&cronplan.MinuteExp{
		Wildcard: true,
		Bottom:   intptr(5),
	}, cron.Minute.Exps[0])
}

func TestMinutesIncrRange(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("1-10/3 * * * ? *")
	assert.NoError(err)
	assert.Equal(&cronplan.MinuteExp{
		Range: &cronplan.MinuteRange{
			Start: minute(1),
			End:   minute(10),
		},
		Bottom: intptr(3),
	}, cron.Minute.Exps[0])
}

func TestMinutesComplex(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("*,0,0-59,0/5,*/5 * * * ? *")
	assert.NoError(err)
	assert.True(cron.Minute.Exps[0].Wildcard)
	assert.Equal(minute(0), cron.Minute.Exps[1].Number)
	assert.Equal(&cronplan.MinuteRange{
		Start: minute(0),
		End:   minute(59),
	}, cron.Minute.Exps[2].Range)
	assert.Equal(&cronplan.MinuteExp{
		Number: minute(0),
		Bottom: intptr(5),
	}, cron.Minute.Exps[3])
	assert.Equal(&cronplan.MinuteExp{
		Wildcard: true,
		Bottom:   intptr(5),
	}, cron.Minute.Exps[4])
}

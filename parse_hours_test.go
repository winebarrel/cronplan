package cronplan_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/cronplan"
)

func TestHoursAll(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * * * ? *")
	assert.NoError(err)
	assert.True(cron.Hour.Exps[0].Wildcard)
}

func TestHoursNumber(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* 0 * * ? *")
	assert.NoError(err)
	assert.Equal(hour(0), cron.Hour.Exps[0].Number)
}

func TestHoursNumberRange(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* 0-23 * * ? *")
	assert.NoError(err)
	assert.Equal(&cronplan.HourRange{
		Start: hour(0),
		End:   hour(23),
	}, cron.Hour.Exps[0].Range)
}

func TestHoursIncr(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* 0/5 * * ? *")
	assert.NoError(err)
	assert.Equal(&cronplan.HourExp{
		Number: hour(0),
		Bottom: intptr(5),
	}, cron.Hour.Exps[0])
}

func TestHoursIncrWildcard(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* */5 * * ? *")
	assert.NoError(err)
	assert.Equal(&cronplan.HourExp{
		Wildcard: true,
		Bottom:   intptr(5),
	}, cron.Hour.Exps[0])
}

func TestHoursIncrRange(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* 1-10/3 * * ? *")
	assert.NoError(err)
	assert.Equal(&cronplan.HourExp{
		Range: &cronplan.HourRange{
			Start: hour(1),
			End:   hour(10),
		},
		Bottom: intptr(3),
	}, cron.Hour.Exps[0])
}

func TestHoursComplex(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* *,0,0-23,0/5,*/5 * * ? *")
	assert.NoError(err)
	assert.True(cron.Hour.Exps[0].Wildcard)
	assert.Equal(hour(0), cron.Hour.Exps[1].Number)
	assert.Equal(&cronplan.HourRange{
		Start: hour(0),
		End:   hour(23),
	}, cron.Hour.Exps[2].Range)
	assert.Equal(&cronplan.HourExp{
		Number: hour(0),
		Bottom: intptr(5),
	}, cron.Hour.Exps[3])
	assert.Equal(&cronplan.HourExp{
		Wildcard: true,
		Bottom:   intptr(5),
	}, cron.Hour.Exps[4])
}

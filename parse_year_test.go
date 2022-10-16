package cronplan_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/cronplan"
)

func TestYearAll(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * * * ? *")
	assert.NoError(err)
	assert.True(cron.Year.Exps[0].Wildcard)
}

func TestYearNumber(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * * * ? 2022")
	assert.NoError(err)
	assert.Equal(year(2022), cron.Year.Exps[0].Number)
}

func TestYearNumberRange(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * * * ? 1970-2199")
	assert.NoError(err)
	assert.Equal(&cronplan.YearRange{
		Start: year(1970),
		End:   year(2199),
	}, cron.Year.Exps[0].Range)
}

func TestYearIncr(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * * * ? 1970/2")
	assert.NoError(err)
	assert.Equal(&cronplan.YearExp{
		Number: year(1970),
		Bottom: intptr(2),
	}, cron.Year.Exps[0])
}

func TestYearIncrWildcard(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * * * ? */2")
	assert.NoError(err)
	assert.Equal(&cronplan.YearExp{
		Wildcard: true,
		Bottom:   intptr(2),
	}, cron.Year.Exps[0])
}

func TestYearIncrRange(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * * * ? 2000-2020/2")
	assert.NoError(err)
	assert.Equal(&cronplan.YearExp{
		Range: &cronplan.YearRange{
			Start: year(2000),
			End:   year(2020),
		},
		Bottom: intptr(2),
	}, cron.Year.Exps[0])
}

func TestYearComplex(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * * * ? *,2022,1970-2199,1970/2,*/2")
	assert.NoError(err)
	assert.True(cron.Year.Exps[0].Wildcard)
	assert.Equal(year(2022), cron.Year.Exps[1].Number)
	assert.Equal(&cronplan.YearRange{
		Start: year(1970),
		End:   year(2199),
	}, cron.Year.Exps[2].Range)
	assert.Equal(&cronplan.YearExp{
		Number: year(1970),
		Bottom: intptr(2),
	}, cron.Year.Exps[3])
	assert.Equal(&cronplan.YearExp{
		Wildcard: true,
		Bottom:   intptr(2),
	}, cron.Year.Exps[4])
}

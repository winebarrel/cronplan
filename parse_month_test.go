package cronplan_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/cronplan"
)

func TestMonthAll(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parser.ParseString("", "* * * * ? *")
	assert.NoError(err)
	assert.True(cron.Month.Exps[0].Wildcard)
}

func TestMonthNumber(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parser.ParseString("", "* * * 1 ? *")
	assert.NoError(err)
	assert.Equal(month(time.January), cron.Month.Exps[0].Month)
}

func TestMonthNumberRange(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parser.ParseString("", "* * * 1-12 ? *")
	assert.NoError(err)
	assert.Equal(&cronplan.MonthRange{
		Start: month(time.January),
		End:   month(time.December),
	}, cron.Month.Exps[0].Range)
}

func TestMonthIncr(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parser.ParseString("", "* * * 1/5 ? *")
	assert.NoError(err)
	assert.Equal(&cronplan.MonthExp{
		Month:  month(time.January),
		Bottom: intptr(5),
	}, cron.Month.Exps[0])
}

func TestMonthIncrWildcard(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parser.ParseString("", "* * * */5 ? *")
	assert.NoError(err)
	assert.Equal(&cronplan.MonthExp{
		Wildcard: true,
		Bottom:   intptr(5),
	}, cron.Month.Exps[0])
}

func TestMonthName(t *testing.T) {
	assert := assert.New(t)
	tt := map[string]time.Month{
		"JAN": time.January,
		"FEB": time.February,
		"MAR": time.March,
		"APR": time.April,
		"MAY": time.May,
		"JUN": time.June,
		"JUL": time.July,
		"AUG": time.August,
		"SEP": time.September,
		"OCT": time.October,
		"NOV": time.November,
		"DEC": time.December,
	}

	for k, v := range tt {
		cron, err := cronplan.Parser.ParseString("", fmt.Sprintf("* * * %s ? *", k))
		assert.NoError(err)
		assert.Equal(month(v), cron.Month.Exps[0].Month, k)

		cron, err = cronplan.Parser.ParseString("", fmt.Sprintf("* * * %s ? *", strings.ToLower(k)))
		assert.NoError(err)
		assert.Equal(month(v), cron.Month.Exps[0].Month, k)
	}
}

func TestMonthNameRange(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parser.ParseString("", "* * * JAN-DEC ? *")
	assert.NoError(err)
	assert.Equal(&cronplan.MonthRange{
		Start: month(time.January),
		End:   month(time.December),
	}, cron.Month.Exps[0].Range)
}

func TestMonthNameIncrRange(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parser.ParseString("", "* * * JAN-DEC/3 ? *")
	assert.NoError(err)
	assert.Equal(
		&cronplan.MonthExp{
			Range: &cronplan.MonthRange{
				Start: month(time.January),
				End:   month(time.December),
			},
			Bottom: intptr(3),
		},
		cron.Month.Exps[0])
}

func TestMonthNameIncr(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parser.ParseString("", "* * * JAN/3 ? *")
	assert.NoError(err)
	assert.Equal(
		&cronplan.MonthExp{
			Month:  month(time.January),
			Bottom: intptr(3),
		},
		cron.Month.Exps[0])
}

func TestMonthComplex(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parser.ParseString("", "* * * *,1,1-12,1/5,*/5,JAN,JAN-DEC ? *")
	assert.NoError(err)
	assert.True(cron.Month.Exps[0].Wildcard)
	assert.Equal(month(time.January), cron.Month.Exps[1].Month)
	assert.Equal(&cronplan.MonthRange{
		Start: month(time.January),
		End:   month(time.December),
	}, cron.Month.Exps[2].Range)
	assert.Equal(&cronplan.MonthExp{
		Month:  month(time.January),
		Bottom: intptr(5),
	}, cron.Month.Exps[3])
	assert.Equal(&cronplan.MonthExp{
		Wildcard: true,
		Bottom:   intptr(5),
	}, cron.Month.Exps[4])
	assert.Equal(month(time.January), cron.Month.Exps[5].Month)
	assert.Equal(&cronplan.MonthRange{
		Start: month(time.January),
		End:   month(time.December),
	}, cron.Month.Exps[6].Range)
}

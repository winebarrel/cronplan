package cronplan_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/cronplan"
)

func TestMatchMinutes(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		exp   string
		tests []struct {
			tm       time.Time
			expected bool
		}
	}{
		{
			exp: "2 * * * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 1, 2, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 1, 3, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "2,3 * * * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 1, 2, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 1, 3, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 1, 4, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "2-4 * * * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 1, 2, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 1, 3, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 1, 4, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 1, 5, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "*/3 * * * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 1, 1, 0, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 1, 2, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 1, 3, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 1, 4, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 1, 5, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "1/3 * * * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 1, 1, 0, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 1, 2, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 1, 3, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 1, 4, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 1, 5, 0, 0, time.UTC), false},
			},
		},
	}

	for _, t := range tt {
		cron, err := cronplan.Parse(t.exp)
		assert.NoError(err)

		for _, t2 := range t.tests {
			assert.Equal(t2.expected, cron.Match(t2.tm), t.exp, t2.tm)
		}
	}
}

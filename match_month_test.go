package cronplan_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/cronplan"
)

func TestMatchMonth(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		exp   string
		tests []struct {
			tm       time.Time
			expected bool
		}
	}{
		{
			exp: "* * * 2 ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 1, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 2, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 3, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2023, 2, 1, 1, 1, 0, 0, time.UTC), true},
			},
		},
		{
			exp: "* * * 2,3 ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 1, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 2, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 3, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 4, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2023, 2, 1, 1, 1, 0, 0, time.UTC), true},
			},
		},
		{
			exp: "* * * 2-4 ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 1, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 2, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 3, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 4, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 5, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2023, 2, 1, 1, 1, 0, 0, time.UTC), true},
			},
		},
		{
			exp: "* * * */3 ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 1, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 2, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 3, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 4, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 5, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 6, 1, 1, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* * * 1/3 ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 1, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 2, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 3, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 4, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 5, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 6, 1, 1, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* * * FEB,MAR ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 1, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 2, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 3, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 4, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2023, 2, 1, 1, 1, 0, 0, time.UTC), true},
			},
		},
		{
			exp: "* * * FEB-APR ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 1, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 2, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 3, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 4, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 5, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2023, 2, 1, 1, 1, 0, 0, time.UTC), true},
			},
		},
		{
			exp: "* * * FEB-JUN/2 ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 1, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 2, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 3, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 4, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 5, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2023, 6, 1, 1, 1, 0, 0, time.UTC), true},
			},
		},
		{
			exp: "* * * 2-6/2 ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 1, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 2, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 3, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 4, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 5, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2023, 6, 1, 1, 1, 0, 0, time.UTC), true},
			},
		},
		{
			exp: "* * * JUN-FEB ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 5, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2023, 6, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2023, 7, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2023, 8, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2023, 9, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2023, 10, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2023, 11, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2023, 12, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2023, 1, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2023, 2, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2023, 3, 1, 1, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* * * 6-2 ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 5, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2023, 6, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2023, 7, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2023, 8, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2023, 9, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2023, 10, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2023, 11, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2023, 12, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2023, 1, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2023, 2, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2023, 3, 1, 1, 1, 0, 0, time.UTC), false},
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

package cronplan_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/cronplan/v2"
)

func TestMatchHours(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		exp   string
		tests []struct {
			tm       time.Time
			expected bool
		}
	}{
		{
			exp: "* 2 * * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 2, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 3, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* 2,3 * * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 2, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 3, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 4, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* 2-4 * * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 2, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 3, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 4, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 5, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* */3 * * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 1, 0, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 2, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 3, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 4, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 5, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* 1/3 * * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 1, 0, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 2, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 3, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 4, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 5, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* 1/0 * * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 1, 0, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 2, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 3, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 4, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 5, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 2, 1, 1, 0, 0, time.UTC), true},
			},
		},
		{
			exp: "* 1-10/3 * * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 1, 0, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 2, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 3, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 4, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 5, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 6, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 7, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 8, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 9, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 10, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 11, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 12, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 13, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* 1-10/0 * * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 1, 0, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 2, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 3, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 4, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 5, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 6, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 7, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 8, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 9, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 10, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 11, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 12, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 13, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 2, 0, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 2, 1, 1, 0, 0, time.UTC), true},
			},
		},
		{
			exp: "* 22-2 * * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 1, 21, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 22, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 23, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 2, 0, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 2, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 2, 2, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 2, 3, 1, 0, 0, time.UTC), false},
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

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
		{
			exp: "1-2/0 * * * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 1, 1, 0, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 1, 2, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 1, 3, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 1, 4, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 1, 5, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 2, 1, 0, 0, time.UTC), true},
			},
		},
		{
			exp: "1-10/3 * * * ? *",
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
				{time.Date(2022, 10, 1, 1, 6, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 1, 7, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 1, 8, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 1, 9, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 1, 10, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 1, 11, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 1, 12, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 1, 13, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "58-2 * * * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 1, 1, 57, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 1, 58, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 1, 59, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 2, 0, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 2, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 2, 2, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 1, 3, 0, 0, time.UTC), false},
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

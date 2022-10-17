package cronplan_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/cronplan"
)

func TestMatchDayObMonth(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		exp   string
		tests []struct {
			tm       time.Time
			expected bool
		}
	}{
		{
			exp: "* * 2 * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 2, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 3, 1, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* * 2,3 * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 2, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 3, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 4, 1, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* * 2-4 * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 2, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 3, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 4, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 5, 1, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* * */3 * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 2, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 3, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 4, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 5, 1, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* * 1/3 * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 2, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 3, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 4, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 5, 1, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* * 1/0 * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 2, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 3, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 4, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 5, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 11, 1, 1, 1, 0, 0, time.UTC), true},
			},
		},
		{
			exp: "* * 2/3 * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 2, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 3, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 4, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 5, 1, 1, 0, 0, time.UTC), true},
			},
		},
		{
			exp: "* * L * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 31, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 30, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2023, 2, 28, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2024, 2, 28, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2024, 2, 29, 1, 1, 0, 0, time.UTC), true},
			},
		},
		{
			exp: "* * L-2 * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 28, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 29, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 30, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 31, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2024, 2, 26, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2024, 2, 27, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2024, 2, 28, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2024, 2, 29, 1, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* * 3W * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 7, 3, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 7, 4, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 8, 3, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 9, 2, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 9, 3, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 3, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 11, 3, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 12, 2, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 12, 3, 1, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* * 1W * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 1, 0, 0, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 2, 0, 0, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 3, 0, 0, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 4, 0, 0, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* * 31W * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2023, 2, 28, 0, 0, 0, 0, time.UTC), false},
				{time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC), true},
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

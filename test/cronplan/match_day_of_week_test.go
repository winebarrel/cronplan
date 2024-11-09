package cronplan_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/cronplan"
)

func TestMatchDayOfWeek(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		exp   string
		tests []struct {
			tm       time.Time
			expected bool
		}
	}{
		{
			exp: "* * ? * 2 *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 2, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 3, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 4, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 5, 1, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* * ? * 2,3 *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 2, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 3, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 4, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 5, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 6, 1, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* * ? * 2-4 *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 2, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 3, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 4, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 5, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 6, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 7, 1, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* * ? * TUE,WED *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 2, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 3, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 4, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 5, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 6, 1, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* * ? * TUE-SUN *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 2, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 3, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 4, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 5, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 6, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 7, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 8, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 9, 1, 1, 0, 0, time.UTC), true},
			},
		},
		{
			exp: "* * ? * TUE-THU *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 2, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 3, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 4, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 5, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 6, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 7, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 8, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 9, 1, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* * ? *  SUN-FRI *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 2, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 3, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 4, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 5, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 6, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 7, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 8, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 9, 1, 1, 0, 0, time.UTC), true},
			},
		},
		{
			exp: "* * ? * */3 *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 2, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 3, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 4, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 5, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 6, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 7, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 8, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 9, 1, 1, 0, 0, time.UTC), true},
			},
		},
		{
			exp: "* * ? * 1/3 *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 2, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 3, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 4, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 5, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 6, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 7, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 8, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 9, 1, 1, 0, 0, time.UTC), true},
			},
		},
		{
			exp: "* * ? * 2/3 *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 2, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 3, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 4, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 5, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 6, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 7, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 8, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 9, 1, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* * ? * 2/0 *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 2, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 3, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 4, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 5, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 6, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 7, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 8, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 9, 1, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* * ? * 1/4 *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 2, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 3, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 4, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 5, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 6, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 7, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 8, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 9, 1, 1, 0, 0, time.UTC), true},
			},
		},
		{
			exp: "* * ? * 2/4 *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 2, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 3, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 4, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 5, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 6, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 7, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 8, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 9, 1, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* * ? * 2-4/2 *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 2, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 3, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 4, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 5, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 6, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 7, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 8, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 9, 1, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* * ? * 2#3 *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 10, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 17, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 18, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 11, 20, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 11, 21, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 11, 28, 1, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* * ? * MON#3 *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 10, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 17, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 18, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 11, 20, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 11, 21, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 11, 28, 1, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* * ? * 6L *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2023, 10, 27, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2023, 11, 25, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2023, 12, 18, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2024, 1, 26, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2024, 2, 22, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2024, 3, 20, 1, 1, 0, 0, time.UTC), false},
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

package cronplan_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/cronplan/v2"
)

func TestMatchYear(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		exp   string
		tests []struct {
			tm       time.Time
			expected bool
		}
	}{
		{
			exp: "* * * * ? 2022",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2021, 10, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2023, 10, 1, 1, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* * * * ? 2022,2023",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2021, 10, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2023, 10, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2024, 10, 1, 1, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* * * * ? 2022-2024",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2021, 10, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2023, 10, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2024, 10, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2025, 10, 1, 1, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* * * * ? */11",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(1959, 10, 1, 1, 0, 0, 0, time.UTC), false},
				{time.Date(1970, 10, 1, 1, 0, 0, 0, time.UTC), true},
				{time.Date(2024, 10, 1, 1, 0, 0, 0, time.UTC), false},
				{time.Date(2025, 10, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2026, 10, 1, 1, 2, 0, 0, time.UTC), false},
				{time.Date(2036, 10, 1, 1, 3, 0, 0, time.UTC), true},
				{time.Date(2047, 10, 1, 1, 4, 0, 0, time.UTC), true},
			},
		},
		{
			exp: "* * * * ? 2024/11",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2013, 10, 1, 1, 0, 0, 0, time.UTC), false},
				{time.Date(2023, 10, 1, 1, 0, 0, 0, time.UTC), false},
				{time.Date(2024, 10, 1, 1, 0, 0, 0, time.UTC), true},
				{time.Date(2025, 10, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2035, 10, 1, 1, 3, 0, 0, time.UTC), true},
				{time.Date(2046, 10, 1, 1, 4, 0, 0, time.UTC), true},
			},
		},
		{
			exp: "* * * * ? 2024-2190/0",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2013, 10, 1, 1, 0, 0, 0, time.UTC), false},
				{time.Date(2023, 10, 1, 1, 0, 0, 0, time.UTC), false},
				{time.Date(2024, 10, 1, 1, 0, 0, 0, time.UTC), true},
				{time.Date(2025, 10, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2035, 10, 1, 1, 3, 0, 0, time.UTC), false},
				{time.Date(2046, 10, 1, 1, 4, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* * * * ? 2022-2026/2",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2021, 10, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2023, 10, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2024, 10, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2025, 10, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2026, 10, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2027, 10, 1, 1, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* * * * ? 2198-1971",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2197, 10, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2198, 10, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2199, 10, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(1970, 10, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(1971, 10, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(1972, 10, 1, 1, 1, 0, 0, time.UTC), false},
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

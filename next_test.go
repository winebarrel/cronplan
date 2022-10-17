package cronplan_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/cronplan"
)

func TestNext(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		exp      string
		from     time.Time
		expected time.Time
	}{
		{
			exp:      "30 * * * ? *",
			from:     time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2022, 10, 10, 0, 30, 0, 0, time.UTC),
		},
		{
			exp:      "31 * * * ? *",
			from:     time.Date(2022, 10, 10, 0, 32, 0, 0, time.UTC),
			expected: time.Date(2022, 10, 10, 1, 31, 0, 0, time.UTC),
		},
		{
			exp:      "31 5 * * ? *",
			from:     time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2022, 10, 10, 5, 31, 0, 0, time.UTC),
		},
		{
			exp:      "31 5 * * ? *",
			from:     time.Date(2022, 10, 10, 5, 32, 0, 0, time.UTC),
			expected: time.Date(2022, 10, 11, 5, 31, 0, 0, time.UTC),
		},
		{
			exp:      "32 6 13,16 * ? *",
			from:     time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2022, 10, 13, 6, 32, 0, 0, time.UTC),
		},
		{
			exp:      "32 6 13,16 * ? *",
			from:     time.Date(2022, 10, 13, 6, 33, 0, 0, time.UTC),
			expected: time.Date(2022, 10, 16, 6, 32, 0, 0, time.UTC),
		},
		{
			exp:      "33 7 ? * FRI *",
			from:     time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2022, 10, 14, 7, 33, 0, 0, time.UTC),
		},
		{
			exp:      "33 7 ? * FRI *",
			from:     time.Date(2022, 10, 14, 7, 34, 0, 0, time.UTC),
			expected: time.Date(2022, 10, 21, 7, 33, 0, 0, time.UTC),
		},
		{
			exp:      "34 8 15 NOV ? *",
			from:     time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2022, 11, 15, 8, 34, 0, 0, time.UTC),
		},
		{
			exp:      "34 8 15 SEP ? *",
			from:     time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2023, 9, 15, 8, 34, 0, 0, time.UTC),
		},
		{
			exp:      "35 9 17 DEC ? 2023",
			from:     time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2023, 12, 17, 9, 35, 0, 0, time.UTC),
		},
		{
			exp:      "35 9 17 DEC ? 2020",
			from:     time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			expected: time.Time{},
		},
		{
			exp:      "35 9 L DEC ? 2023",
			from:     time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2023, 12, 31, 9, 35, 0, 0, time.UTC),
		},
		{
			exp:      "35 9 L-2 DEC ? 2023",
			from:     time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2023, 12, 29, 9, 35, 0, 0, time.UTC),
		},
		{
			exp:      "33 7 ? * FRI#4 *",
			from:     time.Date(2022, 10, 14, 7, 34, 0, 0, time.UTC),
			expected: time.Date(2022, 10, 28, 7, 33, 0, 0, time.UTC),
		},
		{
			exp:      "35 9 5W NOV ? *",
			from:     time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2022, 11, 4, 9, 35, 0, 0, time.UTC),
		},
		{
			exp:      "34 12-8 15 NOV ? *",
			from:     time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2022, 11, 15, 0, 34, 0, 0, time.UTC),
		},
		{
			exp:      "34 12-8 * * ? *",
			from:     time.Date(2022, 10, 10, 9, 0, 0, 0, time.UTC),
			expected: time.Date(2022, 10, 10, 12, 34, 0, 0, time.UTC),
		},
	}

	for _, t := range tt {
		cron, err := cronplan.Parse(t.exp)
		assert.NoError(err)
		next := cron.Next(t.from)
		assert.Equal(t.expected, next, t)
	}
}

func TestNextN_10(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("*/5 * * * ? *")
	assert.NoError(err)
	schedule := cron.NextN(time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC), 10)
	assert.Equal(
		[]time.Time{
			time.Date(2022, time.October, 10, 0, 0, 0, 0, time.UTC),
			time.Date(2022, time.October, 10, 0, 5, 0, 0, time.UTC),
			time.Date(2022, time.October, 10, 0, 10, 0, 0, time.UTC),
			time.Date(2022, time.October, 10, 0, 15, 0, 0, time.UTC),
			time.Date(2022, time.October, 10, 0, 20, 0, 0, time.UTC),
			time.Date(2022, time.October, 10, 0, 25, 0, 0, time.UTC),
			time.Date(2022, time.October, 10, 0, 30, 0, 0, time.UTC),
			time.Date(2022, time.October, 10, 0, 35, 0, 0, time.UTC),
			time.Date(2022, time.October, 10, 0, 40, 0, 0, time.UTC),
			time.Date(2022, time.October, 10, 0, 45, 0, 0, time.UTC),
		},
		schedule,
	)
}

func TestNextN_3(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		exp      string
		from     time.Time
		expected []time.Time
	}{
		{
			exp:  "30 * * * ? *",
			from: time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 10, 10, 0, 30, 0, 0, time.UTC),
				time.Date(2022, 10, 10, 1, 30, 0, 0, time.UTC),
				time.Date(2022, 10, 10, 2, 30, 0, 0, time.UTC),
			},
		},
		{
			exp:  "31 * * * ? *",
			from: time.Date(2022, 10, 10, 0, 32, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 10, 10, 1, 31, 0, 0, time.UTC),
				time.Date(2022, 10, 10, 2, 31, 0, 0, time.UTC),
				time.Date(2022, 10, 10, 3, 31, 0, 0, time.UTC),
			},
		},
		{
			exp:  "31 5 * * ? *",
			from: time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 10, 10, 5, 31, 0, 0, time.UTC),
				time.Date(2022, 10, 11, 5, 31, 0, 0, time.UTC),
				time.Date(2022, 10, 12, 5, 31, 0, 0, time.UTC),
			},
		},
		{
			exp:  "31 5 * * ? *",
			from: time.Date(2022, 10, 10, 5, 32, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 10, 11, 5, 31, 0, 0, time.UTC),
				time.Date(2022, 10, 12, 5, 31, 0, 0, time.UTC),
				time.Date(2022, 10, 13, 5, 31, 0, 0, time.UTC),
			},
		},
		{
			exp:  "32 6 13,16 * ? *",
			from: time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 10, 13, 6, 32, 0, 0, time.UTC),
				time.Date(2022, 10, 16, 6, 32, 0, 0, time.UTC),
				time.Date(2022, 11, 13, 6, 32, 0, 0, time.UTC),
			},
		},
		{
			exp:  "32 6 13,16 * ? *",
			from: time.Date(2022, 10, 13, 6, 33, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 10, 16, 6, 32, 0, 0, time.UTC),
				time.Date(2022, 11, 13, 6, 32, 0, 0, time.UTC),
				time.Date(2022, 11, 16, 6, 32, 0, 0, time.UTC),
			},
		},
		{
			exp:  "33 7 ? * FRI *",
			from: time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 10, 14, 7, 33, 0, 0, time.UTC),
				time.Date(2022, 10, 21, 7, 33, 0, 0, time.UTC),
				time.Date(2022, 10, 28, 7, 33, 0, 0, time.UTC),
			},
		},
		{
			exp:  "33 7 ? * FRI *",
			from: time.Date(2022, 10, 14, 7, 34, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 10, 21, 7, 33, 0, 0, time.UTC),
				time.Date(2022, 10, 28, 7, 33, 0, 0, time.UTC),
				time.Date(2022, 11, 4, 7, 33, 0, 0, time.UTC),
			},
		},
		{
			exp:  "34 8 15 NOV ? *",
			from: time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 11, 15, 8, 34, 0, 0, time.UTC),
				time.Date(2023, 11, 15, 8, 34, 0, 0, time.UTC),
				time.Date(2024, 11, 15, 8, 34, 0, 0, time.UTC),
			},
		},
		{
			exp:  "34 8 15 SEP ? *",
			from: time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2023, 9, 15, 8, 34, 0, 0, time.UTC),
				time.Date(2024, 9, 15, 8, 34, 0, 0, time.UTC),
				time.Date(2025, 9, 15, 8, 34, 0, 0, time.UTC),
			},
		},
		{
			exp:  "35 9 17 DEC ? 2023",
			from: time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2023, 12, 17, 9, 35, 0, 0, time.UTC),
			},
		},
		{
			exp:      "35 9 17 DEC ? 2020",
			from:     time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			expected: []time.Time{},
		},
	}

	for _, t := range tt {
		cron, err := cronplan.Parse(t.exp)
		assert.NoError(err)
		next := cron.NextN(t.from, 3)
		assert.Equal(t.expected, next, t)
	}
}

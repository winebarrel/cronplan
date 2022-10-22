package cronplan_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/cronplan"
)

func TestBetween(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		exp      string
		from     time.Time
		to       time.Time
		expected []time.Time
	}{
		{
			exp:  "0 * * * ? *",
			from: time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			to:   time.Date(2022, 10, 10, 2, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
				time.Date(2022, 10, 10, 1, 0, 0, 0, time.UTC),
				time.Date(2022, 10, 10, 2, 0, 0, 0, time.UTC),
			},
		},
		{
			exp:  "30 * * * ? *",
			from: time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			to:   time.Date(2022, 10, 10, 2, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 10, 10, 0, 30, 0, 0, time.UTC),
				time.Date(2022, 10, 10, 1, 30, 0, 0, time.UTC),
			},
		},
		{
			exp:  "31 * * * ? *",
			from: time.Date(2022, 10, 10, 0, 32, 0, 0, time.UTC),
			to:   time.Date(2022, 10, 10, 2, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 10, 10, 1, 31, 0, 0, time.UTC),
			},
		},
		{
			exp:  "31 5 * * ? *",
			from: time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			to:   time.Date(2022, 10, 12, 0, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 10, 10, 5, 31, 0, 0, time.UTC),
				time.Date(2022, 10, 11, 5, 31, 0, 0, time.UTC),
			},
		},
		{
			exp:  "31 5 * * ? *",
			from: time.Date(2022, 10, 10, 5, 32, 0, 0, time.UTC),
			to:   time.Date(2022, 10, 13, 0, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 10, 11, 5, 31, 0, 0, time.UTC),
				time.Date(2022, 10, 12, 5, 31, 0, 0, time.UTC),
			},
		},
		{
			exp:  "32 6 13,16 * ? *",
			from: time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			to:   time.Date(2022, 10, 20, 0, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 10, 13, 6, 32, 0, 0, time.UTC),
				time.Date(2022, 10, 16, 6, 32, 0, 0, time.UTC),
			},
		},
		{
			exp:  "32 6 13,16 * ? *",
			from: time.Date(2022, 10, 13, 6, 33, 0, 0, time.UTC),
			to:   time.Date(2022, 10, 20, 0, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 10, 16, 6, 32, 0, 0, time.UTC),
			},
		},
		{
			exp:  "33 7 ? * FRI *",
			from: time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			to:   time.Date(2022, 10, 30, 0, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 10, 14, 7, 33, 0, 0, time.UTC),
				time.Date(2022, 10, 21, 7, 33, 0, 0, time.UTC),
				time.Date(2022, 10, 28, 7, 33, 0, 0, time.UTC),
			},
		},
		{
			exp:  "33 7 ? * FRI *",
			from: time.Date(2022, 10, 14, 7, 34, 0, 0, time.UTC),
			to:   time.Date(2022, 10, 30, 0, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 10, 21, 7, 33, 0, 0, time.UTC),
				time.Date(2022, 10, 28, 7, 33, 0, 0, time.UTC),
			},
		},
		{
			exp:  "34 8 15 NOV ? *",
			from: time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			to:   time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 11, 15, 8, 34, 0, 0, time.UTC),
				time.Date(2023, 11, 15, 8, 34, 0, 0, time.UTC),
			},
		},
		{
			exp:  "34 8 15 SEP ? *",
			from: time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			to:   time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2023, 9, 15, 8, 34, 0, 0, time.UTC),
			},
		},
		{
			exp:  "35 9 17 DEC ? 2023",
			from: time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			to:   time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2023, 12, 17, 9, 35, 0, 0, time.UTC),
			},
		},
		{
			exp:      "35 9 17 DEC ? 2020",
			from:     time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			to:       time.Date(2024, 10, 10, 0, 0, 0, 0, time.UTC),
			expected: []time.Time{},
		},
		{
			exp:  "35 9 L DEC ? 2023",
			from: time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			to:   time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2023, 12, 31, 9, 35, 0, 0, time.UTC),
			},
		},
		{
			exp:  "35 9 L DEC ? 2023-2024",
			from: time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			to:   time.Date(2024, 12, 31, 10, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2023, 12, 31, 9, 35, 0, 0, time.UTC),
				time.Date(2024, 12, 31, 9, 35, 0, 0, time.UTC),
			},
		},
		{
			exp:  "35 9 L-2 DEC ? 2023",
			from: time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			to:   time.Date(2024, 12, 31, 10, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2023, 12, 29, 9, 35, 0, 0, time.UTC),
			},
		},
		{
			exp:      "35 9 L-2 DEC ? 2023",
			from:     time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			to:       time.Date(2023, 12, 29, 9, 0, 0, 0, time.UTC),
			expected: []time.Time{},
		},
		{
			exp:  "33 7 ? * FRI#4 *",
			from: time.Date(2022, 10, 14, 7, 34, 0, 0, time.UTC),
			to:   time.Date(2022, 11, 30, 0, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 10, 28, 7, 33, 0, 0, time.UTC),
				time.Date(2022, 11, 25, 7, 33, 0, 0, time.UTC),
			},
		},
		{
			exp:  "35 9 5W NOV ? *",
			from: time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			to:   time.Date(2023, 11, 30, 0, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 11, 4, 9, 35, 0, 0, time.UTC),
				time.Date(2023, 11, 6, 9, 35, 0, 0, time.UTC),
			},
		},
		{
			exp:  "34 12-8 15 NOV ? *",
			from: time.Date(2022, 11, 15, 0, 0, 0, 0, time.UTC),
			to:   time.Date(2022, 11, 15, 10, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 11, 15, 0, 34, 0, 0, time.UTC),
				time.Date(2022, 11, 15, 1, 34, 0, 0, time.UTC),
				time.Date(2022, 11, 15, 2, 34, 0, 0, time.UTC),
				time.Date(2022, 11, 15, 3, 34, 0, 0, time.UTC),
				time.Date(2022, 11, 15, 4, 34, 0, 0, time.UTC),
				time.Date(2022, 11, 15, 5, 34, 0, 0, time.UTC),
				time.Date(2022, 11, 15, 6, 34, 0, 0, time.UTC),
				time.Date(2022, 11, 15, 7, 34, 0, 0, time.UTC),
				time.Date(2022, 11, 15, 8, 34, 0, 0, time.UTC),
			},
		},
		{
			exp:  "34 12-8 * * ? *",
			from: time.Date(2022, 10, 10, 9, 0, 0, 0, time.UTC),
			to:   time.Date(2022, 10, 10, 14, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 10, 10, 12, 34, 0, 0, time.UTC),
				time.Date(2022, 10, 10, 13, 34, 0, 0, time.UTC),
			},
		},
	}

	for _, t := range tt {
		cron, err := cronplan.Parse(t.exp)
		assert.NoError(err)
		next := cron.Between(t.from, t.to)
		assert.Equal(t.expected, next, t)
	}
}

func TestBetween_From_eq_To(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * * * ? *")
	assert.NoError(err)
	from := time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC)
	to := time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC)
	schedule := cron.Between(from, to)
	assert.Equal([]time.Time{}, schedule)
}

func TestBetween_From_gt_To(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("* * * * ? *")
	assert.NoError(err)
	from := time.Date(2022, 10, 10, 0, 0, 1, 0, time.UTC)
	to := time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC)
	schedule := cron.Between(from, to)
	assert.Equal([]time.Time{}, schedule)
}

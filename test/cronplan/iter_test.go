package cronplan_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/winebarrel/cronplan/v2"
)

func TestIter(t *testing.T) {
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
				{},
				{},
			},
		},
		{
			exp:      "35 9 17 DEC ? 2020",
			from:     time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			expected: []time.Time{{}, {}, {}},
		},
		{
			exp:  "35 9 L DEC ? 2023",
			from: time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2023, 12, 31, 9, 35, 0, 0, time.UTC),
				{},
				{},
			},
		},
		{
			exp:  "35 9 L-2 DEC ? 2023",
			from: time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2023, 12, 29, 9, 35, 0, 0, time.UTC),
				{},
				{},
			},
		},
		{
			exp:  "33 7 ? * FRI#4 *",
			from: time.Date(2022, 10, 14, 7, 34, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 10, 28, 7, 33, 0, 0, time.UTC),
				time.Date(2022, 11, 25, 7, 33, 0, 0, time.UTC),
				time.Date(2022, 12, 23, 7, 33, 0, 0, time.UTC),
			},
		},
		{
			exp:  "32 6 ? * FRIL *",
			from: time.Date(2022, 10, 14, 6, 32, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 10, 28, 6, 32, 0, 0, time.UTC),
				time.Date(2022, 11, 25, 6, 32, 0, 0, time.UTC),
				time.Date(2022, 12, 30, 6, 32, 0, 0, time.UTC),
			},
		},
		{
			exp:  "31 5 ? * 2L *",
			from: time.Date(2022, 10, 14, 5, 31, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 10, 31, 5, 31, 0, 0, time.UTC),
				time.Date(2022, 11, 28, 5, 31, 0, 0, time.UTC),
				time.Date(2022, 12, 26, 5, 31, 0, 0, time.UTC),
			},
		},
		{
			exp:  "35 9 5W NOV ? *",
			from: time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 11, 4, 9, 35, 0, 0, time.UTC),
				time.Date(2023, 11, 6, 9, 35, 0, 0, time.UTC),
				time.Date(2024, 11, 5, 9, 35, 0, 0, time.UTC),
			},
		},
		{
			exp:  "34 12-8 15 NOV ? *",
			from: time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 11, 15, 0, 34, 0, 0, time.UTC),
				time.Date(2022, 11, 15, 1, 34, 0, 0, time.UTC),
				time.Date(2022, 11, 15, 2, 34, 0, 0, time.UTC),
			},
		},
		{
			exp:  "34 12-8 * * ? *",
			from: time.Date(2022, 10, 10, 9, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 10, 10, 12, 34, 0, 0, time.UTC),
				time.Date(2022, 10, 10, 13, 34, 0, 0, time.UTC),
				time.Date(2022, 10, 10, 14, 34, 0, 0, time.UTC),
			},
		},
		{
			exp:  "35 13 LW * ? *",
			from: time.Date(2022, 10, 10, 9, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 10, 31, 13, 35, 0, 0, time.UTC),
				time.Date(2022, 11, 30, 13, 35, 0, 0, time.UTC),
				time.Date(2022, 12, 30, 13, 35, 0, 0, time.UTC),
			},
		},
	}

	for _, test := range tt {
		cron, err := cronplan.Parse(test.exp)
		assert.NoError(err)
		iter := cron.Iter(test.from)

		for _, e := range test.expected {
			next := iter.Next()
			assert.Equal(e, next, test)
		}
	}
}

func TestIterHasNext(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		exp      string
		from     time.Time
		expected []bool
	}{
		{
			exp:      "30 * * * ? *",
			from:     time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			expected: []bool{true, true, true},
		},
		{
			exp:      "35 9 17 DEC ? 2023",
			from:     time.Date(2023, 12, 17, 9, 35, 0, 0, time.UTC),
			expected: []bool{true, false, false},
		},
		{
			exp:      "35 9 17 DEC ? 2023",
			from:     time.Date(2023, 12, 17, 9, 36, 0, 0, time.UTC),
			expected: []bool{false, false, false},
		},
	}

	for _, test := range tt {
		cron, err := cronplan.Parse(test.exp)
		assert.NoError(err)
		iter := cron.Iter(test.from)

		for _, e := range test.expected {
			next := iter.HasNext()
			iter.Next()
			assert.Equal(e, next, test)
		}
	}
}

func TestIterSeq(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

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
		{
			exp:  "35 9 L DEC ? 2023",
			from: time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2023, 12, 31, 9, 35, 0, 0, time.UTC),
			},
		},
		{
			exp:  "35 9 L-2 DEC ? 2023",
			from: time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2023, 12, 29, 9, 35, 0, 0, time.UTC),
			},
		},
		{
			exp:  "33 7 ? * FRI#4 *",
			from: time.Date(2022, 10, 14, 7, 34, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 10, 28, 7, 33, 0, 0, time.UTC),
				time.Date(2022, 11, 25, 7, 33, 0, 0, time.UTC),
				time.Date(2022, 12, 23, 7, 33, 0, 0, time.UTC),
			},
		},
		{
			exp:  "32 6 ? * FRIL *",
			from: time.Date(2022, 10, 14, 6, 32, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 10, 28, 6, 32, 0, 0, time.UTC),
				time.Date(2022, 11, 25, 6, 32, 0, 0, time.UTC),
				time.Date(2022, 12, 30, 6, 32, 0, 0, time.UTC),
			},
		},
		{
			exp:  "31 5 ? * 2L *",
			from: time.Date(2022, 10, 14, 5, 31, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 10, 31, 5, 31, 0, 0, time.UTC),
				time.Date(2022, 11, 28, 5, 31, 0, 0, time.UTC),
				time.Date(2022, 12, 26, 5, 31, 0, 0, time.UTC),
			},
		},
		{
			exp:  "35 9 5W NOV ? *",
			from: time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 11, 4, 9, 35, 0, 0, time.UTC),
				time.Date(2023, 11, 6, 9, 35, 0, 0, time.UTC),
				time.Date(2024, 11, 5, 9, 35, 0, 0, time.UTC),
			},
		},
		{
			exp:  "34 12-8 15 NOV ? *",
			from: time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 11, 15, 0, 34, 0, 0, time.UTC),
				time.Date(2022, 11, 15, 1, 34, 0, 0, time.UTC),
				time.Date(2022, 11, 15, 2, 34, 0, 0, time.UTC),
			},
		},
		{
			exp:  "34 12-8 * * ? *",
			from: time.Date(2022, 10, 10, 9, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 10, 10, 12, 34, 0, 0, time.UTC),
				time.Date(2022, 10, 10, 13, 34, 0, 0, time.UTC),
				time.Date(2022, 10, 10, 14, 34, 0, 0, time.UTC),
			},
		},
		{
			exp:  "35 13 LW * ? *",
			from: time.Date(2022, 10, 10, 9, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 10, 31, 13, 35, 0, 0, time.UTC),
				time.Date(2022, 11, 30, 13, 35, 0, 0, time.UTC),
				time.Date(2022, 12, 30, 13, 35, 0, 0, time.UTC),
			},
		},
	}

	for _, test := range tt {
		cron, err := cronplan.Parse(test.exp)
		assert.NoError(err)
		iter := cron.Iter(test.from)

		i := 0

		for next := range iter.Seq() {
			require.Greater(len(test.expected), i)
			e := test.expected[i]
			assert.Equal(e, next, test)
			i++

			if i >= 3 {
				break
			}
		}
	}
}

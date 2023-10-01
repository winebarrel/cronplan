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
			exp:      "32 6 ? * FRIL *",
			from:     time.Date(2022, 10, 14, 6, 32, 0, 0, time.UTC),
			expected: time.Date(2022, 10, 28, 6, 32, 0, 0, time.UTC),
		},
		{
			exp:      "31 5 ? * 2L *",
			from:     time.Date(2022, 10, 14, 5, 31, 0, 0, time.UTC),
			expected: time.Date(2022, 10, 31, 5, 31, 0, 0, time.UTC),
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
		{
			exp:      "35 13 LW * ? *",
			from:     time.Date(2022, 10, 10, 9, 0, 0, 0, time.UTC),
			expected: time.Date(2022, 10, 31, 13, 35, 0, 0, time.UTC),
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

func TestNextN_10_LastFri(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("0 0 ? * 6L *")
	assert.NoError(err)
	schedule := cron.NextN(time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC), 10)
	assert.Equal(
		[]time.Time{
			time.Date(2023, time.October, 27, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.November, 24, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 29, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.January, 26, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.February, 23, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.March, 29, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.April, 26, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.May, 31, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.June, 28, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.July, 26, 0, 0, 0, 0, time.UTC),
		},
		schedule,
	)
}

func TestNextN_10_LastSat(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("0 0 ? * SATL *")
	assert.NoError(err)
	schedule := cron.NextN(time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC), 10)
	assert.Equal(
		[]time.Time{
			time.Date(2023, time.October, 28, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.November, 25, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 30, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.January, 27, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.February, 24, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.March, 30, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.April, 27, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.May, 25, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.June, 29, 0, 0, 0, 0, time.UTC),
			time.Date(2024, time.July, 27, 0, 0, 0, 0, time.UTC),
		},
		schedule,
	)
}

func TestNextN_10_LastWdayWithoutWday(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("0 0 ? * L *")
	assert.NoError(err)
	schedule := cron.NextN(time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC), 10)
	assert.Equal(
		[]time.Time{
			time.Date(2023, time.October, 7, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.October, 14, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.October, 21, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.October, 28, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.November, 04, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.November, 11, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.November, 18, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.November, 25, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC),
			time.Date(2023, time.December, 9, 0, 0, 0, 0, time.UTC),
		},
		schedule,
	)
}

func TestNextN_TZ(t *testing.T) {
	jst, err := time.LoadLocation("Asia/Tokyo")

	if err != nil {
		panic(err)
	}

	assert := assert.New(t)
	cron, err := cronplan.Parse("*/5 * * * ? *")
	assert.NoError(err)

	schedule := cron.NextN(time.Date(2022, 10, 10, 0, 0, 0, 0, jst), 10)
	assert.Equal(
		[]time.Time{
			time.Date(2022, time.October, 10, 0, 0, 0, 0, jst),
			time.Date(2022, time.October, 10, 0, 5, 0, 0, jst),
			time.Date(2022, time.October, 10, 0, 10, 0, 0, jst),
			time.Date(2022, time.October, 10, 0, 15, 0, 0, jst),
			time.Date(2022, time.October, 10, 0, 20, 0, 0, jst),
			time.Date(2022, time.October, 10, 0, 25, 0, 0, jst),
			time.Date(2022, time.October, 10, 0, 30, 0, 0, jst),
			time.Date(2022, time.October, 10, 0, 35, 0, 0, jst),
			time.Date(2022, time.October, 10, 0, 40, 0, 0, jst),
			time.Date(2022, time.October, 10, 0, 45, 0, 0, jst),
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
			exp:  "33 7 ? * 6L *",
			from: time.Date(2022, 10, 14, 7, 34, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 10, 28, 7, 33, 0, 0, time.UTC),
				time.Date(2022, 11, 25, 7, 33, 0, 0, time.UTC),
				time.Date(2022, 12, 30, 7, 33, 0, 0, time.UTC),
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
			exp:  "34 8 LW * ? *",
			from: time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC),
			expected: []time.Time{
				time.Date(2022, 10, 31, 8, 34, 0, 0, time.UTC),
				time.Date(2022, 11, 30, 8, 34, 0, 0, time.UTC),
				time.Date(2022, 12, 30, 8, 34, 0, 0, time.UTC),
			},
		},
	}

	for _, t := range tt {
		cron, err := cronplan.Parse(t.exp)
		assert.NoError(err)
		next := cron.NextN(t.from, 3)
		assert.Equal(t.expected, next, t)
	}
}

func TestNextN_0(t *testing.T) {
	assert := assert.New(t)
	cron, err := cronplan.Parse("*/5 * * * ? *")
	assert.NoError(err)
	schedule := cron.NextN(time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC), 0)
	assert.Equal([]time.Time{}, schedule)
}

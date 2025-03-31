package util_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/cronplan/internal/util"
)

func TestCastWeekday(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		wday     string
		expected time.Weekday
	}{
		{"mon", time.Monday},
		{"tue", time.Tuesday},
		{"wed", time.Wednesday},
		{"thu", time.Thursday},
		{"fri", time.Friday},
		{"sat", time.Saturday},
		{"sun", time.Sunday},
	}

	for _, t := range tt {
		actual, err := util.CastWeekday(t.wday)
		assert.NoError(err)
		assert.Equal(t.expected, actual, t.wday)
	}
}

func TestCastWeekdayError(t *testing.T) {
	assert := assert.New(t)
	_, err := util.CastWeekday("xxx")
	assert.EqualError(err, "cannot convert to weekday from xxx")
}

func TestCastMonth(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		mon      string
		expected time.Month
	}{
		{"jan", time.January},
		{"feb", time.February},
		{"mar", time.March},
		{"apr", time.April},
		{"may", time.May},
		{"jun", time.June},
		{"jul", time.July},
		{"aug", time.August},
		{"sep", time.September},
		{"oct", time.October},
		{"nov", time.November},
		{"dec", time.December},
	}

	for _, t := range tt {
		actual, err := util.CastMonth(t.mon)
		assert.NoError(err)
		assert.Equal(t.expected, actual, t.mon)
	}
}

func TestCastMonthError(t *testing.T) {
	assert := assert.New(t)
	_, err := util.CastMonth("xxx")
	assert.EqualError(err, "cannot convert to month from xxx")
}

func TestLastOfMonth(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		tm       time.Time
		expected int
	}{
		{time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2023, 2, 1, 9, 0, 0, 0, time.UTC), 28},
		{time.Date(2023, 3, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2023, 4, 1, 9, 0, 0, 0, time.UTC), 30},
		{time.Date(2023, 5, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2023, 6, 1, 9, 0, 0, 0, time.UTC), 30},
		{time.Date(2023, 7, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2023, 8, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2023, 9, 1, 9, 0, 0, 0, time.UTC), 30},
		{time.Date(2023, 10, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2023, 11, 1, 9, 0, 0, 0, time.UTC), 30},
		{time.Date(2023, 12, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2024, 2, 1, 9, 0, 0, 0, time.UTC), 29},
		{time.Date(2024, 3, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2024, 4, 1, 9, 0, 0, 0, time.UTC), 30},
		{time.Date(2024, 5, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2024, 6, 1, 9, 0, 0, 0, time.UTC), 30},
		{time.Date(2024, 7, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2024, 8, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2024, 9, 1, 9, 0, 0, 0, time.UTC), 30},
		{time.Date(2024, 10, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2024, 11, 1, 9, 0, 0, 0, time.UTC), 30},
		{time.Date(2024, 12, 1, 9, 0, 0, 0, time.UTC), 31},

		{time.Date(2023, 1, 31, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2023, 2, 28, 9, 0, 0, 0, time.UTC), 28},
		{time.Date(2023, 3, 31, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2023, 4, 30, 9, 0, 0, 0, time.UTC), 30},
		{time.Date(2023, 5, 31, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2023, 6, 30, 9, 0, 0, 0, time.UTC), 30},
		{time.Date(2023, 7, 31, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2023, 8, 31, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2023, 9, 30, 9, 0, 0, 0, time.UTC), 30},
		{time.Date(2023, 10, 31, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2023, 11, 30, 9, 0, 0, 0, time.UTC), 30},
		{time.Date(2023, 12, 31, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2024, 1, 31, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2024, 2, 29, 9, 0, 0, 0, time.UTC), 29},
		{time.Date(2024, 3, 31, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2024, 4, 30, 9, 0, 0, 0, time.UTC), 30},
		{time.Date(2024, 5, 31, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2024, 6, 30, 9, 0, 0, 0, time.UTC), 30},
		{time.Date(2024, 7, 31, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2024, 8, 31, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2024, 9, 30, 9, 0, 0, 0, time.UTC), 30},
		{time.Date(2024, 10, 31, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2024, 11, 30, 9, 0, 0, 0, time.UTC), 30},
		{time.Date(2024, 12, 31, 9, 0, 0, 0, time.UTC), 31},
	}

	for _, t := range tt {
		assert.Equal(t.expected, util.LastOfMonth(t.tm), t.tm)
	}
}

func TestLastWdayOfMonth(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		tm       time.Time
		wday     time.Weekday
		expected int
	}{
		{time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC), time.Sunday, 29},
		{time.Date(2023, 2, 1, 9, 0, 0, 0, time.UTC), time.Monday, 27},
		{time.Date(2023, 3, 1, 9, 0, 0, 0, time.UTC), time.Tuesday, 28},
		{time.Date(2023, 4, 1, 9, 0, 0, 0, time.UTC), time.Wednesday, 26},
		{time.Date(2023, 5, 1, 9, 0, 0, 0, time.UTC), time.Thursday, 25},
		{time.Date(2023, 6, 1, 9, 0, 0, 0, time.UTC), time.Friday, 30},
		{time.Date(2023, 7, 1, 9, 0, 0, 0, time.UTC), time.Saturday, 29},
		{time.Date(2023, 8, 1, 9, 0, 0, 0, time.UTC), time.Sunday, 27},
		{time.Date(2023, 9, 1, 9, 0, 0, 0, time.UTC), time.Monday, 25},
		{time.Date(2023, 10, 1, 9, 0, 0, 0, time.UTC), time.Tuesday, 31},
		{time.Date(2023, 11, 1, 9, 0, 0, 0, time.UTC), time.Wednesday, 29},
		{time.Date(2023, 12, 1, 9, 0, 0, 0, time.UTC), time.Thursday, 28},
		{time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), time.Friday, 26},
		{time.Date(2024, 2, 1, 9, 0, 0, 0, time.UTC), time.Saturday, 24},
		{time.Date(2024, 3, 1, 9, 0, 0, 0, time.UTC), time.Sunday, 31},
		{time.Date(2024, 4, 1, 9, 0, 0, 0, time.UTC), time.Monday, 29},
		{time.Date(2024, 5, 1, 9, 0, 0, 0, time.UTC), time.Tuesday, 28},
		{time.Date(2024, 6, 1, 9, 0, 0, 0, time.UTC), time.Wednesday, 26},
		{time.Date(2024, 7, 1, 9, 0, 0, 0, time.UTC), time.Thursday, 25},
		{time.Date(2024, 8, 1, 9, 0, 0, 0, time.UTC), time.Friday, 30},
		{time.Date(2024, 9, 1, 9, 0, 0, 0, time.UTC), time.Saturday, 28},
		{time.Date(2024, 10, 1, 9, 0, 0, 0, time.UTC), time.Sunday, 27},
		{time.Date(2024, 11, 1, 9, 0, 0, 0, time.UTC), time.Monday, 25},
		{time.Date(2024, 12, 1, 9, 0, 0, 0, time.UTC), time.Tuesday, 31},

		{time.Date(2023, 1, 31, 9, 0, 0, 0, time.UTC), time.Sunday, 29},
		{time.Date(2023, 2, 28, 9, 0, 0, 0, time.UTC), time.Monday, 27},
		{time.Date(2023, 3, 31, 9, 0, 0, 0, time.UTC), time.Tuesday, 28},
		{time.Date(2023, 4, 30, 9, 0, 0, 0, time.UTC), time.Wednesday, 26},
		{time.Date(2023, 5, 31, 9, 0, 0, 0, time.UTC), time.Thursday, 25},
		{time.Date(2023, 6, 30, 9, 0, 0, 0, time.UTC), time.Friday, 30},
		{time.Date(2023, 7, 31, 9, 0, 0, 0, time.UTC), time.Saturday, 29},
		{time.Date(2023, 8, 31, 9, 0, 0, 0, time.UTC), time.Sunday, 27},
		{time.Date(2023, 9, 30, 9, 0, 0, 0, time.UTC), time.Monday, 25},
		{time.Date(2023, 10, 31, 9, 0, 0, 0, time.UTC), time.Tuesday, 31},
		{time.Date(2023, 11, 30, 9, 0, 0, 0, time.UTC), time.Wednesday, 29},
		{time.Date(2023, 12, 31, 9, 0, 0, 0, time.UTC), time.Thursday, 28},
		{time.Date(2024, 1, 31, 9, 0, 0, 0, time.UTC), time.Friday, 26},
		{time.Date(2024, 2, 29, 9, 0, 0, 0, time.UTC), time.Saturday, 24},
		{time.Date(2024, 3, 31, 9, 0, 0, 0, time.UTC), time.Sunday, 31},
		{time.Date(2024, 4, 30, 9, 0, 0, 0, time.UTC), time.Monday, 29},
		{time.Date(2024, 5, 31, 9, 0, 0, 0, time.UTC), time.Tuesday, 28},
		{time.Date(2024, 6, 30, 9, 0, 0, 0, time.UTC), time.Wednesday, 26},
		{time.Date(2024, 7, 31, 9, 0, 0, 0, time.UTC), time.Thursday, 25},
		{time.Date(2024, 8, 31, 9, 0, 0, 0, time.UTC), time.Friday, 30},
		{time.Date(2024, 9, 30, 9, 0, 0, 0, time.UTC), time.Saturday, 28},
		{time.Date(2024, 10, 31, 9, 0, 0, 0, time.UTC), time.Sunday, 27},
		{time.Date(2024, 11, 30, 9, 0, 0, 0, time.UTC), time.Monday, 25},
		{time.Date(2024, 12, 31, 9, 0, 0, 0, time.UTC), time.Tuesday, 31},
	}

	for _, t := range tt {
		assert.Equal(t.expected, util.LastWdayOfMonth(t.tm, t.wday), t.tm)
	}
}

func TestNearestWeekday(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		tm       time.Time
		day      int
		expected int
	}{
		{time.Date(2022, 11, 1, 0, 0, 0, 0, time.UTC), 3, 3},
		{time.Date(2022, 11, 1, 0, 0, 0, 0, time.UTC), 4, 4},
		{time.Date(2022, 11, 1, 0, 0, 0, 0, time.UTC), 5, 4},
		{time.Date(2022, 11, 1, 0, 0, 0, 0, time.UTC), 6, 7},
		{time.Date(2022, 11, 1, 0, 0, 0, 0, time.UTC), 7, 7},
		{time.Date(2022, 11, 1, 0, 0, 0, 0, time.UTC), 8, 8},
		{time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), 1, 2},
		{time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC), 1, 1},
		{time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC), 1, 1},
		{time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC), 1, 3},
		{time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC), 1, 1},
		{time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC), 1, 1},
		{time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC), 1, 3},
		{time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC), 1, 1},
		{time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC), 1, 1},
		{time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC), 1, 2},
		{time.Date(2023, 11, 1, 0, 0, 0, 0, time.UTC), 1, 1},
		{time.Date(2023, 12, 1, 0, 0, 0, 0, time.UTC), 1, 1},
		{time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), 31, 31},
		{time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC), 31, 0},
		{time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC), 31, 31},
		{time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC), 31, 0},
		{time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC), 31, 31},
		{time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC), 31, 30},
		{time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC), 31, 31},
		{time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC), 31, 31},
		{time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC), 31, 29},
		{time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC), 31, 31},
		{time.Date(2023, 11, 1, 0, 0, 0, 0, time.UTC), 31, 30},
		{time.Date(2023, 12, 1, 0, 0, 0, 0, time.UTC), 31, 29},
		{time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), 31, 31},
		{time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC), 31, 29},
		{time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC), 31, 29},
		{time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC), 31, 30},
		{time.Date(2024, 5, 1, 0, 0, 0, 0, time.UTC), 31, 31},
		{time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC), 31, 0},
		{time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC), 31, 31},
		{time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC), 31, 30},
		{time.Date(2024, 9, 1, 0, 0, 0, 0, time.UTC), 31, 30},
		{time.Date(2024, 10, 1, 0, 0, 0, 0, time.UTC), 31, 31},
		{time.Date(2024, 11, 1, 0, 0, 0, 0, time.UTC), 31, 29},
		{time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC), 31, 31},

		{time.Date(2022, 11, 30, 0, 0, 0, 0, time.UTC), 3, 3},
		{time.Date(2022, 11, 30, 0, 0, 0, 0, time.UTC), 4, 4},
		{time.Date(2022, 11, 30, 0, 0, 0, 0, time.UTC), 5, 4},
		{time.Date(2022, 11, 30, 0, 0, 0, 0, time.UTC), 6, 7},
		{time.Date(2022, 11, 30, 0, 0, 0, 0, time.UTC), 7, 7},
		{time.Date(2022, 11, 30, 0, 0, 0, 0, time.UTC), 8, 8},
		{time.Date(2023, 1, 31, 0, 0, 0, 0, time.UTC), 1, 2},
		{time.Date(2023, 2, 28, 0, 0, 0, 0, time.UTC), 1, 1},
		{time.Date(2023, 3, 31, 0, 0, 0, 0, time.UTC), 1, 1},
		{time.Date(2023, 4, 30, 0, 0, 0, 0, time.UTC), 1, 3},
		{time.Date(2023, 5, 31, 0, 0, 0, 0, time.UTC), 1, 1},
		{time.Date(2023, 6, 30, 0, 0, 0, 0, time.UTC), 1, 1},
		{time.Date(2023, 7, 31, 0, 0, 0, 0, time.UTC), 1, 3},
		{time.Date(2023, 8, 31, 0, 0, 0, 0, time.UTC), 1, 1},
		{time.Date(2023, 9, 30, 0, 0, 0, 0, time.UTC), 1, 1},
		{time.Date(2023, 10, 31, 0, 0, 0, 0, time.UTC), 1, 2},
		{time.Date(2023, 11, 30, 0, 0, 0, 0, time.UTC), 1, 1},
		{time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC), 1, 1},
		{time.Date(2023, 1, 31, 0, 0, 0, 0, time.UTC), 31, 31},
		{time.Date(2023, 2, 28, 0, 0, 0, 0, time.UTC), 31, 0},
		{time.Date(2023, 3, 31, 0, 0, 0, 0, time.UTC), 31, 31},
		{time.Date(2023, 4, 30, 0, 0, 0, 0, time.UTC), 31, 0},
		{time.Date(2023, 5, 31, 0, 0, 0, 0, time.UTC), 31, 31},
		{time.Date(2023, 6, 30, 0, 0, 0, 0, time.UTC), 31, 30},
		{time.Date(2023, 7, 31, 0, 0, 0, 0, time.UTC), 31, 31},
		{time.Date(2023, 8, 31, 0, 0, 0, 0, time.UTC), 31, 31},
		{time.Date(2023, 9, 30, 0, 0, 0, 0, time.UTC), 31, 29},
		{time.Date(2023, 10, 31, 0, 0, 0, 0, time.UTC), 31, 31},
		{time.Date(2023, 11, 30, 0, 0, 0, 0, time.UTC), 31, 30},
		{time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC), 31, 29},
		{time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC), 31, 31},
		{time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC), 31, 29},
		{time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC), 31, 29},
		{time.Date(2024, 4, 30, 0, 0, 0, 0, time.UTC), 31, 30},
		{time.Date(2024, 5, 31, 0, 0, 0, 0, time.UTC), 31, 31},
		{time.Date(2024, 6, 30, 0, 0, 0, 0, time.UTC), 31, 0},
		{time.Date(2024, 7, 31, 0, 0, 0, 0, time.UTC), 31, 31},
		{time.Date(2024, 8, 31, 0, 0, 0, 0, time.UTC), 31, 30},
		{time.Date(2024, 9, 30, 0, 0, 0, 0, time.UTC), 31, 30},
		{time.Date(2024, 10, 31, 0, 0, 0, 0, time.UTC), 31, 31},
		{time.Date(2024, 11, 30, 0, 0, 0, 0, time.UTC), 31, 29},
		{time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC), 31, 31},
	}

	for _, t := range tt {
		assert.Equal(t.expected, util.NearestWeekday(t.tm, t.day), fmt.Sprintf("%s %v", t.tm, t))
	}
}

func TestNthDayOfWeek(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		tm       time.Time
		w        time.Weekday
		nth      int
		expected int
	}{
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Monday, 1, 3},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Monday, 2, 10},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Monday, 3, 17},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Monday, 4, 24},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Monday, 5, 31},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Tuesday, 1, 4},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Tuesday, 2, 11},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Tuesday, 3, 18},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Tuesday, 4, 25},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Wednesday, 1, 5},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Wednesday, 2, 12},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Wednesday, 3, 19},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Wednesday, 4, 26},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Thursday, 1, 6},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Thursday, 2, 13},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Thursday, 3, 20},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Thursday, 4, 27},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Friday, 1, 7},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Friday, 2, 14},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Friday, 3, 21},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Friday, 4, 28},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Saturday, 1, 1},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Saturday, 2, 8},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Saturday, 3, 15},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Saturday, 4, 22},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Saturday, 5, 29},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Sunday, 1, 2},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Sunday, 2, 9},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Sunday, 3, 16},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Sunday, 4, 23},
		{time.Date(2022, 10, 3, 9, 0, 0, 0, time.UTC), time.Sunday, 5, 30},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Monday, 1, 7},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Monday, 2, 14},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Monday, 3, 21},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Monday, 4, 28},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Tuesday, 1, 1},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Tuesday, 2, 8},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Tuesday, 3, 15},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Tuesday, 4, 22},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Tuesday, 5, 29},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Wednesday, 1, 2},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Wednesday, 2, 9},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Wednesday, 3, 16},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Wednesday, 4, 23},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Wednesday, 5, 30},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Thursday, 1, 3},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Thursday, 2, 10},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Thursday, 3, 17},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Thursday, 4, 24},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Friday, 1, 4},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Friday, 2, 11},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Friday, 3, 18},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Friday, 4, 25},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Saturday, 1, 5},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Saturday, 2, 12},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Saturday, 3, 19},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Saturday, 4, 26},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Sunday, 1, 6},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Sunday, 2, 13},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Sunday, 3, 20},
		{time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC), time.Sunday, 4, 27},
		{time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC), time.Sunday, 5, 0},
		{time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC), time.Thursday, 5, 29},
	}

	for _, t := range tt {
		assert.Equal(t.expected, util.NthDayOfWeek(t.tm, t.w, t.nth), fmt.Sprintf("%s %v", t.tm, t))
	}
}

func TestLastWeekdayOfMonth(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		tm       time.Time
		expected int
	}{
		{time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2023, 2, 1, 9, 0, 0, 0, time.UTC), 28},
		{time.Date(2023, 3, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2023, 4, 1, 9, 0, 0, 0, time.UTC), 28},
		{time.Date(2023, 5, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2023, 6, 1, 9, 0, 0, 0, time.UTC), 30},
		{time.Date(2023, 7, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2023, 8, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2023, 9, 1, 9, 0, 0, 0, time.UTC), 29},
		{time.Date(2023, 10, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2023, 11, 1, 9, 0, 0, 0, time.UTC), 30},
		{time.Date(2023, 12, 1, 9, 0, 0, 0, time.UTC), 29},
		{time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2024, 2, 1, 9, 0, 0, 0, time.UTC), 29},
		{time.Date(2024, 3, 1, 9, 0, 0, 0, time.UTC), 29},
		{time.Date(2024, 4, 1, 9, 0, 0, 0, time.UTC), 30},
		{time.Date(2024, 5, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2024, 6, 1, 9, 0, 0, 0, time.UTC), 28},
		{time.Date(2024, 7, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2024, 8, 1, 9, 0, 0, 0, time.UTC), 30},
		{time.Date(2024, 9, 1, 9, 0, 0, 0, time.UTC), 30},
		{time.Date(2024, 10, 1, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2024, 11, 1, 9, 0, 0, 0, time.UTC), 29},
		{time.Date(2024, 12, 1, 9, 0, 0, 0, time.UTC), 31},

		{time.Date(2023, 1, 31, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2023, 2, 28, 9, 0, 0, 0, time.UTC), 28},
		{time.Date(2023, 3, 31, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2023, 4, 30, 9, 0, 0, 0, time.UTC), 28},
		{time.Date(2023, 5, 31, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2023, 6, 30, 9, 0, 0, 0, time.UTC), 30},
		{time.Date(2023, 7, 31, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2023, 8, 31, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2023, 9, 30, 9, 0, 0, 0, time.UTC), 29},
		{time.Date(2023, 10, 31, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2023, 11, 30, 9, 0, 0, 0, time.UTC), 30},
		{time.Date(2023, 12, 31, 9, 0, 0, 0, time.UTC), 29},
		{time.Date(2024, 1, 31, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2024, 2, 29, 9, 0, 0, 0, time.UTC), 29},
		{time.Date(2024, 3, 31, 9, 0, 0, 0, time.UTC), 29},
		{time.Date(2024, 4, 30, 9, 0, 0, 0, time.UTC), 30},
		{time.Date(2024, 5, 31, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2024, 6, 30, 9, 0, 0, 0, time.UTC), 28},
		{time.Date(2024, 7, 31, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2024, 8, 31, 9, 0, 0, 0, time.UTC), 30},
		{time.Date(2024, 9, 30, 9, 0, 0, 0, time.UTC), 30},
		{time.Date(2024, 10, 31, 9, 0, 0, 0, time.UTC), 31},
		{time.Date(2024, 11, 30, 9, 0, 0, 0, time.UTC), 29},
		{time.Date(2024, 12, 31, 9, 0, 0, 0, time.UTC), 31},
	}

	for _, t := range tt {
		assert.Equal(t.expected, util.LastWeekdayOfMonth(t.tm), t.tm)
	}
}

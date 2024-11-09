package cronplan_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/cronplan"
)

func TestIntegrationMatch(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		exp   string
		tests []struct {
			tm       time.Time
			expected bool
		}
	}{
		// https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html
		{
			exp: "0 10 * * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 1, 9, 10, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 10, 10, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 10, 0, 0, 0, time.UTC), true},
			},
		},
		{
			exp: "0 18 ? * MON-FRI *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 1, 18, 0, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 2, 18, 0, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 3, 18, 0, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 4, 18, 0, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 5, 18, 0, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 6, 18, 0, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 7, 18, 0, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 8, 18, 0, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 9, 18, 0, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "0 8 1 * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 8, 1, 8, 0, 0, 0, time.UTC), true},
				{time.Date(2022, 8, 2, 8, 0, 0, 0, time.UTC), false},
				{time.Date(2022, 9, 1, 8, 0, 0, 0, time.UTC), true},
				{time.Date(2022, 9, 2, 8, 0, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "0/15 * * * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 8, 1, 8, 0, 0, 0, time.UTC), true},
				{time.Date(2022, 8, 1, 8, 5, 0, 0, time.UTC), false},
				{time.Date(2022, 8, 1, 8, 10, 0, 0, time.UTC), false},
				{time.Date(2022, 8, 1, 8, 15, 0, 0, time.UTC), true},
				{time.Date(2022, 8, 1, 8, 20, 0, 0, time.UTC), false},
				{time.Date(2022, 8, 1, 8, 25, 0, 0, time.UTC), false},
				{time.Date(2022, 8, 1, 8, 30, 0, 0, time.UTC), true},
				{time.Date(2022, 8, 1, 8, 35, 0, 0, time.UTC), false},
				{time.Date(2022, 8, 1, 8, 40, 0, 0, time.UTC), false},
				{time.Date(2022, 8, 1, 8, 45, 0, 0, time.UTC), true},
				{time.Date(2022, 8, 1, 8, 50, 0, 0, time.UTC), false},
				{time.Date(2022, 8, 1, 8, 55, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "0/5 8-17 ? * MON-FRI *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 8, 1, 8, 5, 0, 0, time.UTC), true},
				{time.Date(2022, 8, 1, 8, 10, 0, 0, time.UTC), true},
				{time.Date(2022, 8, 1, 8, 18, 0, 0, time.UTC), false},
				{time.Date(2022, 8, 2, 8, 5, 0, 0, time.UTC), true},
				{time.Date(2022, 8, 2, 8, 10, 0, 0, time.UTC), true},
				{time.Date(2022, 8, 2, 8, 18, 0, 0, time.UTC), false},
				{time.Date(2022, 8, 3, 8, 5, 0, 0, time.UTC), true},
				{time.Date(2022, 8, 3, 8, 10, 0, 0, time.UTC), true},
				{time.Date(2022, 8, 3, 8, 18, 0, 0, time.UTC), false},
				{time.Date(2022, 8, 4, 8, 5, 0, 0, time.UTC), true},
				{time.Date(2022, 8, 4, 8, 10, 0, 0, time.UTC), true},
				{time.Date(2022, 8, 4, 8, 18, 0, 0, time.UTC), false},
				{time.Date(2022, 8, 5, 8, 5, 0, 0, time.UTC), true},
				{time.Date(2022, 8, 5, 8, 10, 0, 0, time.UTC), true},
				{time.Date(2022, 8, 5, 8, 18, 0, 0, time.UTC), false},
				{time.Date(2022, 8, 6, 8, 5, 0, 0, time.UTC), false},
				{time.Date(2022, 8, 6, 8, 10, 0, 0, time.UTC), false},
				{time.Date(2022, 8, 6, 8, 15, 0, 0, time.UTC), false},
				{time.Date(2022, 8, 7, 8, 5, 0, 0, time.UTC), false},
				{time.Date(2022, 8, 7, 8, 10, 0, 0, time.UTC), false},
				{time.Date(2022, 8, 7, 8, 15, 0, 0, time.UTC), false},
			},
		},
		// https://docs.oracle.com/cd/E12058_01/doc/doc.1014/e12030/cron_expression.htm
		{
			exp: "0 12 * * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 8, 1, 12, 0, 0, 0, time.UTC), true},
				{time.Date(2022, 8, 1, 12, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "15 10 ? * * *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 8, 1, 10, 15, 0, 0, time.UTC), true},
				{time.Date(2022, 8, 1, 10, 16, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "15 10 * * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 8, 1, 10, 15, 0, 0, time.UTC), true},
				{time.Date(2022, 8, 1, 10, 16, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "15 10 * * ? 2005",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2005, 8, 1, 10, 15, 0, 0, time.UTC), true},
				{time.Date(2022, 8, 1, 10, 15, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* 14 * * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 8, 1, 14, 0, 0, 0, time.UTC), true},
				{time.Date(2022, 8, 1, 14, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 8, 1, 15, 0, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "0/5 14 * * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 8, 1, 14, 0, 0, 0, time.UTC), true},
				{time.Date(2022, 8, 1, 14, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 8, 1, 14, 2, 0, 0, time.UTC), false},
				{time.Date(2022, 8, 1, 14, 3, 0, 0, time.UTC), false},
				{time.Date(2022, 8, 1, 14, 4, 0, 0, time.UTC), false},
				{time.Date(2022, 8, 1, 14, 5, 0, 0, time.UTC), true},
				{time.Date(2022, 8, 1, 14, 6, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "0/5 14,18 * * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 8, 1, 14, 0, 0, 0, time.UTC), true},
				{time.Date(2022, 8, 1, 14, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 8, 1, 18, 4, 0, 0, time.UTC), false},
				{time.Date(2022, 8, 1, 18, 5, 0, 0, time.UTC), true},
				{time.Date(2022, 8, 1, 15, 0, 0, 0, time.UTC), false},
				{time.Date(2022, 8, 1, 15, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "0-5 14 * * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 8, 1, 14, 0, 0, 0, time.UTC), true},
				{time.Date(2022, 8, 1, 14, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 8, 1, 14, 2, 0, 0, time.UTC), true},
				{time.Date(2022, 8, 1, 14, 3, 0, 0, time.UTC), true},
				{time.Date(2022, 8, 1, 14, 4, 0, 0, time.UTC), true},
				{time.Date(2022, 8, 1, 14, 5, 0, 0, time.UTC), true},
				{time.Date(2022, 8, 1, 14, 6, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "10,44 14 ? 3 WED *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 3, 2, 14, 10, 0, 0, time.UTC), true},
				{time.Date(2022, 3, 2, 14, 44, 0, 0, time.UTC), true},
				{time.Date(2022, 3, 2, 14, 11, 0, 0, time.UTC), false},
				{time.Date(2022, 3, 2, 14, 45, 0, 0, time.UTC), false},
				{time.Date(2022, 3, 3, 14, 10, 0, 0, time.UTC), false},
				{time.Date(2022, 3, 3, 14, 44, 0, 0, time.UTC), false},
				{time.Date(2022, 3, 9, 14, 10, 0, 0, time.UTC), true},
				{time.Date(2022, 3, 9, 14, 44, 0, 0, time.UTC), true},
			},
		},
		{
			exp: "15 10 ? * MON-FRI *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 2, 10, 15, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 3, 10, 15, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 4, 10, 15, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 5, 10, 15, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 6, 10, 15, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 7, 10, 15, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 8, 10, 15, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "15 10 15 * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 15, 10, 15, 0, 0, time.UTC), true},
				{time.Date(2022, 11, 15, 10, 15, 0, 0, time.UTC), true},
				{time.Date(2022, 11, 16, 10, 15, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "15 10 L * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 9, 30, 10, 15, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 31, 10, 15, 0, 0, time.UTC), true},
				{time.Date(2022, 11, 30, 10, 15, 0, 0, time.UTC), true},
				{time.Date(2022, 11, 29, 10, 15, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "15 10 ? * 6#3 *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 9, 16, 10, 15, 0, 0, time.UTC), true},
				{time.Date(2022, 9, 23, 10, 15, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 21, 10, 15, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 20, 10, 15, 0, 0, time.UTC), false},
				{time.Date(2022, 11, 18, 10, 15, 0, 0, time.UTC), true},
				{time.Date(2022, 11, 15, 10, 15, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "0 12 1/5 * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 1, 12, 0, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 2, 12, 0, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 3, 12, 0, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 4, 12, 0, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 5, 12, 0, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 6, 12, 0, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 7, 12, 0, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "11 11 11 11 ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2021, 11, 11, 11, 11, 0, 0, time.UTC), true},
				{time.Date(2022, 11, 11, 11, 11, 0, 0, time.UTC), true},
				{time.Date(2023, 11, 11, 11, 11, 0, 0, time.UTC), true},
				{time.Date(2021, 11, 11, 11, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "0 0 ? * 6L *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2023, 10, 27, 0, 0, 0, 0, time.UTC), true},
				{time.Date(2023, 11, 24, 0, 0, 0, 0, time.UTC), true},
				{time.Date(2023, 12, 29, 0, 0, 0, 0, time.UTC), true},
				{time.Date(2023, 10, 20, 0, 0, 0, 0, time.UTC), false},
				{time.Date(2023, 11, 17, 0, 0, 0, 0, time.UTC), false},
				{time.Date(2023, 12, 22, 0, 0, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "0 0 ? * 7L *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2023, 10, 28, 0, 0, 0, 0, time.UTC), true},
				{time.Date(2023, 11, 25, 0, 0, 0, 0, time.UTC), true},
				{time.Date(2023, 12, 30, 0, 0, 0, 0, time.UTC), true},
				{time.Date(2023, 10, 21, 0, 0, 0, 0, time.UTC), false},
				{time.Date(2023, 11, 18, 0, 0, 0, 0, time.UTC), false},
				{time.Date(2023, 12, 21, 0, 0, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "0 0 ? * L *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2023, 10, 7, 0, 0, 0, 0, time.UTC), true},
				{time.Date(2023, 10, 14, 0, 0, 0, 0, time.UTC), true},
				{time.Date(2023, 10, 28, 0, 0, 0, 0, time.UTC), true},
				{time.Date(2023, 11, 4, 0, 0, 0, 0, time.UTC), true},
				{time.Date(2023, 11, 11, 0, 0, 0, 0, time.UTC), true},
				{time.Date(2023, 11, 18, 0, 0, 0, 0, time.UTC), true},
				{time.Date(2023, 11, 25, 0, 0, 0, 0, time.UTC), true},
				{time.Date(2023, 12, 2, 0, 0, 0, 0, time.UTC), true},
				{time.Date(2023, 12, 9, 0, 0, 0, 0, time.UTC), true},
				{time.Date(2023, 10, 8, 0, 0, 0, 0, time.UTC), false},
				{time.Date(2023, 10, 15, 0, 0, 0, 0, time.UTC), false},
				{time.Date(2023, 10, 29, 0, 0, 0, 0, time.UTC), false},
				{time.Date(2023, 11, 3, 0, 0, 0, 0, time.UTC), false},
				{time.Date(2023, 11, 10, 0, 0, 0, 0, time.UTC), false},
				{time.Date(2023, 11, 17, 0, 0, 0, 0, time.UTC), false},
				{time.Date(2023, 11, 24, 0, 0, 0, 0, time.UTC), false},
				{time.Date(2023, 12, 1, 0, 0, 0, 0, time.UTC), false},
			},
		},
	}

	for _, t := range tt {
		cron, err := cronplan.Parser.ParseString("", t.exp)
		assert.NoError(err)

		for _, t2 := range t.tests {
			assert.Equal(t2.expected, cron.Match(t2.tm), t.exp, t2.tm)
		}
	}
}

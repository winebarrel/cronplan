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
	}

	for _, t := range tt {
		cron, err := cronplan.Parse(t.exp)
		assert.NoError(err)

		for _, t2 := range t.tests {
			assert.Equal(t2.expected, cron.Match(t2.tm), t.exp, t2.tm)
		}
	}
}

func TestMatchHours(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		exp   string
		tests []struct {
			tm       time.Time
			expected bool
		}
	}{
		{
			exp: "* 2 * * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 2, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 3, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* 2,3 * * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 2, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 3, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 4, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* 2-4 * * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 2, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 3, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 4, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 5, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* */3 * * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 1, 0, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 1, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 2, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 3, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 4, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 5, 1, 0, 0, time.UTC), false},
			},
		},
		{
			exp: "* 1/3 * * ? *",
			tests: []struct {
				tm       time.Time
				expected bool
			}{
				{time.Date(2022, 10, 1, 0, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 1, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 2, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 3, 1, 0, 0, time.UTC), false},
				{time.Date(2022, 10, 1, 4, 1, 0, 0, time.UTC), true},
				{time.Date(2022, 10, 1, 5, 1, 0, 0, time.UTC), false},
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
	}

	for _, t := range tt {
		cron, err := cronplan.Parse(t.exp)
		assert.NoError(err)

		for _, t2 := range t.tests {
			assert.Equal(t2.expected, cron.Match(t2.tm), t.exp, t2.tm)
		}
	}
}

func TestMatchMonth(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		exp   string
		tests []struct {
			tm       time.Time
			expected bool
		}
	}{
		// {
		// 	exp: "* * * 2 ? *",
		// 	tests: []struct {
		// 		tm       time.Time
		// 		expected bool
		// 	}{
		// 		{time.Date(2022, 1, 1, 1, 1, 0, 0, time.UTC), false},
		// 		{time.Date(2022, 2, 1, 1, 1, 0, 0, time.UTC), false},
		// 		{time.Date(2022, 3, 1, 1, 1, 0, 0, time.UTC), true},
		// 		// {time.Date(2023, 2, 1, 1, 1, 0, 0, time.UTC), true},
		// 	},
		// },
		// {
		// 	exp: "* * * 2,3 ? *",
		// 	tests: []struct {
		// 		tm       time.Time
		// 		expected bool
		// 	}{
		// 		{time.Date(2022, 1, 1, 1, 1, 0, 0, time.UTC), false},
		// 		{time.Date(2022, 2, 1, 1, 1, 0, 0, time.UTC), true},
		// 		{time.Date(2022, 3, 1, 1, 1, 0, 0, time.UTC), true},
		// 		{time.Date(2022, 4, 1, 1, 1, 0, 0, time.UTC), false},
		// 		{time.Date(2023, 2, 1, 1, 1, 0, 0, time.UTC), true},
		// 	},
		// },
		// {
		// 	exp: "* * * 2-4 ? *",
		// 	tests: []struct {
		// 		tm       time.Time
		// 		expected bool
		// 	}{
		// 		{time.Date(2022, 1, 1, 1, 1, 0, 0, time.UTC), false},
		// 		{time.Date(2022, 2, 1, 1, 1, 0, 0, time.UTC), true},
		// 		{time.Date(2022, 3, 1, 1, 1, 0, 0, time.UTC), true},
		// 		{time.Date(2022, 4, 1, 1, 1, 0, 0, time.UTC), true},
		// 		{time.Date(2022, 5, 1, 1, 1, 0, 0, time.UTC), false},
		// 		{time.Date(2023, 2, 1, 1, 1, 0, 0, time.UTC), true},
		// 	},
		// },
		// {
		// 	exp: "* * * */3 ? *",
		// 	tests: []struct {
		// 		tm       time.Time
		// 		expected bool
		// 	}{
		// 		{time.Date(2022, 1, 1, 1, 1, 0, 0, time.UTC), true},
		// 		{time.Date(2022, 2, 1, 1, 1, 0, 0, time.UTC), false},
		// 		{time.Date(2022, 3, 1, 1, 1, 0, 0, time.UTC), false},
		// 		{time.Date(2022, 4, 1, 1, 1, 0, 0, time.UTC), true},
		// 		{time.Date(2022, 5, 1, 1, 1, 0, 0, time.UTC), false},
		// 		{time.Date(2022, 6, 1, 1, 1, 0, 0, time.UTC), false},
		// 	},
		// },
		// {
		// 	exp: "* * * 1/3 ? *",
		// 	tests: []struct {
		// 		tm       time.Time
		// 		expected bool
		// 	}{
		// 		{time.Date(2022, 1, 1, 1, 1, 0, 0, time.UTC), true},
		// 		{time.Date(2022, 2, 1, 1, 1, 0, 0, time.UTC), false},
		// 		{time.Date(2022, 3, 1, 1, 1, 0, 0, time.UTC), false},
		// 		{time.Date(2022, 4, 1, 1, 1, 0, 0, time.UTC), true},
		// 		{time.Date(2022, 5, 1, 1, 1, 0, 0, time.UTC), false},
		// 		{time.Date(2022, 6, 1, 1, 1, 0, 0, time.UTC), false},
		// 	},
		// },
		// {

		// 	exp: "* * * FEB,MAR ? *",
		// 	tests: []struct {
		// 		tm       time.Time
		// 		expected bool
		// 	}{
		// 		{time.Date(2022, 1, 1, 1, 1, 0, 0, time.UTC), false},
		// 		{time.Date(2022, 2, 1, 1, 1, 0, 0, time.UTC), true},
		// 		{time.Date(2022, 3, 1, 1, 1, 0, 0, time.UTC), true},
		// 		{time.Date(2022, 4, 1, 1, 1, 0, 0, time.UTC), false},
		// 		{time.Date(2023, 2, 1, 1, 1, 0, 0, time.UTC), true},
		// 	},
		// },
		// {
		// 	exp: "* * * FEB-APR ? *",
		// 	tests: []struct {
		// 		tm       time.Time
		// 		expected bool
		// 	}{
		// 		{time.Date(2022, 1, 1, 1, 1, 0, 0, time.UTC), false},
		// 		{time.Date(2022, 2, 1, 1, 1, 0, 0, time.UTC), true},
		// 		{time.Date(2022, 3, 1, 1, 1, 0, 0, time.UTC), true},
		// 		{time.Date(2022, 4, 1, 1, 1, 0, 0, time.UTC), true},
		// 		{time.Date(2022, 5, 1, 1, 1, 0, 0, time.UTC), false},
		// 		{time.Date(2023, 2, 1, 1, 1, 0, 0, time.UTC), true},
		// 	},
		// },
	}

	for _, t := range tt {
		cron, err := cronplan.Parse(t.exp)
		assert.NoError(err)

		for _, t2 := range t.tests {
			assert.Equal(t2.expected, cron.Match(t2.tm), t.exp, t2.tm)
		}
	}
}

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
	}

	for _, t := range tt {
		cron, err := cronplan.Parse(t.exp)
		assert.NoError(err)

		for _, t2 := range t.tests {
			assert.Equal(t2.expected, cron.Match(t2.tm), t.exp, t2.tm)
		}
	}
}

// func TestMatchYear(t *testing.T) {
// 	assert := assert.New(t)

// 	tt := []struct {
// 		exp   string
// 		tests []struct {
// 			tm       time.Time
// 			expected bool
// 		}
// 	}{
// 		{
// 			exp: "* * * * ? 2022",
// 			tests: []struct {
// 				tm       time.Time
// 				expected bool
// 			}{
// 				{time.Date(2021, 10, 1, 1, 1, 0, 0, time.UTC), false},
// 				{time.Date(2022, 10, 1, 1, 1, 0, 0, time.UTC), true},
// 				{time.Date(2023, 10, 1, 1, 1, 0, 0, time.UTC), false},
// 			},
// 		},
// 		{
// 			exp: "* * * * ? 2022,2023",
// 			tests: []struct {
// 				tm       time.Time
// 				expected bool
// 			}{
// 				{time.Date(2021, 10, 1, 1, 1, 0, 0, time.UTC), false},
// 				{time.Date(2022, 10, 1, 1, 1, 0, 0, time.UTC), true},
// 				{time.Date(2023, 10, 1, 1, 1, 0, 0, time.UTC), true},
// 				{time.Date(2024, 10, 1, 1, 1, 0, 0, time.UTC), false},
// 			},
// 		},
// 		{
// 			exp: "* * * * ? 2022-2024",
// 			tests: []struct {
// 				tm       time.Time
// 				expected bool
// 			}{
// 				{time.Date(2021, 10, 1, 1, 1, 0, 0, time.UTC), false},
// 				{time.Date(2022, 10, 1, 1, 1, 0, 0, time.UTC), true},
// 				{time.Date(2023, 10, 1, 1, 1, 0, 0, time.UTC), true},
// 				{time.Date(2024, 10, 1, 1, 1, 0, 0, time.UTC), true},
// 				{time.Date(2025, 10, 1, 1, 1, 0, 0, time.UTC), false},
// 			},
// 		},
// 		{
// 			exp: "* * * * ? */11",
// 			tests: []struct {
// 				tm       time.Time
// 				expected bool
// 			}{
// 				{time.Date(1959, 10, 1, 1, 0, 0, 0, time.UTC), false},
// 				{time.Date(1970, 10, 1, 1, 0, 0, 0, time.UTC), true},
// 				{time.Date(2024, 10, 1, 1, 0, 0, 0, time.UTC), false},
// 				{time.Date(2025, 10, 1, 1, 1, 0, 0, time.UTC), true},
// 				{time.Date(2026, 10, 1, 1, 2, 0, 0, time.UTC), false},
// 				{time.Date(2036, 10, 1, 1, 3, 0, 0, time.UTC), true},
// 				{time.Date(2047, 10, 1, 1, 4, 0, 0, time.UTC), true},
// 			},
// 		},
// 		{
// 			exp: "* * * * ? 2024/11",
// 			tests: []struct {
// 				tm       time.Time
// 				expected bool
// 			}{
// 				{time.Date(2013, 10, 1, 1, 0, 0, 0, time.UTC), false},
// 				{time.Date(2023, 10, 1, 1, 0, 0, 0, time.UTC), false},
// 				{time.Date(2024, 10, 1, 1, 0, 0, 0, time.UTC), true},
// 				{time.Date(2025, 10, 1, 1, 1, 0, 0, time.UTC), false},
// 				{time.Date(2035, 10, 1, 1, 3, 0, 0, time.UTC), true},
// 				{time.Date(2046, 10, 1, 1, 4, 0, 0, time.UTC), true},
// 			},
// 		},
// 	}

// 	for _, t := range tt {
// 		cron, err := cronplan.Parse( t.exp)
// 		assert.NoError(err)

// 		for _, t2 := range t.tests {
// 			assert.Equal(t2.expected, cron.Match(t2.tm), t.exp, t2.tm)
// 		}
// 	}
// }

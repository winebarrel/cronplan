package cronplan_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/cronplan/v2"
)

func TestParseErrors(t *testing.T) {
	tt := []struct {
		name   string
		exp    string
		errSub string
	}{
		{
			name:   "both day-of-month and day-of-week are question mark",
			exp:    "* * ? * ? *",
			errSub: "'?' cannot be set to both day-of-month and day-of-week",
		},
		{
			name:   "neither day-of-month nor day-of-week is question mark",
			exp:    "* * * * * *",
			errSub: "either day-of-month or day-of-week must be '?'",
		},
		{
			name:   "minute out of range",
			exp:    "60 * * * ? *",
			errSub: "minute must be 0-59 (value=60)",
		},
		{
			name:   "hour out of range",
			exp:    "* 24 * * ? *",
			errSub: "hour must be 0-23 (value=24)",
		},
		{
			name:   "day-of-month out of range",
			exp:    "* * 32 * ? *",
			errSub: "day-of-month must be 1-31 (value=32)",
		},
		{
			name:   "month out of range",
			exp:    "* * * 13 ? *",
			errSub: "month number must be 1-12 (value=13)",
		},
		{
			name:   "day-of-week out of range",
			exp:    "* * * * 8 *",
			errSub: "day-of-week number must be 1-7 (value=8)",
		},
		{
			name:   "year below minimum",
			exp:    "* * * * ? 1969",
			errSub: "year must be 1970-2199 (value=1969)",
		},
		{
			name:   "year above maximum",
			exp:    "* * * * ? 2200",
			errSub: "year must be 1970-2199 (value=2200)",
		},
		{
			name:   "empty expression",
			exp:    "",
			errSub: "unexpected token",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			_, err := cronplan.Parse(tc.exp)
			assert.Error(err)
			assert.ErrorContains(err, tc.errSub)
		})
	}
}

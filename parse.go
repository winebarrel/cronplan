//nolint:govet
package cronplan

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/winebarrel/cronplan/internal/util"
)

var (
	cronLexer = lexer.MustSimple([]lexer.SimpleRule{
		{Name: `Number`, Pattern: `\d+`},
		{Name: `Month`, Pattern: `(?i)(?:` + strings.Join(util.ShortMonthNames, "|") + `)`},
		{Name: `Weekday`, Pattern: `(?i)(?:` + strings.Join(util.ShortWeekdayNames, "|") + `)`},
		{Name: `Symbol`, Pattern: `[,\-\*\?/LW#]`},
		{Name: `SP`, Pattern: `\s+`},
	})

	Parser = participle.MustBuild[Expression](
		participle.Lexer(cronLexer),
	)
)

// minute =====================================================================

type Minute int

func (v *Minute) Capture(values []string) error {
	s := values[0]
	r := regexp.MustCompile(`^\d+$`)

	if !r.MatchString(s) {
		return fmt.Errorf("connot convert to minute from %s", s)
	}

	n, _ := strconv.Atoi(s)

	if n < 0 || 59 < n {
		return fmt.Errorf("minute must be 0-59 (value=%d)", n)
	}

	*v = Minute(n)

	return nil
}

func (v *Minute) Int() int {
	return int(*v)
}

func (v *Minute) String() string {
	return strconv.Itoa(v.Int())
}

type MinuteRange struct {
	Start *Minute `@Number`
	End   *Minute `"-" @Number`
}

func (v *MinuteRange) String() string {
	return fmt.Sprintf("%s-%s", v.Start, v.End)
}

type MinuteExp struct {
	Wildcard bool         `( @"*"`
	Range    *MinuteRange `  | @@`
	Number   *Minute      `  | @Number )`
	Bottom   *int         `( "/" @Number )?`
}

func (e *MinuteExp) String() string {
	var s string

	if e.Wildcard {
		s = "*"
	} else if e.Range != nil {
		s = e.Range.String()
	} else if e.Number != nil {
		s = e.Number.String()
	}

	if e.Bottom != nil {
		s = fmt.Sprintf("%s/%d", s, *e.Bottom)
	}

	return s
}

type MinuteField struct {
	Exps []*MinuteExp `@@ ( "," @@ )*`
}

func (v *MinuteField) String() string {
	ss := make([]string, 0, len(v.Exps))

	for _, e := range v.Exps {
		ss = append(ss, e.String())
	}

	return strings.Join(ss, ",")
}

// hour =======================================================================

type Hour int

func (v *Hour) Capture(values []string) error {
	s := values[0]
	r := regexp.MustCompile(`^\d+$`)

	if !r.MatchString(s) {
		return fmt.Errorf("connot convert to hour from %s", s)
	}

	n, _ := strconv.Atoi(s)

	if n < 0 || 23 < n {
		return fmt.Errorf("hour must be 0-23 (value=%d)", n)
	}

	*v = Hour(n)

	return nil
}

func (v *Hour) Int() int {
	return int(*v)
}

func (v *Hour) String() string {
	return strconv.Itoa(v.Int())
}

type HourRange struct {
	Start *Hour `@Number`
	End   *Hour `"-" @Number`
}

func (v *HourRange) String() string {
	return fmt.Sprintf("%s-%s", v.Start, v.End)
}

type HourExp struct {
	Wildcard bool       `( @"*"`
	Range    *HourRange `  | @@`
	Number   *Hour      `  | @Number )`
	Bottom   *int       `( "/" @Number )?`
}

func (e *HourExp) String() string {
	var s string

	if e.Wildcard {
		s = "*"
	} else if e.Range != nil {
		s = e.Range.String()
	} else if e.Number != nil {
		s = e.Number.String()
	}

	if e.Bottom != nil {
		s = fmt.Sprintf("%s/%d", s, *e.Bottom)
	}

	return s
}

type HourField struct {
	Exps []*HourExp `@@ ( "," @@ )*`
}

func (v *HourField) String() string {
	ss := make([]string, 0, len(v.Exps))

	for _, e := range v.Exps {
		ss = append(ss, e.String())
	}

	return strings.Join(ss, ",")
}

// day-of-month ===============================================================

type DayOfMonth int

func (v *DayOfMonth) Capture(values []string) error {
	s := values[0]
	r := regexp.MustCompile(`^\d+$`)

	if !r.MatchString(s) {
		return fmt.Errorf("connot convert to day-of-month from %s", s)
	}

	n, _ := strconv.Atoi(s)

	if n < 1 || 31 < n {
		return fmt.Errorf("hour must be 1-31 (value=%d)", n)
	}

	*v = DayOfMonth(n)

	return nil
}

func (v *DayOfMonth) Int() int {
	return int(*v)
}

func (v *DayOfMonth) String() string {
	return strconv.Itoa(v.Int())
}

type DayOfMonthRange struct {
	Start *DayOfMonth `@Number`
	End   *DayOfMonth `"-" @Number`
}

func (v *DayOfMonthRange) String() string {
	return fmt.Sprintf("%s-%s", v.Start, v.End)
}

type NearestWeekday int

func (v *NearestWeekday) Capture(values []string) error {
	s := values[0]
	r := regexp.MustCompile(`^\d+$`)

	if !r.MatchString(s) {
		return fmt.Errorf("connot convert to nearest_weekday from %sW", s)
	}

	n, _ := strconv.Atoi(s)

	if n < 1 || 31 < n {
		return fmt.Errorf("'<num>W' must be 1-31 (value=%d)", n)
	}

	*v = NearestWeekday(n)

	return nil
}

func (v *NearestWeekday) Int() int {
	return int(*v)
}

func (v *NearestWeekday) String() string {
	return fmt.Sprintf("%dW", v.Int())
}

type LastDayOfMonth int

func (v *LastDayOfMonth) Capture(values []string) error {
	s := values[0]

	if s == "L" {
		*v = 0
		return nil
	}

	r := regexp.MustCompile(`^\d+$`)

	if !r.MatchString(s) {
		return fmt.Errorf("connot convert to last_day-of-month from L-%s", s)
	}

	n, _ := strconv.Atoi(s)

	if n < 1 || 30 < n {
		return fmt.Errorf("'L-<num>' must be 1-30 (value=%d)", n)
	}

	*v = LastDayOfMonth(n)

	return nil
}

func (v *LastDayOfMonth) Int() int {
	return int(*v)
}

func (v *LastDayOfMonth) String() string {
	if v.Int() > 0 {
		return fmt.Sprintf("L-%d", v.Int())
	} else {
		return "L"
	}
}

type DayOfMonthExp struct {
	NearestWeekday *NearestWeekday  `( @Number "W" )`
	Wildcard       bool             `| ( ( @"*"`
	Range          *DayOfMonthRange `      | @@`
	Number         *DayOfMonth      `      | @Number )`
	Bottom         *int             `    ( "/" @Number )? )`
	Last           *LastDayOfMonth  `| ( @"L" ( "-" @Number )? )`
}

func (e *DayOfMonthExp) String() string {
	var s string

	if e.Wildcard {
		s = "*"
	} else if e.Range != nil {
		s = e.Range.String()
	} else if e.Number != nil {
		s = e.Number.String()
	} else if e.Last != nil {
		s = e.Last.String()
	} else if e.NearestWeekday != nil {
		s = e.NearestWeekday.String()
	}

	if e.Bottom != nil {
		s = fmt.Sprintf("%s/%d", s, *e.Bottom)
	}

	return s
}

type DayOfMonthField struct {
	Exps []*DayOfMonthExp `( @@ ( "," @@ )* )`
	Any  bool             `| @"?"`
}

func (v *DayOfMonthField) String() string {
	if v.Any {
		return "?"
	}

	ss := make([]string, 0, len(v.Exps))

	for _, e := range v.Exps {
		ss = append(ss, e.String())
	}

	return strings.Join(ss, ",")
}

// month ======================================================================

type Month int

func (v *Month) Capture(values []string) error {
	s := values[0]
	r := regexp.MustCompile(`^\d+$`)

	if r.MatchString(s) {
		n, _ := strconv.Atoi(s)

		if n < 1 || 12 < n {
			return fmt.Errorf("month number must be 1-12 (value=%d)", n)
		}

		*v = Month(n)
	} else {
		month, err := util.CastMonth(s)

		if err != nil {
			return err
		}

		*v = Month(month)
	}

	return nil
}

func (v *Month) Int() int {
	return int(*v)
}

func (v *Month) Month() time.Month {
	return time.Month(v.Int())
}

func (v *Month) String() string {
	return util.ShortMonthNames[v.Int()-1]
}

type MonthRange struct {
	Start *Month `( @Number | @Month )`
	End   *Month `"-" ( @Number | @Month )`
}

func (v *MonthRange) String() string {
	return fmt.Sprintf("%s-%s", v.Start, v.End)
}

type MonthExp struct {
	Wildcard bool        `( @"*"`
	Range    *MonthRange `  | @@`
	Month    *Month      `  | ( @Number | @Month ) )`
	Bottom   *int        `( "/" @Number )?`
}

func (e *MonthExp) String() string {
	var s string

	if e.Wildcard {
		s = "*"
	} else if e.Range != nil {
		s = e.Range.String()
	} else if e.Month != nil {
		s = e.Month.String()
	}

	if e.Bottom != nil {
		s = fmt.Sprintf("%s/%d", s, *e.Bottom)
	}

	return s
}

type MonthField struct {
	Exps []*MonthExp `@@ ( "," @@ )*`
}

func (v *MonthField) String() string {
	ss := make([]string, 0, len(v.Exps))

	for _, e := range v.Exps {
		ss = append(ss, e.String())
	}

	return strings.Join(ss, ",")
}

// day-of-week ================================================================

type Weekday time.Weekday

func (v *Weekday) Capture(values []string) error {
	s := values[0]
	r := regexp.MustCompile(`^\d+$`)

	if r.MatchString(s) {
		n, _ := strconv.Atoi(s)

		if n < 1 || 7 < n {
			return fmt.Errorf("day-of-week number must be 1-7 (value=%d)", n)
		}

		*v = Weekday(n - 1)
	} else {
		wday, err := util.CastWeekday(s)

		if err != nil {
			return err
		}

		*v = Weekday(wday)
	}

	return nil
}

func (v *Weekday) Int() int {
	return int(*v)
}

func (v *Weekday) Weekday() time.Weekday {
	return time.Weekday(v.Int())
}

func (v *Weekday) String() string {
	return util.ShortWeekdayNames[v.Int()]
}

type WeekdayRange struct {
	Start *Weekday `( @Number | @Weekday )`
	End   *Weekday `"-" ( @Number | @Weekday )`
}

func (v *WeekdayRange) String() string {
	return fmt.Sprintf("%s-%s", v.Start, v.End)
}

type NthDayOfWeek struct {
	Wday *Weekday `( @Number | @Weekday )`
	Nth  int      `"#" @Number`
}

func (v *NthDayOfWeek) String() string {
	return fmt.Sprintf("%s#%d", v.Wday, v.Nth)
}

type LastDayOfWeek struct {
	// NOTE: It seems that "L" without a number can be entered,
	//		   but the schedule is strange so it is not allowed.
	Wday Weekday `(@Number | @Weekday) "L"`
}

func (v *LastDayOfWeek) Weekday() time.Weekday {
	return time.Weekday(v.Wday)
}

func (v *LastDayOfWeek) String() string {
	return fmt.Sprintf("%sL", v.Wday.String())
}

type DayOfWeekExp struct {
	Nth      *NthDayOfWeek  `@@`
	Last     *LastDayOfWeek `| @@`
	Wildcard bool           `| ( ( @"*"`
	Range    *WeekdayRange  `      | @@`
	Wday     *Weekday       `      | ( @Number | @Weekday ) )`
	Bottom   *int           `    ( "/" @Number )? )`
}

func (e *DayOfWeekExp) String() string {
	var s string

	if e.Wildcard {
		s = "*"
	} else if e.Range != nil {
		s = e.Range.String()
	} else if e.Wday != nil {
		s = e.Wday.String()
	} else if e.Nth != nil {
		s = e.Nth.String()
	} else if e.Last != nil {
		s = e.Last.String()
	}

	if e.Bottom != nil {
		s = fmt.Sprintf("%s/%d", s, *e.Bottom)
	}

	return s
}

type DayOfWeekField struct {
	Exps []*DayOfWeekExp `( @@ ( "," @@ )* )`
	Any  bool            `| @"?"`
}

func (v *DayOfWeekField) String() string {
	if v.Any {
		return "?"
	}

	ss := make([]string, 0, len(v.Exps))

	for _, e := range v.Exps {
		ss = append(ss, e.String())
	}

	return strings.Join(ss, ",")
}

// year =======================================================================

type Year int

func (v *Year) Capture(values []string) error {
	s := values[0]
	r := regexp.MustCompile(`^\d+$`)

	if !r.MatchString(s) {
		return fmt.Errorf("connot convert to year from %s", s)
	}

	n, _ := strconv.Atoi(s)

	if n < 1970 || 2199 < n {
		return fmt.Errorf("year must be 1970-2199 (value=%d)", n)
	}

	*v = Year(n)

	return nil
}

func (v *Year) Int() int {
	return int(*v)
}

func (v *Year) String() string {
	return strconv.Itoa(v.Int())
}

type YearRange struct {
	Start *Year `@Number`
	End   *Year `"-" @Number`
}

func (v *YearRange) String() string {
	return fmt.Sprintf("%s-%s", v.Start, v.End)
}

type YearExp struct {
	Wildcard bool       `( @"*"`
	Range    *YearRange `  | @@`
	Number   *Year      `  | @Number )`
	Bottom   *int       `( "/" @Number )?`
}

func (e *YearExp) String() string {
	var s string

	if e.Wildcard {
		s = "*"
	} else if e.Range != nil {
		s = e.Range.String()
	} else if e.Number != nil {
		s = e.Number.String()
	}

	if e.Bottom != nil {
		s = fmt.Sprintf("%s/%d", s, *e.Bottom)
	}

	return s
}

type YearField struct {
	Exps []*YearExp `@@ ( "," @@ )*`
}

func (v *YearField) String() string {
	ss := make([]string, 0, len(v.Exps))

	for _, e := range v.Exps {
		ss = append(ss, e.String())
	}

	return strings.Join(ss, ",")
}

// expression =================================================================

type Expression struct {
	Minute     *MinuteField     `@@`
	Hour       *HourField       `SP @@`
	DayOfMonth *DayOfMonthField `SP @@`
	Month      *MonthField      `SP @@`
	DayOfWeek  *DayOfWeekField  `SP @@`
	Year       *YearField       `SP @@`
}

func Parse(exp string) (*Expression, error) {
	exp = strings.TrimSpace(exp)
	cron, err := Parser.ParseString("", exp)

	if err != nil {
		return nil, err
	}

	if cron.DayOfMonth.Any && cron.DayOfWeek.Any {
		return nil, fmt.Errorf("'?' cannot be set to both day-of-month and day-of-week")
	} else if !cron.DayOfMonth.Any && !cron.DayOfWeek.Any {
		return nil, fmt.Errorf("either day-of-month or day-of-week must be '?'")
	}

	return cron, nil
}

func (v *Expression) String() string {
	return fmt.Sprintf("%s %s %s %s %s %s",
		v.Minute,
		v.Hour,
		v.DayOfMonth,
		v.Month,
		v.DayOfWeek,
		v.Year,
	)
}

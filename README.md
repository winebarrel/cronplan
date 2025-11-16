# cronplan

[![CI](https://github.com/winebarrel/cronplan/actions/workflows/ci.yml/badge.svg)](https://github.com/winebarrel/cronplan/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/winebarrel/cronplan/v2.svg)](https://pkg.go.dev/github.com/winebarrel/cronplan/v2)
[![Go Report Card](https://goreportcard.com/badge/github.com/winebarrel/cronplan/v2)](https://goreportcard.com/report/github.com/winebarrel/cronplan/v2)

## Overview

Cron expression parser for Amazon EventBridge.

### Try with curl

```sh
$ curl cronplan.in -d '5 0 10 * ? *'
Tue, 10 Oct 2023 00:05:00
Fri, 10 Nov 2023 00:05:00
Sun, 10 Dec 2023 00:05:00
Wed, 10 Jan 2024 00:05:00
Sat, 10 Feb 2024 00:05:00
Sun, 10 Mar 2024 00:05:00
Wed, 10 Apr 2024 00:05:00
Fri, 10 May 2024 00:05:00
Mon, 10 Jun 2024 00:05:00
Wed, 10 Jul 2024 00:05:00
```

## Installation

```sh
go get github.com/winebarrel/cronplan
```

## Usage

```go
package main

import (
	"fmt"
	"time"

	"github.com/winebarrel/cronplan/v2"
)

func main() {
	cron, err := cronplan.Parse("0 10 * * ? *")

	if err != nil {
		panic(err)
	}

	fmt.Println(
		cron.Minute.Exps[0].Number, //=> 0
		cron.Hour.Exps[0].Number,   //=> 10
		cron.String(),              //=> "0 10 * * ? *"
	)

	cron.Match(time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC))
	//=> false
	cron.Match(time.Date(2022, 11, 3, 10, 0, 0, 0, time.UTC))
	//=> true

	// NOTE: If you don't want to include `from`, add `1 * time.Minute`
	cron.Next(time.Date(2022, 11, 3, 10, 0, 0, 0, time.UTC))
	//=> 2022-11-03 10:00:00 +0000 UTC
	cron.Next(time.Date(2022, 11, 3, 11, 0, 0, 0, time.UTC))
	//=> 2022-11-04 10:00:00 +0000 UTC
	cron.NextN(time.Date(2022, 11, 3, 10, 0, 0, 0, time.UTC), 3)
	//=> [2022-11-03 10:00:00 +0000 UTC 2022-11-04 10:00:00 +0000 UTC 2022-11-05 10:00:00 +0000 UTC]

	cron.Between(
		time.Date(2022, 11, 3, 10, 0, 0, 0, time.UTC),
		time.Date(2022, 11, 4, 10, 0, 0, 0, time.UTC),
	)
	//=> [2022-11-03 10:00:00 +0000 UTC 2022-11-04 10:00:00 +0000 UTC]

	iter := cron.Iter(time.Date(2022, 11, 3, 10, 0, 0, 0, time.UTC))

	for i := range 3 {
		fmt.Println(i, iter.Next())
		//=> 0 2022-11-03 10:00:00 +0000 UTC
		//=> 1 2022-11-04 10:00:00 +0000 UTC
		//=> 2 2022-11-05 10:00:00 +0000 UTC
	}

	for next := range iter.Seq() {
		fmt.Println(next)
		//=> 2022-11-03 10:00:00 +0000 UTC
		break
	}
}
```

### Scheduler implementation example

https://github.com/winebarrel/cronplan/blob/main/_example/cron/main.go

## Behavior of "L" in day-of-week

If you specify "L" for day-of-week, the last day of the week of each month is usually matched.

```
# cron(0 0 ? * 6L *)
Fri, 27 Oct 2023 00:00:00
Fri, 24 Nov 2023 00:00:00
Fri, 29 Dec 2023 00:00:00
Fri, 26 Jan 2024 00:00:00
Fri, 23 Feb 2024 00:00:00
```

However, if you do not specify the day of the week before "L", the behavior will be the same as when you specify "SAT".

```
# cron(0 0 ? * L *) = cron(0 0 ? * SAT *)
Sat, 07 Oct 2023 00:00:00
Sat, 14 Oct 2023 00:00:00
Sat, 21 Oct 2023 00:00:00
Sat, 28 Oct 2023 00:00:00
Sat, 04 Nov 2023 00:00:00
```

## Behavior of "31W" in day-of-month

If you specify "31W" for day-of-month, months without a 31st day will be skipped.

(I'm not sure if this is the correct behavior)

```
# cron(5 0 31W * ? 2026)
Fri, 30 Jan 2026 00:05:00
Tue, 31 Mar 2026 00:05:00
Fri, 29 May 2026 00:05:00
Fri, 31 Jul 2026 00:05:00
Mon, 31 Aug 2026 00:05:00
Fri, 30 Oct 2026 00:05:00
Thu, 31 Dec 2026 00:05:00
```

If you want the last weekday of the month, use "LW".

```
# cron(5 0 LW * ? 2026)
Fri, 30 Jan 2026 00:05:00
Fri, 27 Feb 2026 00:05:00
Tue, 31 Mar 2026 00:05:00
Thu, 30 Apr 2026 00:05:00
Fri, 29 May 2026 00:05:00
Tue, 30 Jun 2026 00:05:00
Fri, 31 Jul 2026 00:05:00
Mon, 31 Aug 2026 00:05:00
Wed, 30 Sep 2026 00:05:00
Fri, 30 Oct 2026 00:05:00
```

# cronplan CLI

CLI to show next triggers.

## Installation

```
brew install winebarrel/cronplan/cronplan
```

## Usage

```
Usage: cronplan [OPTION] CRON_EXPR
  -h int
    	hour to add
  -n int
    	number of next triggers (default 10)
  -version
    	print version and exit
```

```
$ cronplan '*/10 10 ? * MON-FRI *'
Tue, 11 Oct 2022 10:00:00
Tue, 11 Oct 2022 10:10:00
Tue, 11 Oct 2022 10:20:00
Tue, 11 Oct 2022 10:30:00
Tue, 11 Oct 2022 10:40:00
Tue, 11 Oct 2022 10:50:00
Wed, 12 Oct 2022 10:00:00
Wed, 12 Oct 2022 10:10:00
Wed, 12 Oct 2022 10:20:00
Wed, 12 Oct 2022 10:30:00

$ cronplan -h -9 '*/10 10 ? * MON-FRI *'
Tue, 11 Oct 2022 01:00:00
Tue, 11 Oct 2022 01:10:00
Tue, 11 Oct 2022 01:20:00
Tue, 11 Oct 2022 01:30:00
Tue, 11 Oct 2022 01:40:00
Tue, 11 Oct 2022 01:50:00
Wed, 12 Oct 2022 01:00:00
Wed, 12 Oct 2022 01:10:00
Wed, 12 Oct 2022 01:20:00
Wed, 12 Oct 2022 01:30:00
```

# cronmatch CLI

CLI to check if datetime matches cron expression.

## Installation

```
brew install winebarrel/cronplan/cronmatch
```

## Usage

```
Usage: cronmatch [OPTION] CRON_EXPR DATE
  -h int
    	hour to add
  -no-color
    	disable color output
  -version
    	print version and exit
```

```
$ cronmatch -h -9 '0 1 * * ? *' '2022/10/20 10:00'
'0 1 * * ? *' matches '2022/10/20 10:00' (offset: -9h)

$ cronmatch '0 10 * * ? *' 'Oct 10, 2022, 10:10'
'0 10 * * ? *' does not match 'Oct 10, 2022, 10:10'
```

cf. https://pkg.go.dev/github.com/araddon/dateparse#readme-extended-example

# cronviz CLI

CLI to visualize cron schedule.

inspired by [cronv](https://github.com/takumakanari/cronv), [aws-cronv](https://www.npmjs.com/package/aws-cronv).

## Installation

```
brew install winebarrel/cronplan/cronviz
```

## Usage

```
Usage: cronviz [OPTION] [FILE]
  -f string
    	from date (default current date)
  -h int
    	hour to add
  -p string
    	period (default "1d")
  -version
    	print version and exit
```

```
$ cat cron.txt
batch1  0 * * * ? *
batch2  30 */2 * * ? *
batch3  15,45 */3 * * ? *

$ cronviz cron.txt > output.html
$ open output.html
```

cf. https://raw.githack.com/winebarrel/cronplan/main/_example/timeline.html

# crongrep CLI

CLI to grep with cron expression.

## Installation

```
brew install winebarrel/cronplan/crongrep
```

## Usage

```
Usage: crongrep [OPTION] CRON_EXPR
  -version
    	print version and exit
```

```
$ cronplan -n 5 '10 12 */5 * ? *'
Fri, 06 Oct 2023 12:10:00
Wed, 11 Oct 2023 12:10:00
Mon, 16 Oct 2023 12:10:00
Sat, 21 Oct 2023 12:10:00
Thu, 26 Oct 2023 12:10:00

$ cronplan -n 5 '10 12 */5 * ? *' | crongrep '* * ? * WED-FRI *'
Fri, 06 Oct 2023 12:10:00
Wed, 11 Oct 2023 12:10:00
Thu, 26 Oct 2023 12:10:00
```

# cronskd CLI

CLI to show a schedule of cron expressions.

## Installation

```
brew install winebarrel/cronplan/cronskd
```

## Usage

```
Usage: cronskd [OPTION] [FILE]
  -e string
    	end date (default: end of day)
  -s string
    	start date (default: beginning of day)
  -version
    	print version and exit
```

```
$ cat exprs.txt
0 10 * * ? *
15 12 * * ? *
0 18 ? * MON-FRI *
0 8 1 * ? *
5 8-10 ? * MON-FRI *

$ cronskd -s '2024-11-11' exprs.txt
Mon, 11 Nov 2024 08:05:00    5 8-10 ? * MON-FRI *
Mon, 11 Nov 2024 09:05:00    5 8-10 ? * MON-FRI *
Mon, 11 Nov 2024 10:00:00    0 10 * * ? *
Mon, 11 Nov 2024 10:05:00    5 8-10 ? * MON-FRI *
Mon, 11 Nov 2024 12:15:00    15 12 * * ? *
Mon, 11 Nov 2024 18:00:00    0 18 ? * MON-FRI *

$ cronskd -s '2024/11/12 10:00' -e 'Nov 13, 2024, 12:00' exprs.txt
Tue, 12 Nov 2024 10:00:00	0 10 * * ? *
Tue, 12 Nov 2024 10:05:00	5 8-10 ? * MON-FRI *
Tue, 12 Nov 2024 12:15:00	15 12 * * ? *
Tue, 12 Nov 2024 18:00:00	0 18 ? * MON-FRI *
Wed, 13 Nov 2024 08:05:00	5 8-10 ? * MON-FRI *
Wed, 13 Nov 2024 09:05:00	5 8-10 ? * MON-FRI *
Wed, 13 Nov 2024 10:00:00	0 10 * * ? *
Wed, 13 Nov 2024 10:05:00	5 8-10 ? * MON-FRI *
```

cf. https://pkg.go.dev/github.com/araddon/dateparse#readme-extended-example

## Related Links

* https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-scheduled-rule-pattern.html#eb-cron-expressions
* https://github.com/winebarrel/terraform-provider-cronplan
* [Cron Trigger Tutorial](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html)

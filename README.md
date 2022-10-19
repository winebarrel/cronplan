# cronplan

Cron expression parser for Amazon EventBridge.

## Installation

```sh
go get github.com/winebarrel/cronplan@latest
```

## Usage

```go
package main

import (
	"fmt"
	"time"

	"github.com/winebarrel/cronplan"
)

func main() {
	cron, err := cronplan.Parse("0 10 * * ? *")

	if err != nil {
		panic(err)
	}

	fmt.Println(cron.Minute.Exps[0].Number) //=> 0
	fmt.Println(cron.Hour.Exps[0].Number)   //=> 10
	fmt.Println(cron.String())              //=> "0 10 * * ? *"

	fmt.Println(cron.Match(time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC)))
	//=> false
	fmt.Println(cron.Match(time.Date(2022, 11, 3, 10, 0, 0, 0, time.UTC)))
	//=> true

	fmt.Println(cron.Next(time.Date(2022, 11, 3, 10, 0, 0, 0, time.UTC)))
	//=> 2022-11-03 10:00:00 +0000 UTC
	fmt.Println(cron.Next(time.Date(2022, 11, 3, 11, 0, 0, 0, time.UTC)))
	//=> 2022-11-04 10:00:00 +0000 UTC
	fmt.Println(cron.NextN(time.Date(2022, 11, 3, 10, 0, 0, 0, time.UTC), 3))
	//=> [2022-11-03 10:00:00 +0000 UTC 2022-11-04 10:00:00 +0000 UTC 2022-11-05 10:00:00 +0000 UTC]
}
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
$ cronplan "*/10 10 ? * MON-FRI *"
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

$ cronplan -h -9 "*/10 10 ? * MON-FRI *"
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
  -version
    	print version and exit
```

```
$ cronmatch -h -9 "0 1 * * ? *" "2022/10/20 10:00"
'2022/10/20 10:00' (offset: -9h) does not match '0 1 * * ? *'

$ cronmatch  "0 10 * * ? *" "Oct 10, 2022, 10:10"
'Oct 10, 2022, 10:10' (offset: 0h) does not match '0 10 * * ? *'
```

cf. https://github.com/araddon/dateparse

# Related Links

* [Schedule Expressions for Rules - Amazon CloudWatch Events](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html)
* [cronplan.in](http://cronplan.in/)

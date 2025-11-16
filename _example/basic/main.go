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
		//=> 2022-11-06 10:00:00 +0000 UTC
		break
	}
}

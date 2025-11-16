package main

import (
	"fmt"
	"time"

	"github.com/winebarrel/cronplan/v2"
)

func cron(expr string, from time.Time, proc func()) (<-chan struct{}, error) {
	cron, err := cronplan.Parse(expr)

	if err != nil {
		return nil, err
	}

	waiter := make(chan struct{})

	go func() {
		for next := range cron.Iter(from).Seq() {
			<-time.After(time.Until(next))
			go proc()
		}

		close(waiter)
	}()

	return waiter, nil
}

func main() {
	startTime := time.Now().Add(1 * time.Minute)

	waiter, err := cron("* * * * ? *", startTime, func() {
		fmt.Println(time.Now())
	})

	if err != nil {
		panic(err)
	}

	<-waiter
}

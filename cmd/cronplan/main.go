package main

import (
	"fmt"
	"log"
	"time"

	"github.com/winebarrel/cronplan"
)

func init() {
	log.SetFlags(0)
}

func main() {
	flags := parseFlags()
	cron, err := cronplan.Parse(flags.expr)

	if err != nil {
		log.Fatal(err)
	}

	triggers := cron.NextN(time.Now(), flags.n)

	for _, t := range triggers {
		fmt.Println(t.Add(time.Duration(flags.h) * time.Hour).Format("Mon, 02 Jan 2006 15:04:05"))
	}
}

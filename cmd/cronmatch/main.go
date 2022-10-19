package main

import (
	"fmt"
	"log"
	"os"

	"github.com/araddon/dateparse"
	"github.com/winebarrel/cronplan"
)

func init() {
	log.SetFlags(0)
}

func main() {
	flags := parseFlags()

	cron, err := cronplan.Parse(flags.expr)

	if err != nil {
		log.Fatalf("failed to cron expr: %s", err)
	}

	t, err := dateparse.ParseAny(flags.t)

	if err != nil {
		log.Fatalf("failed to parse date: %s", err)
	}

	t = t.AddDate(0, 0, flags.h)

	m := cron.Match(t)

	if m {
		fmt.Printf("'%s' (offset: %dh) matches '%s'\n", flags.t, flags.h, flags.expr)
	} else {
		fmt.Printf("'%s' (offset: %dh) does not match '%s'\n", flags.t, flags.h, flags.expr)
		os.Exit(1)
	}
}

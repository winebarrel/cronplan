package main

import (
	"log"
	"os"
	"time"

	"github.com/araddon/dateparse"
	"github.com/fatih/color"
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

	t = t.Add(time.Duration(flags.h) * time.Hour)
	m := cron.Match(t)

	if m {
		color.Green("'%s' matches '%s' (offset: %dh)\n", flags.expr, flags.t, flags.h)
	} else {
		color.Red("'%s' does not match '%s' (offset: %dh)\n", flags.expr, flags.t, flags.h)
		os.Exit(1)
	}
}

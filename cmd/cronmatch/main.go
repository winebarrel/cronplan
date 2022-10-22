package main

import (
	"fmt"
	"log"
	"os"
	"strings"
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
		var buf strings.Builder
		fmt.Fprintf(&buf, "'%s' matches '%s'", flags.expr, flags.t)

		if flags.h != 0 {
			fmt.Fprintf(&buf, " (offset: %dh)", flags.h)
		}

		color.Green(buf.String())
	} else {
		var buf strings.Builder
		fmt.Fprintf(&buf, "'%s' does not match '%s'", flags.expr, flags.t)

		if flags.h != 0 {
			fmt.Fprintf(&buf, " (offset: %dh)", flags.h)
		}

		color.Red(buf.String())
		os.Exit(1)
	}
}

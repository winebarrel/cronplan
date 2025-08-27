package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/araddon/dateparse"
	"github.com/winebarrel/cronplan/v2"
)

func init() {
	log.SetFlags(0)
}

func main() {
	flags := parseFlags()

	var scanner *bufio.Scanner

	if flags.file == "-" {
		scanner = bufio.NewScanner(os.Stdin)
	} else {
		file, err := os.Open(flags.file)

		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		scanner = bufio.NewScanner(file)
	}

	type exprNext struct {
		expr string
		next time.Time
	}

	schedule := []exprNext{}

	for scanner.Scan() {
		expr := scanner.Text()
		expr = strings.TrimSpace(expr)
		expr = strings.TrimPrefix(expr, "cron(")
		expr = strings.TrimSuffix(expr, ")")

		cron, err := cronplan.Parse(expr)

		if err != nil {
			log.Fatal(err)
		}

		var start, end time.Time

		if flags.start == "" {
			now := time.Now()
			start = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		} else {
			start, err = dateparse.ParseAny(flags.start)

			if err != nil {
				log.Fatal(err)
			}
		}

		if flags.end == "" {
			end = time.Date(start.Year(), start.Month(), start.Day(), 23, 59, 50, 0, start.Location())
		} else {
			end, err = dateparse.ParseAny(flags.end)

			if err != nil {
				log.Fatal(err)
			}
		}

		nexts := cron.Between(start, end)

		for _, n := range nexts {
			schedule = append(schedule, exprNext{expr: expr, next: n})
		}
	}

	sort.Slice(schedule, func(i, j int) bool {
		return schedule[i].next.Before(schedule[j].next)
	})

	for _, ln := range schedule {
		fmt.Printf("%s\t%s\n", ln.next.Format("Mon, 02 Jan 2006 15:04:05"), ln.expr)
	}
}

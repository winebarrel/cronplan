package main

import (
	"bufio"
	_ "embed"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"
	"time"

	"github.com/araddon/dateparse"
	"github.com/k1LoW/duration"
	"github.com/winebarrel/cronplan"
)

//go:embed timeline.html.tmpl
var timelineTmpl string

func init() {
	log.SetFlags(0)

}

type Row struct {
	Expr  string
	Times []time.Time
}

func main() {
	flags := parseFlags()

	var from time.Time
	var err error

	if flags.from == "" {
		from = time.Now()
		from = time.Date(from.Year(), from.Month(), from.Day(), 0, 0, 0, 0, from.Location())
	} else {
		from, err = dateparse.ParseAny(flags.from)

		if err != nil {
			log.Fatalf("failed to parse from date: %s", err)
		}
	}

	period, err := duration.Parse(flags.period)

	if err != nil {
		log.Fatalf("failed to parse period: %s", err)
	}

	to := from.Add(period)
	var file io.ReadCloser

	if flags.input == "" {
		file = os.Stdin
	} else {
		file, err = os.OpenFile(flags.input, os.O_RDONLY, 0)

		if err != nil {
			log.Fatalf("failed to open %s: %s", flags.input, err)
		}
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	r := regexp.MustCompile(`\s+`)

	schedule := map[string]*Row{}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		fields := r.Split(line, 2)

		if len(fields) < 2 {
			log.Fatalf("too few fields: %s", line)
		}

		name := fields[0]
		expr := fields[1]
		cron, err := cronplan.Parse(expr)

		if err != nil {
			log.Fatalf("failed to parse cron expr: %s/%s: %s", name, expr, err)
		}

		schedule[name] = &Row{
			Expr:  expr,
			Times: cron.Between(from, to),
		}
	}

	if len(schedule) == 0 {
		log.Fatal("input is empty")
	}

	t := template.Must(template.New("timeline.html.tmpl").Funcs(map[string]interface{}{
		"monthn": func(m time.Month) int {
			return int(m)
		},
	}).Parse(timelineTmpl))

	err = t.Execute(os.Stdout, schedule)

	if err != nil {
		log.Fatalf("failed to execute template: %s", err)
	}
}

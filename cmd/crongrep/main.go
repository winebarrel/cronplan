package main

import (
	"bufio"
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
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		t, err := dateparse.ParseAny(line)

		if err != nil {
			log.Fatal(err)
		}

		if cron.Match(t) {
			fmt.Println(line)
		}
	}
}

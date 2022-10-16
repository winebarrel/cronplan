package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	version string
)

type flags struct {
	n    int
	h    int
	expr string
}

func init() {
	cmdLine := flag.NewFlagSet(filepath.Base(os.Args[0]), flag.ExitOnError)

	cmdLine.Usage = func() {
		fmt.Fprintf(cmdLine.Output(), "Usage: %s [OPTION] CRON_EXPR\n", cmdLine.Name())
		cmdLine.PrintDefaults()
	}

	flag.CommandLine = cmdLine
}

func parseFlags() *flags {
	flags := &flags{}
	flag.IntVar(&flags.h, "h", 0, "hour to add")
	flag.IntVar(&flags.n, "n", 10, "number of next triggers")
	showVersion := flag.Bool("version", false, "print version and exit")
	flag.Parse()

	if *showVersion {
		printVersionAndExit()
	}

	args := flag.Args()

	if len(args) < 1 {
		printUsageAndExit()
	} else if len(args) > 1 {
		log.Fatal("too many arguments")
	}

	flags.expr = strings.TrimSpace(args[0])

	if flags.expr == "" {
		printUsageAndExit()
	}

	if flags.n < 1 {
		log.Fatal("'-n' must be >= 1")
	}

	return flags
}

func printVersionAndExit() {
	v := version

	if v == "" {
		v = "<nil>"
	}

	fmt.Fprintln(flag.CommandLine.Output(), v)
	os.Exit(0)
}

func printUsageAndExit() {
	flag.CommandLine.Usage()
	os.Exit(0)
}

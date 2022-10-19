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
	h    int
	expr string
	t    string
}

func init() {
	cmdLine := flag.NewFlagSet(filepath.Base(os.Args[0]), flag.ExitOnError)

	cmdLine.Usage = func() {
		fmt.Fprintf(cmdLine.Output(), "Usage: %s [OPTION] CRON_EXPR DATE\n", cmdLine.Name())
		cmdLine.PrintDefaults()
	}

	flag.CommandLine = cmdLine
}

func parseFlags() *flags {
	flags := &flags{}
	flag.IntVar(&flags.h, "h", 0, "hour to add")
	showVersion := flag.Bool("version", false, "print version and exit")
	flag.Parse()

	if *showVersion {
		printVersionAndExit()
	}

	args := flag.Args()

	if len(args) == 0 {
		printUsageAndExit()
	} else if len(args) < 2 {
		log.Fatal("too few arguments")
	} else if len(args) > 2 {
		log.Fatal("too many arguments")
	}

	flags.expr = strings.TrimSpace(args[0])
	flags.t = strings.TrimSpace(args[1])

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

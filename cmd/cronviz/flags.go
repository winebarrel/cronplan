package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	version string
)

type flags struct {
	from   string
	period string
	h      int
	input  string
}

func init() {
	cmdLine := flag.NewFlagSet(filepath.Base(os.Args[0]), flag.ExitOnError)

	cmdLine.Usage = func() {
		fmt.Fprintf(cmdLine.Output(), "Usage: %s [OPTION] [FILE]\n", cmdLine.Name())
		cmdLine.PrintDefaults()
	}

	flag.CommandLine = cmdLine
}

func parseFlags() *flags {
	flags := &flags{}

	flag.StringVar(&flags.from, "f", "", "from date (default current date)")
	flag.StringVar(&flags.period, "p", "1d", "period")
	flag.IntVar(&flags.h, "h", 0, "hour to add")
	showVersion := flag.Bool("version", false, "print version and exit")
	flag.Parse()

	if *showVersion {
		printVersionAndExit()
	}

	args := flag.Args()

	if len(args) > 1 {
		log.Fatal("too many arguments")
	} else if len(args) == 1 {
		flags.input = args[0]
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

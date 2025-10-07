module github.com/winebarrel/cronplan/v2/cmd/cronmatch

go 1.21

toolchain go1.25.2

replace github.com/winebarrel/cronplan/v2 => ../..

require (
	github.com/araddon/dateparse v0.0.0-20210429162001-6b43995a97de
	github.com/fatih/color v1.18.0
	github.com/mattn/go-isatty v0.0.20
	github.com/winebarrel/cronplan/v2 v2.0.0-00010101000000-000000000000
)

require (
	github.com/alecthomas/participle/v2 v2.1.4 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	golang.org/x/sys v0.25.0 // indirect
)

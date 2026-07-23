module github.com/winebarrel/cronplan/v2/cmd/cronmatch

go 1.25.0

toolchain go1.26.5

replace github.com/winebarrel/cronplan/v2 => ../..

require (
	github.com/araddon/dateparse v0.0.0-20210429162001-6b43995a97de
	github.com/fatih/color v1.19.0
	github.com/mattn/go-isatty v0.0.24
	github.com/winebarrel/cronplan/v2 v2.0.0-00010101000000-000000000000
)

require (
	github.com/alecthomas/participle/v2 v2.1.4 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	golang.org/x/sys v0.42.0 // indirect
)

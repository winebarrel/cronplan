module github.com/winebarrel/cronplan/v2/cmd/cronviz

go 1.23

toolchain go1.25.5

replace github.com/winebarrel/cronplan/v2 => ../..

require (
	github.com/araddon/dateparse v0.0.0-20210429162001-6b43995a97de
	github.com/k1LoW/duration v1.2.0
	github.com/winebarrel/cronplan/v2 v2.0.0-00010101000000-000000000000
)

require github.com/alecthomas/participle/v2 v2.1.4 // indirect

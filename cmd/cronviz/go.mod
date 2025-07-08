module github.com/winebarrel/cronplan/cmd/cronviz

go 1.21

toolchain go1.24.5

replace github.com/winebarrel/cronplan => ../..

require (
	github.com/araddon/dateparse v0.0.0-20210429162001-6b43995a97de
	github.com/k1LoW/duration v1.2.0
	github.com/winebarrel/cronplan v0.0.0-00010101000000-000000000000
)

require github.com/alecthomas/participle/v2 v2.1.4 // indirect

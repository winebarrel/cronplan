module github.com/winebarrel/cronplan/cmd/cronskd

go 1.21

toolchain go1.23.3

replace github.com/winebarrel/cronplan => ../..

require (
	github.com/araddon/dateparse v0.0.0-20210429162001-6b43995a97de
	github.com/winebarrel/cronplan v0.0.0-00010101000000-000000000000
)

require github.com/alecthomas/participle/v2 v2.1.1 // indirect

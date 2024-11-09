module github.com/winebarrel/cronplan/test

go 1.21

toolchain go1.21.1

replace github.com/winebarrel/cronplan => ../

require (
	github.com/stretchr/testify v1.9.0
	github.com/winebarrel/cronplan v0.0.0-00010101000000-000000000000
)

require (
	github.com/alecthomas/participle/v2 v2.1.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

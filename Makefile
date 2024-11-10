.PHONY: all
all: vet test build

.PHONY: build
build:
	cd ./cmd/cronplan && go build -o ../../cronplan
	cd ./cmd/cronmatch && go build -o ../../cronmatch
	cd ./cmd/cronviz && go build -o ../../cronviz
	cd ./cmd/crongrep && go build -o ../../crongrep
	cd ./cmd/cronskd && go build -o ../../cronskd

.PHONY: vet
vet:
	go vet -composites=false -structtag=false ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	cd test && go test -v ./...

.PHONY: clean
clean:
	rm -f cronplan cronplan.exe
	rm -f cronmatch cronmatch.exe
	rm -f cronviz cronviz.exe
	rm -f crongrep crongrep.exe
	rm -f cronskd cronskd.exe

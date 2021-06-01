ProjectName := github.com/SomethingBot/gateway
GitCommit   := $(shell git rev-parse HEAD)
GitTag      := $(shell git tag --points-at HEAD)

build:
	go build -ldflags "-X main.GitCommit=$(GitCommit) -X main.GitTag=$(GitTag) -X main.Mode=Dev"

test:
	go vet
	go test

test-integration:
	go test -tags=integration

lint:
	go fmt -v

run: build
	./gateway
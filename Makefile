ProjectName := github.com/SomethingBot/gateway
GitCommit   := $(shell git rev-parse HEAD)
GitTag      := $(shell git tag --points-at HEAD)

build:
	go build -ldflags "-X main.GitCommit=$(GitCommit) -X main.GitTag=$(GitTag) -X main.Mode=Dev"

test:
	go vet ./...
	go test ./...

spawn-etcd:
	podman run -p 2379:2379 --rm --name gateway-etcd-dev quay.io/coreos/etcd:v3.4.16 /usr/local/bin/etcd -listen-client-urls http://0.0.0.0:2379 -advertise-client-urls http://0.0.0.0:2379

spawn-redis:
	podman run -p 6379:6379 --rm --name alias-redis-dev docker.io/library/redis

test-integration:
	go test -tags=integration

lint:
	go fmt ./...

run: build
	GATEWAY_ETCD_ENDPOINTS=localhost:2379 GATEWAY_ADDRESS=0.0.0.0:8987 ./gateway

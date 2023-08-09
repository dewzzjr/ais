OS=Linux
ARCH=x86_64
MOCKGEN_VERSION=1.6.0

dep:
ifeq (, $(shell which mockgen))
	go install github.com/golang/mock/mockgen@v${MOCKGEN_VERSION}
endif
	@echo "Ready"

dep-force:
	go install github.com/golang/mock/mockgen@v${MOCKGEN_VERSION}
	$(MAKE) download-proto

test:
	go vet ./...
	go test -timeout 30s -cover -coverpkg=all -v ./...

mod:
	go mod tidy
	go mod vendor

build:
	GOOS=linux go build -o bin/ais .

generate:
	go generate ./...

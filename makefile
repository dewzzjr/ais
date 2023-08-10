OS?=linux
ARCH?=x86_64
MOCKGEN_VERSION?=1.6.0
MIGRATE_VERSION?=4.16.2

ifneq ($(wildcard .env),)
include .env
export $(shell sed 's/=.*//' .env)
endif

dep:
ifeq (, $(shell which mockgen))
	$(MAKE) download-mockgen
endif
ifeq (, $(shell which migrate))
	$(MAKE) download-migrate
endif
	@echo "Ready"

DB_URL?=mysql://${DB_USER}:${DB_PASSWORD}@tcp(localhost:${DB_PORT})/${DB_NAME}?charset=utf8mb4&parseTime=True

download-migrate:
	go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v${MIGRATE_VERSION}
	migrate --version
	chmod +x $(shell which migrate)

download-mockgen:
	go install github.com/golang/mock/mockgen@v${MOCKGEN_VERSION}
	mockgen --version

test:
	go vet ./...
	go test -timeout 30s -cover -coverpkg=all -v ./...

mod:
	go mod tidy
	go mod vendor

env:
ifneq ($(wildcard .env),)
	cp .env .env.backup
endif
	cp .env.example .env

build:
	go build -o bin/server cmd/server/main.go

remove-docker:
	docker-compose down

build-docker: build
	docker-compose up --build

mysql-docker-up:
	docker start mysql_articles_db

migrate: mysql-docker-up
	migrate -database='${DB_URL}' -source=file://./migrations up
	
reset-migrate: mysql-docker-up
	migrate -database='${DB_URL}' -source=file://./migrations down -all

reset-docker: remove-docker build-docker 

generate:
	go generate ./...

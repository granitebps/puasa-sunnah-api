APP_NAME = api
BUILD_DIR = $(PWD)/build

swag:
	swag init

air:
	air

build-app:
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) .

build-backup:
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/backup cmd/backup.go

lint:
	golangci-lint run ./...

security:
	gosec ./...

wire:
	wire

## test: runs all tests
test:
	go test -v ./...

## cover: opens coverage in browser
cover:
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

## coverage: displays test coverage
coverage:
	go test -cover ./...

vet:
	go vet .
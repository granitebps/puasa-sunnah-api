APP_NAME = api
BUILD_DIR = $(PWD)/build

air:
	air

build-app:
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) .

build-backup:
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/backup cmd/backup.go

build: build-app build-backup

build-app-linux: 
	GOOS=linux GOARCH=amd64 $(MAKE) build-app

build-backup-linux: 
	GOOS=linux GOARCH=amd64 $(MAKE) build-backup

build-linux: build-app-linux build-backup-linux

# install development dependencies ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

check:
	@ gotestsum --version
	@ tparse --version
	@ air -v

deps_i:
	@ go install gotest.tools/gotestsum@latest
	@ go install github.com/mfridman/tparse@latest
	@ go install github.com/air-verse/air@latest
	@ go install github.com/swaggo/swag/cmd/swag@latest
	@ go install github.com/securego/gosec/v2/cmd/gosec@latest
# curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.61.0

# Pre-push command ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

pre-push: swag lint security vet wire test

swag:
	swag init

lint:
	golangci-lint run ./...

security:
	gosec ./...

vet:
	go vet .

wire:
	wire

# Testing command ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

TESTS_ARGS := --format testname --jsonfile gotestsum.json.out
TESTS_ARGS += --max-fails 2
TESTS_ARGS += -- ./...
TESTS_ARGS += -test.parallel 2
TESTS_ARGS += -test.count    1
TESTS_ARGS += -test.failfast
TESTS_ARGS += -test.coverprofile   coverage.out
TESTS_ARGS += -test.timeout        5s
TESTS_ARGS += -race

test: run-test $(TPARSE) ## Run Tests & parse details
	@ cat gotestsum.json.out | tparse -all -notests

run-test: $(GOTESTSUM)
	@ gotestsum $(TESTS_ARGS) -short

## coverage: displays test coverage
coverage:
	go test -cover ./...

## cover: opens coverage in browser
cover:
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

# clean ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

clear: clean-artifacts

clean-artifacts: ## Removes Artifacts
	@ printf "Cleanning artifacts... "
	@ rm -f *.out
	@ rm -f *.html
	@ echo "done."

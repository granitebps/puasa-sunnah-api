APP_NAME = api
BUILD_DIR = $(PWD)/build

air:
	air

build-app:
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) .

build-backup:
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/backup cmd/backup.go

# install development dependencies ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

check:
	@ gotestsum --version
	@ tparse --version
	@ air -v

deps_i:
	@ go install gotest.tools/gotestsum@latest
	@ go install github.com/mfridman/tparse@latest
	@ go install github.com/cosmtrek/air@latest

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

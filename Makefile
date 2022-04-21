OS_NAME := $(shell uname -s | tr A-Z a-z)
ifeq ($(OS_NAME), darwin)
DOC = ./bin/swag-mac init
else
DOC = ./bin/swag-linux init
endif

start:
	go run main.go

build:
	go build main.go

doc:
	$(DOC)

air:
	./bin/air
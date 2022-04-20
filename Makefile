start:
	go run main.go

build:
	go build main.go

doc:
	./bin/swag-mac init

doc-deploy:
	./bin/swag-linux init

air:
	./bin/air
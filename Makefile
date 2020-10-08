.PHONY: mocks
mocks:
	cd ./repository/mocks/; go generate;
	cd ./service/mocks/; go generate;

.PHONY: build
build:
	go build -o cmd/api/simple-web-server cmd/api/main.go

.PHONY: run
run:
	cd cmd/api; ./rundev.sh

tests:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := run

.PHONY: mocks
mocks:
	cd ./store/mocks/; go generate;
	cd ./service/mocks/; go generate;

.PHONY: build
build:
	go build -o cmd/api/simple-web-server cmd/api/main.go

.PHONY: run
run:
	cd cmd/api; ./rundev.sh

.DEFAULT_GOAL := run

mocks:
	cd ./store/mocks/; go generate;
	cd ./service/mocks/; go generate;

build:
	go build -o cmd/api/simple-web-server cmd/api/main.go

run:
	cd cmd/api; ./rundev.sh


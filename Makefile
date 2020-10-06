build:
	go build -o cmd/api/simple-web-server cmd/api/main.go

run:
	cd cmd/api; ./rundev.sh

test:
	go test -v ./...
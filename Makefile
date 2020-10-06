build:
	go build -o cmd/simple-web-server cmd/main.go

run:
	cd cmd; ./rundev.sh

test:
	go test -v ./...
FROM golang:1.15

COPY . /go/src/app

WORKDIR /go/src/app/cmd/api

RUN go build -o rest-api main.go

EXPOSE 8080

CMD ["bash","-c", "source env.docker.sh && rest-api"]

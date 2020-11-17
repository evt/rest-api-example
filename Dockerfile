# Build the project
FROM golang:1.14 as builder

WORKDIR /usr/app
ADD . .

RUN make build

# Create production image for application with needed files
FROM golang:1.14

EXPOSE 8080
RUN apt-get update && apt-get install ca-certificates

COPY --from=builder /usr/app/cmd/api/rest-api .
COPY --from=builder /usr/app/cmd/api/env.docker.sh .
COPY --from=builder /usr/app/store/pg/migrations ./pg/migrations
COPY --from=builder /usr/app/store/mysql/migrations ./mysql/migrations

CMD ["bash","-c", "cat ./env.docker.sh; source ./env.docker.sh; ./rest-api"]

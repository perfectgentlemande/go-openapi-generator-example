FROM golang:1.22 AS builder

ADD . /app
WORKDIR /app
RUN GOOS=linux GOARCH=amd64 go build -o go-mongodb-app ./cmd/go-openapi-generator-example

FROM alpine:3.19 AS app
WORKDIR /app
COPY --from=builder /app/go-mongodb-app /app
COPY --from=builder /app/cmd/go-openapi-generator-example/.env /app
CMD ["/app/go-mongodb-app"]
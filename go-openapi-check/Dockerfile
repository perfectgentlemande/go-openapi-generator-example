FROM golang:1.22 AS build

ADD . /app
WORKDIR /app
RUN GOOS=linux GOARCH=amd64 go build -o go-mongodb-app ./cmd/go-mongodb-crud-api-example

FROM alpine:3.19 AS app
WORKDIR /app
COPY --from=builder /app/go-mongodb-app /app
CMD ["/app/go-mongodb-app"]
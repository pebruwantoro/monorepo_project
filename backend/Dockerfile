FROM golang:alpine AS builder

# Build the binary
WORKDIR /src

RUN mkdir app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . ./app
COPY ./resources/config.json /src/app/resources/config.json

WORKDIR /src/app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=jsoniter,netgo,nomsgpack -ldflags='-s -w -extldflags "-static"' -o executable

# Serve the app
FROM alpine

WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /src/app/executable /app/
COPY --from=builder /src/app/resources/config.json /app/resources/
COPY --from=builder /src/app/migration /app/migration

EXPOSE 9090

ENTRYPOINT [ "/app/executable", "server", "rest" ]
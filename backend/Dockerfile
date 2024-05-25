FROM golang:alpine AS builder

# Token for label docker-build-private-repo-depedency
ARG USERNAME=efishery-eng
ARG TOKEN=ATBBU5ausfxGPmy9gBRy6Luw8dNCCDD3726A

RUN apk add git
RUN git config --global url."https://${USERNAME}:${TOKEN}@bitbucket.org/".insteadOf "https://bitbucket.org/"

# Build the binary
WORKDIR /src

RUN mkdir app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . ./app
COPY config.json /src/app/resources/config.json

WORKDIR /src/app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=jsoniter,netgo,nomsgpack -ldflags='-s -w -extldflags "-static"' -o executable

# Serve the app
FROM alpine

WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /src/app/executable /app/
COPY --from=builder /src/app/resources/config.json /app/resources/
COPY --from=builder /src/app/migration /app/migration

EXPOSE 9000
EXPOSE 9001

ENTRYPOINT [ "/app/executable", "server", "rest" ]
FROM golang:1.18.5-alpine3.16

WORKDIR /go/src/app

ENV GO111MODULE=on

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN env GOOS=linux GOARCH=amd64 go build -o /go/bin/app

ENTRYPOINT ["/go/bin/app"]
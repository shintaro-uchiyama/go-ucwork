FROM golang:1.14.6

EXPOSE 8080
WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go get github.com/cespare/reflex
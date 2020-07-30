FROM golang:1.14.6

WORKDIR /go/src/app

RUN go get github.com/cespare/reflex
CMD reflex -r '(\.go$|go\.mod)' -s go build -- -o hello cmd/hello/main.go
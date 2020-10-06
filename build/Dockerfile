FROM golang:1.15.2-alpine3.12 as builder
WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY server.go .
COPY graph/ ./graph/
COPY pkg/ ./pkg/
RUN go build -o /app server.go

FROM alpine:3.12
# Define GOTRACEBACK to mark this container as using the Go language runtime
# for `skaffold debug` (https://skaffold.dev/docs/workflows/debug/).
ENV GOTRACEBACK=single
CMD ["./app"]
COPY --from=builder /app .
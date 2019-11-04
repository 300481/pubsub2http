# Building stage
FROM golang:1.13.4-buster AS builder

WORKDIR /go/src/github.com/300481/pubsub2http
COPY . .
WORKDIR /go/src/github.com/300481/pubsub2http/cmd/pubsub2http
RUN go get -d -v && \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o /pubsub2http

# Run stage
FROM gcr.io/distroless/static:latest

COPY --from=builder /pubsub2http /pubsub2http

CMD [ "server" ]

ENTRYPOINT [ "/pubsub2http" ]
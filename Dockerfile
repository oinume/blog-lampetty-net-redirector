# build stage
FROM golang:1.15-buster AS builder
WORKDIR /go/src/github.com/oinume/blog-lampetty-net-redirector
COPY . .
RUN make build

# final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata
COPY --from=builder /go/src/github.com/oinume/blog-lampetty-net-redirector/server /bin/server
ENV PORT=${PORT}
ENTRYPOINT [ "/bin/server" ]

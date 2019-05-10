# build stage
FROM golang:1.12-stretch AS builder
WORKDIR /go/src/github.com/oinume/blog-lampetty-net-redirector
COPY . .
ENV GO111MODULE=on
RUN make build

# final stage
#FROM gcr.io/distroless/base
FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata
COPY --from=builder /go/src/github.com/oinume/blog-lampetty-net-redirector/server /bin/server
ENV PORT=${PORT}
#ENTRYPOINT [ "/app" ]
CMD [ "/bin/server" ]

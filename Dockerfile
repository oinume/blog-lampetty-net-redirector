# build stage
FROM golang:1.17-bullseye AS build
WORKDIR /go/src/github.com/oinume/blog-lampetty-net-redirector
COPY . .
RUN make build

# final stage
FROM gcr.io/distroless/base-debian11
COPY --from=build /go/src/github.com/oinume/blog-lampetty-net-redirector/server /bin/server
ENV PORT=${PORT}
ENTRYPOINT [ "/bin/server" ]

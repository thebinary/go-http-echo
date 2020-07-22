# build stage
FROM golang:1.14 AS builder
LABEL maintainer="binary4bytes@gmail.com"

ENV CGO_ENABLED 0
ADD ./ /go/src/github.com/thebinary/go-http-echo
WORKDIR /go/src/github.com/thebinary/go-http-echo
RUN set -ex && \
    CGO_ENABLED=0 GOOS=linux go install -a -ldflags '-extldflags "-static"' .

# final stage
FROM busybox
COPY --from=builder /go/bin/go-http-echo /usr/bin/go-http-echo
ENV PORT 8080
EXPOSE 8080
ENTRYPOINT ["/usr/bin/go-http-echo"]
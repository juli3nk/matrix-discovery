FROM golang:1.14-alpine3.12 AS builder

COPY *.go /go/src/github.com/juli3nk/matrix-discovery/

WORKDIR /go/src/github.com/juli3nk/matrix-discovery

RUN apk --update add \
		ca-certificates \
		gcc \
		git \
		musl-dev

RUN echo 'nobody:x:65534:65534:nobody:/:' > /tmp/passwd \
	&& echo 'nobody:x:65534:' > /tmp/group

RUN go get \
	&& go build -ldflags "-linkmode external -extldflags -static -s -w" -o /tmp/matrix-discovery


FROM scratch

COPY --from=builder /tmp/group /tmp/passwd /etc/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /tmp/matrix-discovery /matrix-discovery

USER nobody:nobody

EXPOSE 8080

ENTRYPOINT ["/matrix-discovery"]

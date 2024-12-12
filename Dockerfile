FROM alpine

LABEL maintainer="antoniorr@correo.ugr.es"
LABEL version="0.0.1"

WORKDIR /app

RUN adduser -D test 

RUN apk update && apk add go gcc bash musl-dev openssl-dev ca-certificates && update-ca-certificates
RUN apk add just

USER test

WORKDIR /app/test

RUN mkdir -p /app/test/.cache/go-build /app/test/go/pkg /app/test/go/bin
ENV GOCACHE=/app/test/.cache/go-build
ENV GOPATH=/app/test/go
ENV GOMODCACHE=/app/test/go/pkg/mod

ENTRYPOINT ["just", "test"]
FROM golang:stretch

LABEL maintainer="antoniorr@correo.ugr.es"
LABEL version="0.0.1"

WORKDIR /app

RUN curl -fsSL https://just.systems/install.sh | bash -s -- --to /usr/local/bin

RUN adduser test 
USER test

WORKDIR /app/test
RUN mkdir -p /app/test/.cache/go-build /app/test/go/pkg /app/test/go/bin

ENV GOCACHE=/app/test/.cache/go-build
ENV GOPATH=/app/test/go
ENV GOMODCACHE=/app/test/go/pkg/mod
ENV PATH=/usr/local/go/bin:$PATH

ENTRYPOINT ["just", "test"]
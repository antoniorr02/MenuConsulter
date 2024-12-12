FROM golang:alpine

LABEL mantainer="antoniorr@correo.ugr.es"
LABEL version="0.0.1"

WORKDIR /app

RUN adduser -D test

RUN apk add just
USER test

WORKDIR /app/test

RUN mkdir -p /app/test/.cache/go-build
ENV GOCACHE=/app/test/.cache/go-build

ENTRYPOINT ["just", "test"]
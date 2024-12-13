FROM golang:alpine

LABEL mantainer="antoniorr@correo.ugr.es" version="0.0.4"

RUN apk add just

RUN adduser -D -h /home/test test
USER test

RUN mkdir -p /home/test/.cache/go-build && chmod -R 747 /home/test/.cache/go-build
ENV GOCACHE=/home/test/.cache/go-build

WORKDIR /app/test

ENTRYPOINT ["just", "test"]
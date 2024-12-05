FROM golang:1.23.3-alpine3.20

LABEL mantainer="antoniorr@correo.ugr.es" version="5.0.0"

WORKDIR /app

RUN adduser -D -u 1001 test && chown test /app

USER test

COPY go.mod go.sum justfile ./

RUN go mod download

WORKDIR /app/test

ENTRYPOINT ["just", "test"]


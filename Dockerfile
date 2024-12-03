FROM golang:1.23.3-alpine3.20

LABEL mantainer="antoniorr@correo.ugr.es" version="5.0.0"

WORKDIR /app

RUN adduser -D test && chown test /app

USER test

COPY go.mod justfile ./

RUN go mod download

WORKDIR /app/test

ENTRYPOINT ["just", "test"]


FROM golang:1.23.4-alpine

LABEL mantainer="antoniorr@correo.ugr.es"

WORKDIR /app

RUN adduser -D -u 1001 test && chown test /app

RUN apk add just
USER test

COPY go.mod go.sum /app/
RUN go mod download

WORKDIR /app/test

ENTRYPOINT ["just", "test"]
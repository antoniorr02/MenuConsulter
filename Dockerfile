FROM golang:1.23.3-alpine3.20

LABEL mantainer="antoniorr@correo.ugr.es" version="5.0.0"

WORKDIR /app

RUN adduser -D -u 1001 test && chown test /app

USER test

COPY . .

RUN go mod download

USER root
RUN apk add just
USER test

WORKDIR /app/test

ENTRYPOINT ["just", "test"]
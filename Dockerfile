FROM golang:1.23.4-alpine

LABEL mantainer="antoniorr@correo.ugr.es"

WORKDIR /app

RUN adduser -D test && chown test /app

RUN apk add just
USER test

COPY go.mod go.sum /app/
RUN go mod download

RUN mkdir -p /app/.cache/go-build && chmod -R 757 /app/.cache
ENV GOCACHE=/app/.cache/go-build

WORKDIR /app/test

ENTRYPOINT ["just", "test"]
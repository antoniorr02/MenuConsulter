FROM golang:alpine

LABEL mantainer="antoniorr@correo.ugr.es"

WORKDIR /app

RUN adduser -D test && chown test /app

RUN apk add just
USER test

COPY go.mod /app/
RUN go mod download

ENV GOCACHE=/app/test/.cache/go-build

WORKDIR /app/test

ENTRYPOINT ["just", "test"]
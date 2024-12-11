FROM golang:alpine

LABEL mantainer="antoniorr@correo.ugr.es"

WORKDIR /app

RUN adduser -D test

RUN apk add just
USER test

WORKDIR /app/test

COPY go.mod go.sum /app/test/
RUN go mod download

RUN mkdir -p /app/test/.cache/go-build
ENV GOCACHE=/app/test/.cache/go-build

ENTRYPOINT ["just", "test"]
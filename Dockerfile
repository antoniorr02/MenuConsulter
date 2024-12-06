FROM golang:1.23.4-alpine3.20

LABEL mantainer="antoniorr@correo.ugr.es"

WORKDIR /app

RUN adduser -D -u 1001 test && chown test /app

USER test

WORKDIR /app/test
COPY go.mod go.sum justfile /app/test/

WORKDIR /app/test/internal
COPY internal/comedor.go internal/menu_extractor_test.go internal/menu_extractor.go internal/menu.go internal/plato.go /app/test/internal/ 

WORKDIR /app/test/data
COPY data/menu.html /app/test/data/

RUN go mod download

USER root
RUN apk add just
USER test

ENTRYPOINT ["just", "test"]
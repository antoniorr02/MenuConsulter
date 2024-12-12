FROM golang

LABEL maintainer="antoniorr@correo.ugr.es"
LABEL version="0.0.1"

WORKDIR /app

RUN apt-get update && apt-get install -y curl \
    && curl -fsSL https://just.systems/install.sh | bash -s -- --to /usr/local/bin \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

RUN useradd -m test
USER test

WORKDIR /app/test

RUN mkdir -p /app/test/.cache/go-build
ENV GOCACHE=/app/test/.cache/go-build

ENTRYPOINT ["just", "test"]

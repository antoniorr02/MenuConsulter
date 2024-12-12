FROM debian:stable-slim

LABEL maintainer="antoniorr@correo.ugr.es"
LABEL version="0.0.1"

WORKDIR /app

RUN apt-get update && \
    apt-get install -y --no-install-recommends wget ca-certificates && \ 
    rm -rf /var/lib/apt/lists/*

RUN wget -q https://go.dev/dl/go1.23.3.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.23.3.linux-amd64.tar.gz && \
    rm go1.23.3.linux-amd64.tar.gz

RUN apt-get update && apt-get install -y curl \
    && curl -fsSL https://just.systems/install.sh | bash -s -- --to /usr/local/bin \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

RUN useradd -m test
USER test

WORKDIR /app/test
RUN mkdir -p /app/test/.cache/go-build /app/test/go/pkg /app/test/go/bin

ENV GOCACHE=/app/test/.cache/go-build
ENV GOPATH=/app/test/go
ENV GOMODCACHE=/app/test/go/pkg/mod
ENV PATH=/usr/local/go/bin:$PATH

ENTRYPOINT ["just", "test"]

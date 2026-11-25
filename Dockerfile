FROM golang:1.24

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        make \
        curl \
        ca-certificates \
        git \
        grep \
    && rm -rf /var/lib/apt/lists/*

RUN go install github.com/golangci/golines@v0.15.0

RUN curl -sSfL https://golangci-lint.run/install.sh \
    | sh -s -- -b /usr/local/bin v2.15.0

WORKDIR /app
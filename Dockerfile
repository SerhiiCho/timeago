FROM golang:1.24-alpine

WORKDIR /app

RUN apk add --no-cache make

COPY . .

ENTRYPOINT ["sh"]

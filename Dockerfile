FROM golang:1.24-alpine

WORKDIR /var/www/html

RUN apk update && apk add --no-cache make

ENTRYPOINT ["sh"]

FROM golang:1.17.4-alpine

WORKDIR /var/www/app

COPY . .

CMD tail -f /dev/null

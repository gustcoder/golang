FROM golang:1.15

WORKDIR /app

COPY . .

RUN GOOS=linux go build -ldflags="-s -w"

CMD ["./app"]

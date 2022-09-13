FROM golang:1.17-alpine

WORKDIR /app

COPY . .

RUN go build -o client

EXPOSE 8080

CMD ./client
FROM golang:1.19-alpine3.16 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux go build -o /telebot cmd/main.go

FROM amd64/alpine:latest

WORKDIR /

COPY --from=builder /telebot /telebot
CMD ["/telebot"]
FROM golang:latest

RUN mkdir -p /app
WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . /app

CMD ["go", "run", "cmd/main.go"]
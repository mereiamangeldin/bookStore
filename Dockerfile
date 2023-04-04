FROM golang:latest

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
WORKDIR /cmd

CMD ["go","run","main.go"]
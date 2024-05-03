FROM golang:1.22.0-alpine3.19
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o api ./cmd/

CMD ["./api"]

EXPOSE 8080
FROM golang:1.24

WORKDIR /app

COPY . .

RUN go mod download

COPY . .

RUN go build -o main ./cmd

CMD ["./main"]
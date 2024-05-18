FROM golang:latest

WORKDIR /root

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main cmd/main.go

ENV PORT=50051
ENV LOG_LEVEL=debug

EXPOSE $PORT

CMD ["./main"]
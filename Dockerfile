FROM golang:1.21

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o main ./cmd/main.go

CMD ["/app/main"]

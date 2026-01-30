FROM golang:1.22-alpine
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o payment cmd/main.go
CMD ["./payment"]
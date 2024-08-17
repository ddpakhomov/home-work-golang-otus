# Этап сборки
FROM golang:1.20 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Этап запуска
FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/main /app/main
COPY config.yml /app
RUN chmod +x /app/main
EXPOSE 8081

CMD ["/app/main"]
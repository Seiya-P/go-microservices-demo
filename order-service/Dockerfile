# syntax=docker/dockerfile:1
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o order-service

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/order-service .
EXPOSE 8080
CMD ["/app/order-service"] 
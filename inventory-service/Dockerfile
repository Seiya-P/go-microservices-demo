# syntax=docker/dockerfile:1
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o inventory-service

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/inventory-service .
CMD ["/app/inventory-service"] 
# syntax=docker/dockerfile:1
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o notification-service

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/notification-service .
CMD ["/app/notification-service"] 
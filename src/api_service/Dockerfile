# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.21.0-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o api_service /app/src/api_service/cmd

# Run stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/api_service .
COPY --from=builder /app/src/api_service/app.env .

EXPOSE 8080
CMD [ "/app/api_service", "--port", "8080", "--host", "" ]
# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.21.0-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o redirection_service /app/src/redirection_service/cmd

# Run stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/redirection_service .
COPY --from=builder /app/src/redirection_service/app.env .

EXPOSE 8080
CMD [ "/app/redirection_service", "--port", "8080", "--host", "" ]
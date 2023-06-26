# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.20-alpine3.15 AS builder
WORKDIR /app
COPY . .
RUN go build /app/cmd/redirection-server

# Run stage
FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/redirection-server .
COPY --from=builder /app/wait-for.sh .

EXPOSE 8080
CMD [ "/app/api-server", "--port", "8080" ]
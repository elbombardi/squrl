# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.20-alpine3.15 AS builder
WORKDIR /app
COPY . .
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go build /app/cmd/demo-data

# Run stage
FROM alpine:3.15
WORKDIR /app
COPY --from=builder /go/bin/migrate .
COPY --from=builder /app/wait-for.sh .
COPY --from=builder /app/orders_service/db/migration/ migration
COPY --from=builder /app/demo-data .
COPY --from=builder /app/mock_data /app/mock_data

RUN ls /app

EXPOSE 8080
CMD [ "/app/migrate"]
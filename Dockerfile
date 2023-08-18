FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o bin/calculatepacks ./internal/api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bin/calculatepacks .

ENTRYPOINT ["./calculatepacks"]

# Build stage
FROM golang:1.23.6-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o api .

# Final stage
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/api .

EXPOSE 8080

CMD ["./api"]

FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o iaq-producer ./cmd/iaq-producer

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/iaq-producer .
CMD ["./iaq-producer"]

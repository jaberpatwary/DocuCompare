# Build stage
FROM golang:1.25-alpine AS builder

RUN apk add --no-cache gcc musl-dev tesseract-ocr-dev leptonica-dev

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=1 go build -o /app/server ./src/main.go

# Production stage
FROM alpine:3.20

RUN apk add --no-cache \
    tesseract-ocr \
    tesseract-ocr-data-eng \
    tesseract-ocr-data-ben \
    ca-certificates

WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/frontend ./frontend
COPY --from=builder /app/.env .

RUN mkdir -p /app/frontend/uploads

EXPOSE 1111

CMD ["./server"]

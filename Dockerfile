# === BUILD STAGE ===

FROM golang:1.24 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

# === FINAL STAGE ===

FROM gcr.io/distroless/base
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
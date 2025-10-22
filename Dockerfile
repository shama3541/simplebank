# Build stage
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# ⚡ Build a static binary (this is the key line)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/main ./main

# Final minimal image
FROM alpine:3.19

WORKDIR /app
COPY --from=builder /app/main .

COPY app.env .


# Run your binary
CMD ["/app/main"]

# -------- Build Stage --------
    FROM golang:1.23.6 AS builder

    WORKDIR /app
    
    # Copy and download dependencies
    COPY go.mod go.sum ./
    RUN go mod download
    
    # Copy source code
    COPY . .
    
    # ⚡ Build a static binary
    RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/main ./main
    
    # ✅ Download and install migrate binary
    RUN apt-get update && apt-get install -y curl tar && \
        curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz -o migrate.tar.gz && \
        tar -xvzf migrate.tar.gz && \
        mv migrate /usr/local/bin/migrate && \
        chmod +x /usr/local/bin/migrate
    
    # -------- Final Minimal Image --------
    FROM alpine:3.19
    
    WORKDIR /app
    
    # Copy built artifacts
    COPY --from=builder /app/main .
    COPY --from=builder /usr/local/bin/migrate ./migrate
    COPY app.env .
    COPY db/migration ./migration
    COPY app.sh .
    RUN chmod +x /app/app.sh
    CMD ["/app/main"]
    ENTRYPOINT ["/app/app.sh"]
    
    
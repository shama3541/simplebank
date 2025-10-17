# Use official Go image
FROM golang:1.23.6 AS builder

# Create a working directory
WORKDIR /app

# Copy all source code
COPY . .

# Build from the correct folder
RUN go build -o main ./main


FROM  alpine:3.13

WORKDIR /app

COPY --from=BUILDER /app/main/ .
# Set the command to run your binary
CMD ["/app/main"]

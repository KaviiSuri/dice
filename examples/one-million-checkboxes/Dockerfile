FROM golang:1.21 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM node:18 AS frontend-builder
WORKDIR /web

COPY web/package*.json ./

RUN npm install

COPY web/ .

RUN npm run build

FROM alpine:latest

# Install ca-certificates
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Go binary from the builder stage
COPY --from=builder /app/main .

# Copy the Vite-built static files from the frontend-builder stage
COPY --from=frontend-builder /web/dist ./web/dist

# Expose port 8000 for the Go server
EXPOSE 8000

# Start the Go server
CMD ["./main"]

FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN PLAYWRIGHT_SKIP_BROWSER_DOWNLOAD=1 CGO_ENABLED=0 GOOS=linux go build -o server .

FROM alpine:3.19
RUN apk add --no-cache \
    chromium \
    nss \
    freetype \
    harfbuzz \
    ttf-freefont \
    ca-certificates
WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/static ./static
COPY --from=builder /app/swagger ./swagger
COPY config.docker.toml config.toml
ENV CHROMIUM_PATH=/usr/bin/chromium-browser
EXPOSE 5555
CMD ["./server"]

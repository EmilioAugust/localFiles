# Stage 1
FROM golang:1.25-alpine AS builder
WORKDIR /app

ENV key=value

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o server ./app/cmd/server

# Stage 2
FROM alpine:3.18
WORKDIR /app

# Копируем бинарник
COPY --from=builder /app/server .

COPY app/web ./web
COPY .env ./

EXPOSE 8080
CMD ["./server"]

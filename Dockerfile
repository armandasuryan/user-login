# syntax=docker/dockerfile:1

# Mengatur versi Go
ARG GO_VERSION=1.23

# Menggunakan Go dengan alpine sebagai base image
FROM golang:${GO_VERSION}-alpine

# Mengatur mode production secara default
ENV GO_ENV production

# Menentukan direktori kerja untuk aplikasi
WORKDIR /usr/src/app

# Copy go.mod dan go.sum agar dependency dapat di-cache
COPY go.mod go.sum ./

# Menggunakan cache mount untuk mempercepat build selanjutnya
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# Copy seluruh aplikasi ke container
COPY . .

# Build aplikasi Go
RUN go build -o app ./cmd/main.go

# Menggunakan user non-root untuk menjalankan aplikasi
USER nobody

# Mengekspos port aplikasi (sesuai aplikasi Go kamu)
EXPOSE 8080

# Menjalankan aplikasi
CMD ["./app"]

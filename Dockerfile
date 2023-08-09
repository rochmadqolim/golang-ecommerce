# Gunakan image dasar Golang
FROM golang:1.20

# Set direktori kerja di dalam kontainer
WORKDIR /app

# Salin kode Go dan file go.mod/go.sum ke direktori kerja kontainer
COPY . .

# Build aplikasi Go
RUN go build -o main .

# Menjalankan aplikasi saat kontainer berjalan
CMD ["./main"]

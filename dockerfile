# Gunakan base image Go
FROM golang:1.21

# Set direktori kerja di dalam container
WORKDIR /app

# Copy semua file ke dalam container
COPY . .

# Download dependensi
RUN go mod tidy

# Build aplikasi
RUN go build -o main .

# Jalankan aplikasi
CMD ["./main"]


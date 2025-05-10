# Gunakan image resmi Go sebagai base image
FROM golang:1.20-alpine


# Set working directory di dalam container
WORKDIR /app

# Copy file go.mod dan go.sum untuk mendownload dependensi
COPY go.mod go.sum ./

# Menjalankan perintah untuk mendownload dependensi
RUN go mod tidy

# Copy semua file dari direktori lokal ke dalam container
COPY . .

# Build aplikasi Go
RUN go build -o main .

# Expose port 8080 agar bisa diakses di luar container
EXPOSE 8080

# Jalankan aplikasi Go
CMD ["./main"]

# Sử dụng image golang alpine nhỏ gọn làm base
FROM golang:1.24-alpine AS builder

# Thiết lập thư mục làm việc
WORKDIR /app

# Sao chép go.mod và go.sum trước để tận dụng cache
COPY go.mod go.sum ./

# Tải dependencies
RUN go mod download

# Sao chép mã nguồn
COPY . .
# Build ứng dụng với cờ tối ưu hóa
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-s -w" -o main ./cmd/server

# Sử dụng image scratch siêu nhẹ cho runtime
FROM scratch

# Sao chép binary từ builder
COPY --from=builder /app/main /main

# Thiết lập lệnh chạy
EXPOSE 8080
ENTRYPOINT ["/main"]
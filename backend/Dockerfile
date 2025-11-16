FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN go mod init blogapi && go mod tidy
RUN go build -o blog-api main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/blog-api .
EXPOSE 8080
CMD ["./blog-api"]

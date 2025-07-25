FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY backend .
RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init
RUN go build -o server .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/server /root/server
COPY --from=builder /app/docs /root/docs
EXPOSE 8080
CMD ["./server"]
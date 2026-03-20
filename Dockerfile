FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /server ./cmd/server

FROM alpine:latest
COPY --from=builder /server /server
EXPOSE 50051
CMD ["/server"]

FROM golang:1.25.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/fool_auth_service
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

RUN ls -l main

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/cmd/fool_auth_service/main .
EXPOSE 8080
CMD ["./main"]
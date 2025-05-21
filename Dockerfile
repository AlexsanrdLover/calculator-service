FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 
ENV GOOS=linux

RUN go mod edit -replace calculator-service=/app
RUN go build -o calculator -ldflags="-s -w" ./cmd/server/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/calculator .
COPY --from=builder /app/docs ./docs

EXPOSE 8080 50051

CMD ["./calculator"]
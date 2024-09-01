FROM golang:1.23.0-alpine3.20 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o authentication .

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/authentication .
EXPOSE 9090
ENTRYPOINT ["./authentication"]

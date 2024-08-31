# FROM golang:1.22.5
# WORKDIR /app
# COPY go.mod go.sum ./
# RUN go mod download
# COPY . .
# RUN go build -o authentication .
# EXPOSE 9090
# ENTRYPOINT [ "./authentication" ]


FROM golang:1.23.0-alpine3.20 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o authentication .

FROM alpine:3.20
WORKDIR /app
# Copy the compiled binary from the builder stage
COPY --from=builder /app/authentication .
EXPOSE 9090
ENTRYPOINT ["./authentication"]

FROM golang:1.23.0-alpine3.20 AS pgbuild
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOARCH=amd64 go build -o go-postgres .

FROM alpine:3.20
WORKDIR /app
COPY --from=pgbuild /app/go-postgres .
EXPOSE 8080
ENTRYPOINT ["./go-postgres"]
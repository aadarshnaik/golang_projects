FROM golang:1.23.0-alpine3.20
WORKDIR /app
COPY . /opt/
RUN go build /opt/
EXPOSE 9090
ENTRYPOINT [ "./authentication" ]
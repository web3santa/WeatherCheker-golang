FROM golang:1.22.1-alpine3.19

WORKDIR /app
COPY go.mod .
COPY main.go .
COPY go.sum .
RUN go mod download
RUN go build -v -o /app/weather
CMD ["./weather"]
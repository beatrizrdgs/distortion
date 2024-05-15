FROM golang:1.21.5-alpine AS builder

WORKDIR /distortion

COPY . .

RUN go mod download
RUN go build -o distortion

EXPOSE 8080

CMD ["./distortion", "server"]
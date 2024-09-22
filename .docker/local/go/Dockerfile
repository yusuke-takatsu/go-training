FROM golang:1.21.5
ENV TZ=Asia/Tokyo

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
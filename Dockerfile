FROM golang:latest

LABEL maintainer="Anuj <underscoreanuj@gmail.com>"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV GO_API_PORT 8000

RUN go build

CMD ["./mux_api"]

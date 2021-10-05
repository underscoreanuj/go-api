FROM golang:latest

LABEL maintainer="Anuj <underscoreanuj@gmail.com>"

# specify the working directory
WORKDIR /app

# required for the package setup
COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

# specify the port at which the API server will listen for requests
ENV GO_API_PORT 8000

# build the project
RUN go build

# run the project
CMD ["./mux_api"]

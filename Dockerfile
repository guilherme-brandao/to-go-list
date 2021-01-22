FROM golang:1.14-alpine as builder

WORKDIR /to-go-list

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the app
RUN go build 

# Run the app
CMD ["./to-go-list"]

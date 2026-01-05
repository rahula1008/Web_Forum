FROM golang:1.25-bookworm AS base

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o go-backend

# Use 3000 as the port
EXPOSE 3000

CMD ["/build/go-backend"]

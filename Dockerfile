FROM golang:1.24.6-alpine3.22 AS build

WORKDIR /myapp

# Modules layer
COPY go.mod go.sum ./
RUN go mod download

# Build layer
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o ./bin/app ./cmd/app

FROM alpine:3.22 AS run

WORKDIR /myapp

COPY --from=build /myapp/bin/app ./bin/app
COPY .env .env

EXPOSE 8080

CMD ["/myapp/bin/app"]
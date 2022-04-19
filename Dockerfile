FROM golang:1.18.1-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
RUN go build -o /build/app

FROM alpine:3.15.4

WORKDIR /app

COPY --from=build /build/app main
COPY .env .env

EXPOSE 8080

ENTRYPOINT ["/app/main"]

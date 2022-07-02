
FROM golang:alpine3.16 AS build
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o app

# Final
FROM alpine:3.16
WORKDIR /

COPY --from=build /app/app .

EXPOSE 4000
ENTRYPOINT ["./app"]
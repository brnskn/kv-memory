FROM golang:alpine AS build

RUN apk update && \
    apk add curl \
            bash \
            ca-certificates && \
    rm -rf /var/cache/apk/*

WORKDIR /app

COPY . .

RUN go build cmd/server/main.go 

FROM alpine:latest

RUN apk --no-cache add ca-certificates bash

WORKDIR /app

COPY --from=build /app/main .

ENTRYPOINT ["./main"]
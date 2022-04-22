FROM golang:1.12.0-alpine3.9 AS builder

RUN mkdir /app

ADD . /app

RUN apk update && apk add --no-cache git

WORKDIR /app

RUN go build -o main .

FROM alpine:3.13

COPY --from=builder /app/main /app/

CMD ["/app/main"]

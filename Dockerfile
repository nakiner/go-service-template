FROM golang:1.23 AS builder

WORKDIR /go/src/app
COPY . .

COPY . /app

RUN make build && make modcache-clean

FROM alpine

RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/app/bin/service .

CMD ["./service"]

FROM golang:1.23 AS builder

WORKDIR /go/src/app
COPY . .

COPY . /app

RUN make build && make modcache-clean

FROM debian:bookworm-slim

RUN apt-get update && apt-get install git curl ca-certificates -y

COPY --from=builder /go/src/app/bin/service .

CMD ["./service"]

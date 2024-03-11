FROM golang:1.20-bullseye AS builder

EXPOSE 80

WORKDIR /usr/src

COPY . .

RUN go build -o fiber-go .

FROM debian:bullseye-slim

COPY --from=builder /usr/src/fiber-go /usr/local/bin/fiber-go

CMD ["fiber-go"]
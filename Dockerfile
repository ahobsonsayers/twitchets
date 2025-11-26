# Builder Image
FROM golang:1.24 AS builder

WORKDIR /twitchets

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v -o ./bin/ .

# Distribution Image
FROM alpine:latest

RUN apk add --no-cache libc6-compat

COPY --from=builder /twitchets/bin/twitchets /usr/bin/twitchets

WORKDIR /twitchets

EXPOSE 5656

ENTRYPOINT ["/usr/bin/twitchets"]

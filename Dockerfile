FROM golang:1.16-alpine AS builder
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -a -ldflags "-extldflags '-static -s'" -o stashbox main.go

FROM alpine:3.13

RUN set -ex \
    && apk add --no-cache ca-certificates
WORKDIR /app

COPY --from=builder /app/dist/stashbox /app/stashbox

CMD ["./stashbox"]

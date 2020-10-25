FROM golang:1.15.3-alpine as builder

WORKDIR /build/
COPY go.* /build/
RUN go mod download

COPY *.go /build/
ENV CGO_ENABLED=0

RUN go build -o /bin/horrible-crawler ./


FROM alpine:latest
RUN apk add tzdata
COPY --from=builder /bin/horrible-crawler /bin/
ENTRYPOINT /bin/horrible-crawler

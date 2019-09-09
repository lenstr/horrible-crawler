FROM golang:1.13-alpine as builder

WORKDIR /build/
COPY go.* /build/
RUN go mod download

COPY main.go /build/main.go
ENV CGO_ENABLED=0

RUN go build -o /bin/horrible-crawler main.go


FROM alpine:latest
RUN apk add tzdata
COPY --from=builder /bin/horrible-crawler /bin/
ENTRYPOINT /bin/horrible-crawler

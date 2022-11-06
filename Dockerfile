FROM golang:1.18 AS builder
WORKDIR /app
COPY cmd/main.go go.mod ./
RUN CGO_ENABLED=0 GOOS=linux go build -o lb .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=builder /app/athena .
ENTRYPOINT [ "/root/athena" ]
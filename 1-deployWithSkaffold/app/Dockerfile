FROM golang:1.15.3-alpine3.12 as builder
COPY server.go .
RUN go build -o /app server.go

FROM alpine:3.12.1
COPY --from=builder /app .
CMD "./app"

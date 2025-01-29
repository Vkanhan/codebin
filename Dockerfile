FROM golang:1.22-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o codebin ./web

FROM alpine:latest
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/codebin .
COPY --from=builder /app/ui /app/ui
COPY .env .env

CMD ["./codebin"]
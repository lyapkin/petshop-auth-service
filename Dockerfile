FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN mkdir -p bin/main && CGO_ENABLED=0 GOOS=linux go build -o bin/main/main ./cmd/main/main.go

FROM alpine:latest

RUN addgroup -S app -g 1000 && \
    adduser -S -G app -u 1000 -s /sbin/nologin auth

WORKDIR /app
RUN chown auth:app /app

COPY --chown=root:app --from=builder /app/bin/main/main .
RUN chmod 550 -R /app

USER auth

EXPOSE 8000

ENTRYPOINT ["./main"]
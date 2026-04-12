FROM golang:1.22-alpine AS builder

WORKDIR /src

RUN apk add --no-cache ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/taskflow-api .

FROM migrate/migrate:v4.18.3 AS migrate

FROM alpine:3.20 AS runtime

WORKDIR /app

RUN addgroup -S app && adduser -S app -G app \
    && apk add --no-cache ca-certificates wget

COPY --from=builder /out/taskflow-api /app/taskflow-api
COPY --from=migrate /usr/local/bin/migrate /usr/local/bin/migrate
COPY migrations /app/migrations
COPY docker-entrypoint.sh /app/docker-entrypoint.sh
RUN sed -i 's/\r$//' /app/docker-entrypoint.sh && chmod +x /app/docker-entrypoint.sh

USER app

EXPOSE 5000

CMD ["/app/docker-entrypoint.sh"]

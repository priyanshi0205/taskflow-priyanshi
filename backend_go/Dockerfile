FROM golang:1.22-alpine AS builder

WORKDIR /src

RUN apk add --no-cache ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/taskflow-api .

FROM alpine:3.20 AS runtime

WORKDIR /app

RUN addgroup -S app && adduser -S app -G app \
    && apk add --no-cache ca-certificates wget

COPY --from=builder /out/taskflow-api /app/taskflow-api

USER app

EXPOSE 5000

CMD ["/app/taskflow-api"]
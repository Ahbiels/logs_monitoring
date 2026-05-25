FROM golang:1.26 AS builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod tidy

COPY . .

RUN go build -o log_monitoring .

FROM debian:bookworm-slim AS runner

WORKDIR /usr/src/app

RUN apt-get update && \
    apt-get install -y ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /usr/src/app/log_monitoring /usr/src/app/log_monitoring

COPY key /usr/src/app/key

EXPOSE 7777

CMD ["./log_monitoring"]
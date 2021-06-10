FROM golang:1.16-alpine as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .

FROM scratch

COPY --from=builder /app /usr/bin/

HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=2 CMD curl -f http://localhost/ || exit 1

ENTRYPOINT ["/usr/bin/weather"]
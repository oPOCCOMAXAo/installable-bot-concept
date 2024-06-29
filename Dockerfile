FROM golang:alpine as builder
RUN apk add --no-cache gcc musl-dev
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /bin/app cmd/app/main.go

FROM alpine
COPY --from=builder /bin/app /bin/app
RUN mkdir -p /data
ENV SERVER_PORT=8080 \
    GIN_MODE=release \
    DB_LOCAL_DSN=/data/db.sqlite3
EXPOSE 8080
ENTRYPOINT ["/bin/app"]

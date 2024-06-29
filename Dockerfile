FROM golang as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /bin/app cmd/app/main.go

FROM alpine
COPY --from=builder /bin/app /bin/app
ENV SERVER_PORT=8080 \
    GIN_MODE=release \
    DB_DSN=/data/db.sqlite3
EXPOSE 8080
ENTRYPOINT ["/bin/app"]

FROM golang:1.22

WORKDIR /app

COPY . .

RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

CMD ["tail", "-f", "/dev/null"]

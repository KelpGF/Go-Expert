FROM golang:1.22

WORKDIR /app

RUN apt update && apt install sqlite3 -y

COPY . .

RUN go mod tidy

CMD ["tail", "-f", "/dev/null"]

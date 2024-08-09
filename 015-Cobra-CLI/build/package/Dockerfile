FROM golang:1.22

WORKDIR /app

COPY . .

RUN go install github.com/spf13/cobra@latest

CMD ["tail", "-f", "/dev/null"]

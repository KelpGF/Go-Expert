FROM golang:1.22

RUN apt-get update && apt-get install apache2-utils -y

WORKDIR /app

COPY . .

CMD ["tail", "-f", "/dev/null"]

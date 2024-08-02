FROM golang:1.22

WORKDIR /app

COPY . .

RUN apt-get update && \
    apt-get install -y protobuf-compiler  && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1

# Evans - gRPC client: https://github.com/ktr0731/evans
RUN go install github.com/ktr0731/evans@latest

CMD ["tail", "-f", "/dev/null"]

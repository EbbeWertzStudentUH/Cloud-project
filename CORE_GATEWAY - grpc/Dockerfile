FROM golang:latest

WORKDIR /app

# protoc generator
RUN apt-get update && apt-get install -y \
    protobuf-compiler \
    && rm -rf /var/lib/apt/lists/*

RUN command protoc --go_out=. --go-grpc_out=. facade.proto

COPY . .

RUN go mod tidy

EXPOSE 3006

ENTRYPOINT ["/app/docker_entrypoint.sh"]

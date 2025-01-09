FROM golang:latest

WORKDIR /app

# protoc generator
RUN apt-get update && apt-get install -y protobuf-compiler && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

COPY . .
RUN command protoc --go_out=. --go-grpc_out=. facade.proto


RUN go mod tidy

EXPOSE 3006

ENTRYPOINT ["/app/docker_entrypoint.sh"]

FROM golang:latest AS base

WORKDIR /app

RUN apt-get update && apt-get install -y protobuf-compiler && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
    apt-get clean && rm -rf /var/lib/apt/lists/*


FROM base AS proto

COPY facade.proto .
RUN protoc --go_out=. --go-grpc_out=. facade.proto



FROM golang:latest AS runtime

WORKDIR /app
COPY --from=proto /app /app
COPY . .
RUN go mod tidy

EXPOSE 3006

ENTRYPOINT ["/app/docker_entrypoint.sh"]

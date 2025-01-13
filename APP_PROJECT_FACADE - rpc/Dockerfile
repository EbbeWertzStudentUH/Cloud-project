FROM golang:latest

WORKDIR /app
COPY . .
RUN go mod tidy

EXPOSE 3006

ENTRYPOINT ["/app/docker_entrypoint.sh"]

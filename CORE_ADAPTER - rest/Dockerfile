FROM rust:1.72 AS builder

WORKDIR /app

# dependencies en build
COPY . .
RUN cargo fetch
RUN cargo build --release

ENTRYPOINT ["/app/docker_entrypoint_build.sh"]




FROM ubuntu:latest AS runtime

WORKDIR /app
COPY --from=builder /app/target/release/GATEWAY-rest .
COPY ./docker_entrypoint_runtime.sh ./docker_entrypoint_runtime.sh
COPY ./.env ./.env

EXPOSE 3001

ENTRYPOINT ["/app/docker_entrypoint_runtime.sh"]


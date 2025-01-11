FROM rust:1.72 AS base

WORKDIR /app
RUN apt-get update && apt-get install -y \
    protobuf-compiler \
    && rm -rf /var/lib/apt/lists/*

COPY Cargo.toml ./
COPY src/main.rs ./src/main.rs
RUN cargo fetch

FROM base AS builder

COPY . .

RUN cargo build --release


FROM ubuntu:latest AS runtime

WORKDIR /app
COPY --from=builder /app/target/release/ADAPTER-rest .
COPY ./docker_entrypoint_runtime.sh ./docker_entrypoint_runtime.sh
COPY ./.env ./.env

EXPOSE 3001

ENTRYPOINT ["/app/docker_entrypoint_runtime.sh"]

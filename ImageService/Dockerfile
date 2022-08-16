FROM rust

WORKDIR /app

COPY ./ ./
RUN mkdir -p /public

RUN apt update
RUN apt install pkg-config

RUN rustup override set nightly
RUN cargo build --release


CMD ["./target/release/ImageService"]
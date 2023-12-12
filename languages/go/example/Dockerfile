FROM rust:1.74-bookworm AS builder
WORKDIR /data
RUN git config --global user.email "bar@foo.com" \
  && git config --global user.name "builder" \
  && git clone --single-branch https://github.com/bitwarden/sdk
WORKDIR /data/sdk
RUN cargo build \
    && apt update \
    && apt install -y nodejs npm \
    && npm install \
    && npm run schemas


FROM golang:1.19-bookworm as go
WORKDIR /data
COPY --from=builder /data/sdk/target/debug/libbitwarden_c.so /usr/local/lib/libbitwarden_c.so
COPY --from=builder /data/sdk/languages/go /data/sdk
COPY . /data/sdk/my-app/
WORKDIR /data/sdk/my-app
RUN go mod tidy \
   && go build -o example


FROM debian:bookworm-slim
RUN apt update && apt install -y ca-certificates
COPY --from=builder /data/sdk/target/debug/libbitwarden_c.so /usr/local/lib/libbitwarden_c.so
COPY --from=go /data/sdk/my-app/example /usr/local/bin/example
ENV LD_LIBRARY_PATH=/usr/local/lib/
ENTRYPOINT [ "/usr/local/bin/example" ]

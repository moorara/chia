# BUILD STAGE
FROM golang:1.14-alpine as builder
RUN apk add --no-cache git
WORKDIR /repo
COPY . .
ENV CGO_ENABLED=0
RUN wget -qO - https://git.io/JeCX6 | sh
RUN cherry build -cross-compile=false

# FINAL STAGE
FROM alpine:3.12
RUN apk add --no-cache ca-certificates
COPY --from=builder /repo/bin/chia /usr/local/bin/
RUN chown -R nobody:nogroup /usr/local/bin/chia
USER nobody
ENTRYPOINT [ "chia" ]

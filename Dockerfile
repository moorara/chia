# BUILD STAGE
FROM golang:1.12-alpine as builder
RUN apk add --no-cache git
WORKDIR /workspace
COPY . .
RUN ./scripts/build.sh --main main.go --binary chia

# FINAL STAGE
FROM alpine:3.9
RUN apk add --no-cache ca-certificates
COPY --from=builder /workspace/chia /usr/local/bin/
RUN chown -R nobody:nogroup /usr/local/bin/chia
USER nobody
ENTRYPOINT [ "chia" ]

# BUILD STAGE
FROM golang:1.11-alpine as builder
RUN apk add --no-cache git
WORKDIR /workspace
COPY . .
RUN ./scripts/build.sh --main main.go --binary chia

# FINAL STAGE
FROM alpine:3.8
RUN apk add --no-cache ca-certificates
COPY --from=builder /workspace/chia /usr/local/bin/
RUN chown -R nobody:nogroup /usr/local/bin/chia
USER nobody
CMD chia

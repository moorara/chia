# BUILD STAGE
FROM golang:1.11-alpine as builder
RUN apk add --no-cache make git
WORKDIR /go/src/github.com/moorara/chia/
COPY . .
RUN make build && \
    cp chia /

# FINAL STAGE
FROM alpine:3.8
RUN apk add --no-cache ca-certificates
COPY --from=builder /chia /usr/local/bin/
RUN chown -R nobody:nogroup /usr/local/bin/chia
USER nobody
CMD chia

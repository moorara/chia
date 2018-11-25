[![Build Status][travisci-image]][travisci-url]
[![Go Report Card][goreport-image]][goreport-url]

# Chia

This is a **work-in-progress** utility for facilitating and automating API testing.

## TO-DO

Supporting the following API paradigms:

  - [ ] REST
  - [ ] gRPC
  - [ ] GraphQL
  - [ ] Messaging (NATS)

Supporting the following features:

  - [ ] HTTPS
  - [ ] MTLS

## Commands

| Command                        | Description                                          |
|--------------------------------|------------------------------------------------------|
| `make run`                     | Run the application locally                          |
| `make build`                   | Build the binary locally                             |
| `make build-all`               | Build the binary locally for all supported platforms |
| `make test`                    | Run the unit tests                                   |
| `make test-short`              | Run the unit tests using `-short` flag               |
| `make coverage`                | Run the unit tests with coverage report              |
| `make docker`                  | Build Docker image                                   |
| `make push`                    | Push built image to registry                         |


[travisci-url]: https://travis-ci.org/moorara/chia
[travisci-image]: https://travis-ci.org/moorara/chia.svg?branch=master

[goreport-url]: https://goreportcard.com/report/github.com/moorara/chia
[goreport-image]: https://goreportcard.com/badge/github.com/moorara/chia

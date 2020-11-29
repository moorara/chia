[![Build Status][workflow-image]][workflow-url]
[![Go Report Card][goreport-image]][goreport-url]
[![Test Coverage][coverage-image]][coverage-url]
[![Maintainability][maintainability-image]][maintainability-url]

# Chia

This is a **WORK-IN-PROGRESS**.

Chia facilitates and automates API testing.

## Quick Start

### Examples

## TO-DO

Supporting the following API paradigms:

  - [ ] REST
  - [ ] gRPC
  - [ ] GraphQL
  - [ ] Messaging (NATS)

Supporting the following features:

  - [ ] HTTPS
  - [ ] mTLS

## Commands

| Command            | Description                                          |
|--------------------|------------------------------------------------------|
| `make build`       | Build the binary locally                             |
| `make build-all`   | Build the binary locally for all supported platforms |
| `make test`        | Run the unit tests                                   |
| `make test-short`  | Run the unit tests using `-short` flag               |
| `make coverage`    | Run the unit tests with coverage report              |
| `make docker`      | Build Docker image                                   |
| `make push`        | Push built image to registry                         |
| `make save-docker` | Save built image to disk                             |
| `make load-docker` | Load saved image from disk                           |


[workflow-url]: https://github.com/moorara/chia/actions
[workflow-image]: https://github.com/moorara/chia/workflows/main/badge.svg
[goreport-url]: https://goreportcard.com/report/github.com/moorara/chia
[goreport-image]: https://goreportcard.com/badge/github.com/moorara/chia
[coverage-url]: https://codeclimate.com/github/moorara/chia/test_coverage
[coverage-image]: https://api.codeclimate.com/v1/badges/e00a5fd21864f3c3f2ec/test_coverage
[maintainability-url]: https://codeclimate.com/github/moorara/chia/maintainability
[maintainability-image]: https://api.codeclimate.com/v1/badges/e00a5fd21864f3c3f2ec/maintainability

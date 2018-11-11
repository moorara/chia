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

| Command                        | Description                             |
|--------------------------------|-----------------------------------------|
| `make dep`                     | Install and updates dependencies        |
| `make run`                     | Run the application locally             |
| `make build`                   | Build the application binary locally    |
| `make test`                    | Run the unit tests                      |
| `make coverage`                | Run the unit tests with coverage report |
| `make docker`                  | Build Docker image                      |
| `make push`                    | Push built image to registry            |

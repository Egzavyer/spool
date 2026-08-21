# Spool

A distributed background job processing system built in Go, with persistent job storage, concurrent workers, and asynchronous task execution.

## Overview

Server receives HTTP request from client to perform a job and validates it. A valid job request is stored in the database with a PENDING status. Workers claim jobs and complete them. The system is fault-tolerant and resistant. Workers can complete multiple tasks in parallel.

## Protocol

The communication protocol and between the client and the server is defined using [protobuf](https://protobuf.dev) with the [Buf CLI](https://buf.build/docs/) tool to generate the structs. Run `buf lint` to check the current `.proto` files agains the linter and `buf generate` to compile them and generate the Go code they describe. Use `just buf-breaking` to check for breaking changes between the current changes and the previous ones. The `buf` workspace is defined in `buf.yaml` and the specifications for generating the code is in `buf.gen.yaml`.

The application uses the `connectRPC` protocol for communication between the client and the server. This is done using the [Connect-Go](https://connectrpc.com/) library which allows to define type-safe `Remote Procedure Calls` using our `protobuf` definition.

## Database

This project uses a PostgreSQL database inside of a [docker container](https://docs.docker.com/guides/postgresql/#docker-compose-configuration). The database can be spun up using `docker compose up` and lives on port 5432. Once the database is running, you can connect to it by using the `just db-connect` command.

## Database Migrations

Migrations to the database are applied using the `golang-migrate` library. The server always tries to migrate up to the latest schema on startup. Migrations can also be applied or reversed manually using the [CLI tool](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) by running `just db-migrate <up|down> [amount]`. Leaving amount blank will apply or reverse all migrations. To create a new migration, run `just db-create <name>`. See [documentation](https://github.com/golang-migrate/migrate/tree/master) for more details.

## Tools

- [just](https://github.com/casey/just) : Open-source command runner
- [golang-migrate](https://github.com/golang-migrate/migrate) : Open-source database schema migration tool
- [Buf CLI](https://github.com/bufbuild/buf) : Open-source protobuf toolchain
- [Docker Compose](https://docs.docker.com/compose/) : Define and run multi-container applications
- [Connect-Go](https://connectrpc.com/docs/go/getting-started/) : Open-source library for connectRPC

## Next Steps

- Look into Connect-Go field validation

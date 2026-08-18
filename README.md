# Spool

## Overview

Server receives HTTP request from client to perform a job and validates it. A valid job request is store in the database with a PENDING status. Workers claim jobs and complete them. The system is fault-tolerant and resistant. Workers can complete multiple tasks in parallel.

## Database

This project uses a PostgreSQL database inside of a [docker container](https://docs.docker.com/guides/postgresql/#docker-compose-configuration). The database can be spun up using `docker compose up` and lives on port 5432. Once the database is running, you can connect to it by using the `just db-connect` command.

## Database Migrations

Migrations to the database are applied using the `golang-migrate` library. The server always tries to migrate up to the latest schema on startup. Migrations can also be applied or reversed manually using the [CLI tool](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) by running `just db-migrate <up|down> [amount]`. Leaving amount blank will apply or reverse all migrations. To create a new migration, run `just db-create <name>`. See [documentation](https://github.com/golang-migrate/migrate/tree/master) for more details.

## Next Steps

- Define basic client-server communication protocol. Look into [protobuf](https://protobuf.dev/getting-started/gotutorial) for protocol definition

set shell := ["powershell","-c"]
set dotenv-load := true

pg_url := env("PG_URL")

buf-breaking:
    buf breaking --against .git#subdir=proto

db-migrate DIRECTION AMOUNT="":
    migrate -source file://db/migrations -database {{pg_url}} {{DIRECTION}} {{AMOUNT}}

db-create NAME:
    migrate create -ext sql -dir db/migrations -seq {{NAME}}

db-connect:
    docker exec -it postgres psql -U postgres -d postgres

create-migration TITLE:
    migrate create -ext sql -dir db/migrations -seq {{TITLE}}

run-server: build-server
    ./build/server/server.exe

build-server:
    go build -o ./build/server/server.exe ./cmd/server

run-client: build-client
    ./build/client/client.exe

build-client:
    go build -o ./build/client/client.exe ./cmd/client

run-worker: build-worker
    ./build/worker/worker.exe

build-worker:
    go clean
    go build -o ./build/worker/worker.exe ./cmd/worker

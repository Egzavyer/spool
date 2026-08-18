set shell := ["powershell","-c"]

run-server: build-server
    ./build/server/server.exe

build-server:
    go clean
    go build -o ./build/server/server.exe ./cmd/server

run-client: build-client
    ./build/client/client.exe

build-client:
    go clean
    go build -o ./build/client/client.exe ./cmd/client

run-worker: build-worker
    ./build/worker/worker.exe

build-worker:
    go clean
    go build -o ./build/worker/worker.exe ./cmd/worker

module spool

go 1.26.6

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.36.12-20260709200747-435963d16310.1
	connectrpc.com/connect v1.20.0
	github.com/golang-migrate/migrate/v4 v4.19.1
	github.com/joho/godotenv v1.5.1
	google.golang.org/protobuf v1.36.12
)

require github.com/lib/pq v1.10.9 // indirect

tool (
	connectrpc.com/connect/cmd/protoc-gen-connect-go
	google.golang.org/protobuf/cmd/protoc-gen-go
)

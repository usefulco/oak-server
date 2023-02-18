package proto

//go:generate protoc --go_out=paths=source_relative:../../internal/provider --go-grpc_out=paths=source_relative:../../internal/provider provider.proto
//go:generate protoc --go_out=paths=source_relative:../../internal/ingest --go-grpc_out=paths=source_relative:../../internal/ingest ingest.proto
//go:generate protoc --go_out=paths=source_relative:../../internal/live --go-grpc_out=paths=source_relative:../../internal/live live.proto

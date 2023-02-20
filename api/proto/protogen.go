package proto

//go:generate protoc --go_out=paths=source_relative:../../internal/aws --go-grpc_out=paths=source_relative:../../internal/aws aws.proto

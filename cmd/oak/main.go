package main

import (
	"log"
	"net"

	pbIngest "github.com/usefulco/oak-server/api/ingest"
	"github.com/usefulco/oak-server/internal/ingest"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	grpcIngestServer := ingest.NewServer()
	pbIngest.RegisterIngestServiceServer(server, grpcIngestServer)

	log.Printf("server listening")

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

package main

import (
	"log"
	"net"

	"github.com/usefulco/oak-server/internal/aws"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	aws.RegisterAWSServiceServer(server, aws.NewAWSProviderServer())

	log.Printf("server listening")

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

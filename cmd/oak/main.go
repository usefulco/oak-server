package main

import (
	"log"
	"net"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/usefulco/oak-server/internal/ingest"
	"google.golang.org/grpc"
)

// TODO:
// - move aws session creation to better place
// - create register provider ... provider.Register(config) // registers AWS Provider

func main() {
	awsSession := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-2"),
	}))

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	grpcIngestServer := ingest.NewServer(awsSession)
	ingest.RegisterIngestServiceServer(server, grpcIngestServer)

	log.Printf("server listening")

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

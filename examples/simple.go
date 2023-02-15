package main

import (
	"context"
	"log"
	"time"

	pbIngest "github.com/usefulco/oak-server/api/ingest"
	pbprovider "github.com/usefulco/oak-server/api/provider"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pbIngest.NewIngestServiceClient(conn)

	// Contact the server and print out response
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CreateIngest(ctx, &pbIngest.CreateIngestRequest{
		Provider:     pbprovider.Provider_AWS,
		Name:         "test_stream_ingestion",
		SourceName:   "mikes_computer",
		SourceIpAddr: "192.168.1.1",
	})
	if err != nil {
		log.Fatalf("could not create: %v", err)
	}

	log.Printf("Success!: %v", r.GetName())
}

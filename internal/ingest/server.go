package ingest

import (
	"context"
	"log"

	pbIngest "github.com/usefulco/oak-server/api/ingest"
)

type ingestServer struct {
	pbIngest.UnimplementedIngestServiceServer
}

func (s *ingestServer) CreateIngest(ctx context.Context, r *pbIngest.CreateIngestRequest) (*pbIngest.Ingest, error) {
	log.Printf("Recieved CreateIngest request")

	return &pbIngest.Ingest{
		Id:   "1234",
		Name: "Sample",
	}, nil
}

func NewServer() pbIngest.IngestServiceServer {
	s := &ingestServer{}

	return s
}

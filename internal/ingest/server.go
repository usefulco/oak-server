package ingest

import (
	"context"
	"log"

	"github.com/davecgh/go-spew/spew"
	pbIngest "github.com/usefulco/oak-server/api/ingest"
)

// TODO:
// - update awsProvider to be providers[]
// - update server methods to require preferred provider, reference from providers slice for each call
// - move test data into gRPC client call so it's not hard-coded here
// - add proper response type

type IngestServer struct {
	pbIngest.UnimplementedIngestServiceServer
	awsProvider Provider
}

func (s *IngestServer) CreateIngest(ctx context.Context, r *pbIngest.CreateIngestRequest) (*pbIngest.Ingest, error) {
	result, err := s.awsProvider.Create(&ProviderCreateInput{
		Name:              "test_stream_ingestion",
		SourceName:        "mikes_computer",
		PermittedSourceIP: "192.168.1.1",
	})
	if err != nil {
		log.Fatalf("failed to create aws setup: %v", err)
	}

	spew.Dump(result)

	return &pbIngest.Ingest{
		Id:   "1234",
		Name: "Sample",
	}, nil
}

func NewServer() pbIngest.IngestServiceServer {
	s := &IngestServer{
		awsProvider: &AwsProvider{},
	}

	return s
}

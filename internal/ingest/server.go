package ingest

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/davecgh/go-spew/spew"
	pbIngest "github.com/usefulco/oak-server/api/ingest"
	pbprovider "github.com/usefulco/oak-server/api/provider"
)

// TODO:
// - Abstract new service scoffolding to module
// - update server methods to require preferred provider, reference from providers slice for each call
// - add proper response type

type IngestServer struct {
	pbIngest.UnimplementedIngestServiceServer
	providers map[pbprovider.Provider]Provider
}

func (s *IngestServer) CreateIngest(ctx context.Context, r *pbIngest.CreateIngestRequest) (*pbIngest.Ingest, error) {
	result, err := s.providers[r.Provider].Create(&ProviderCreateInput{
		Name:              r.Name,
		SourceName:        r.SourceName,
		PermittedSourceIP: r.SourceIpAddr,
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

func NewServer(awsSession *session.Session) pbIngest.IngestServiceServer {
	return &IngestServer{
		providers: map[pbprovider.Provider]Provider{
			pbprovider.Provider_AWS: NewAwsProvider(awsSession),
		},
	}
}

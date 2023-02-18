package ingest

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/davecgh/go-spew/spew"
	"github.com/usefulco/oak-server/internal/provider"
)

// TODO:
// - Abstract new service scoffolding to module
// - update server methods to require preferred provider, reference from providers slice for each call
// - add proper response type

type IngestServer struct {
	UnimplementedIngestServiceServer
	providers map[provider.Provider]Provider
}

func (s *IngestServer) CreateIngest(ctx context.Context, r *CreateIngestRequest) (*Ingest, error) {
	result, err := s.providers[r.Provider].Create(&IngestCreateInput{
		Name:              r.Name,
		SourceName:        r.SourceName,
		PermittedSourceIP: r.SourceIpAddr,
	})
	if err != nil {
		log.Fatalf("failed to create aws setup: %v", err)
	}

	spew.Dump(result)

	return &Ingest{
		Id:   "1234",
		Name: "Sample",
	}, nil
}

func NewServer(awsSession *session.Session) IngestServiceServer {
	return &IngestServer{
		providers: map[provider.Provider]Provider{
			provider.Provider_AWS: NewAwsProvider(awsSession),
		},
	}
}

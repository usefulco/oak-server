package live

import (
	context "context"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/usefulco/oak-server/internal/provider"
)

type LiveServer struct {
	UnimplementedLiveServiceServer
	providers map[provider.Provider]Provider
}

func (s *LiveServer) CreateLive(ctx context.Context, r *CreateLiveRequest) (*Live, error) {
	_, err := s.providers[provider.Provider_AWS].Create(r)
	if err != nil {
		log.Fatalf("failed to create live: %v", err)
	}

	return &Live{}, nil
}

func NewServer(awsSession *session.Session) LiveServiceServer {
	return &LiveServer{
		providers: map[provider.Provider]Provider{
			provider.Provider_AWS: NewAwsProvider(awsSession),
		},
	}
}

package live

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/usefulco/oak-server/internal/provider"
)

type LiveServer struct {
	UnimplementedLiveServiceServer
	providers map[provider.Provider]Provider
}

func NewServer(awsSession *session.Session) LiveServiceServer {
	return &LiveServer{}
}

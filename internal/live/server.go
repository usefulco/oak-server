package live

import (
	"github.com/aws/aws-sdk-go/aws/session"
	pblive "github.com/usefulco/oak-server/api/live"
	pbprovider "github.com/usefulco/oak-server/api/provider"
)

type LiveServer struct {
	pblive.UnimplementedLiveServiceServer
	providers map[pbprovider.Provider]Provider
}

func NewServer(awsSession *session.Session) pblive.LiveServiceServer {
	return &LiveServer{}
}

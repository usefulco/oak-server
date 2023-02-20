package aws

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/usefulco/oak-server/pkg/aws_provider"
)

type AWSServer struct {
	UnimplementedAWSServiceServer
	client *aws_provider.AWSProvider
}

func NewAWSProviderServer() *AWSServer {
	session := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))

	c := aws_provider.NewAWSProvider(session)

	return &AWSServer{
		client: c,
	}
}

func (s *AWSServer) InitializeProvider(ctx context.Context, r *InitializeProviderInput) (*InitializeProviderOutput, error) {
	return &InitializeProviderOutput{}, nil
}

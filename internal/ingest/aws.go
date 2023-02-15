package ingest

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mediaconnect"
)

var providerName = "aws"

type awsProvider struct {
	client *mediaconnect.MediaConnect
}

func NewAwsProvider(session *session.Session) *awsProvider {
	return &awsProvider{
		client: mediaconnect.New(session),
	}
}

func (p *awsProvider) Create(i *ProviderCreateInput) (*ProviderCreateOutput, error) {
	options := mediaconnect.CreateFlowInput{
		Name: &i.Name,
		Source: &mediaconnect.SetSourceRequest{
			Name:          &i.SourceName,
			WhitelistCidr: aws.String(fmt.Sprintf("%s/32", i.PermittedSourceIP)),
			Protocol:      aws.String("srt-listener"),
		},
	}

	createFlowOutput, err := p.client.CreateFlow(&options)
	if err != nil {
		log.Fatalf("failed to create mediaconnect flow: %v", err)
	}

	r := &ProviderCreateOutput{
		ProviderName:      providerName,
		ProviderReference: *createFlowOutput.Flow.FlowArn,
		Location:          *createFlowOutput.Flow.AvailabilityZone,
		Name:              *createFlowOutput.Flow.Name,
		IngestIP:          *createFlowOutput.Flow.Source.IngestIp,
		IngestPort:        *createFlowOutput.Flow.Source.IngestPort,
		IngestProtocol:    *createFlowOutput.Flow.Source.Transport.Protocol,
		Status:            *createFlowOutput.Flow.Status,
	}

	return r, nil
}

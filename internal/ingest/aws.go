package ingest

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/mediaconnect"
	session "github.com/usefulco/oak-server/pkg/aws"
)

// TODO
// - move mediaConnect session outside of method handler

type AwsProvider struct{}

var providerName = "aws"

func (p *AwsProvider) Create(i *ProviderCreateInput) (*ProviderCreateOutput, error) {
	client := mediaconnect.New(session.SharedSession)

	options := mediaconnect.CreateFlowInput{
		Name: &i.Name,
		Source: &mediaconnect.SetSourceRequest{
			Name:          &i.SourceName,
			WhitelistCidr: aws.String(fmt.Sprintf("%s/32", i.PermittedSourceIP)),
			Protocol:      aws.String("srt-listener"),
		},
	}

	createFlowOutput, err := client.CreateFlow(&options)
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

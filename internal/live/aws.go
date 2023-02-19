package live

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/medialive"
	"github.com/davecgh/go-spew/spew"
)

var providerName = "aws"

// TODO: Need to generate S3 bucket via code

type awsProvider struct {
	client *medialive.MediaLive
}

func NewAwsProvider(session *session.Session) *awsProvider {
	return &awsProvider{
		client: medialive.New(session),
	}
}

func (p *awsProvider) Create(i *CreateLiveRequest) (*LiveCreateOutput, error) {
	// create MediaLive Input, not to be confused with Media Connect
	createInputOutput, err := p.client.CreateInput(&medialive.CreateInputInput{
		Name: aws.String("my-input"),
		Type: aws.String(medialive.InputTypeMediaconnect),
		MediaConnectFlows: []*medialive.MediaConnectFlowRequest{
			{
				FlowArn: aws.String(i.AwsConfig.MediaConnectArn),
			},
		},
		RoleArn: aws.String(i.AwsConfig.IamRoleArn),
	})
	if err != nil {
		panic(err)
	}

	// create MediaLiveChannel
	createChannelOutput, err := p.client.CreateChannel(&medialive.CreateChannelInput{
		ChannelClass: aws.String(medialive.ChannelClassSinglePipeline),
		Name:         aws.String("my-channel"),
		RoleArn:      aws.String(i.AwsConfig.IamRoleArn),
		InputAttachments: []*medialive.InputAttachment{
			{
				InputId: createInputOutput.Input.Id,
			},
		},
		Destinations: []*medialive.OutputDestination{
			{
				Settings: []*medialive.OutputDestinationSettings{
					{
						Url: aws.String("s3ssl://747795457281-oak-test-1"),
					},
				},
			},
		},
		EncoderSettings: &medialive.EncoderSettings{
			AudioDescriptions: []*medialive.AudioDescription{
				{},
			},
			OutputGroups: []*medialive.OutputGroup{
				{},
			},
			TimecodeConfig: &medialive.TimecodeConfig{},
			VideoDescriptions: []*medialive.VideoDescription{
				{},
			},
		},
	})
	if err != nil {
		panic(err)
	}

	spew.Dump(createChannelOutput)

	// create MediaLive output

	return nil, nil
}

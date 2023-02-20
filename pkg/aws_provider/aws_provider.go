package aws_provider

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/mediaconnect"
	"github.com/aws/aws-sdk-go/service/medialive"
)

type AWSProvider struct {
	session      *session.Session
	iam          *iam.IAM
	mediaconnect *mediaconnect.MediaConnect
	medialive    *medialive.MediaLive
}

func NewAWSProvider(s *session.Session) *AWSProvider {
	return &AWSProvider{
		session:      s,
		iam:          iam.New(s),
		mediaconnect: mediaconnect.New(s),
		medialive:    medialive.New(s),
	}
}

func (a *AWSProvider) CreateMediaLiveIAM(roleName string) (*string, error) {
	baseRole, err := a.iam.CreateRole(&iam.CreateRoleInput{
		RoleName: aws.String(roleName),
		AssumeRolePolicyDocument: aws.String(`
			"Version": "2012-10-17",
			"Statement": [{
				"Effect": "Allow",
				"Principal": {
					"Service": "medialive.amazonaws.com"
				},
				"Action": ["sts:AssumeRole"]
			}]
		`),
	})
	// TODO: ew
	if err != nil {
		return nil, err
	}

	// TODO: this can't be like this
	primaryPolicy, err := a.iam.CreatePolicy(&iam.CreatePolicyInput{
		PolicyName: aws.String(""),
		PolicyDocument: aws.String(`{
			"Version": "2012-10-17",
			"Statement": [{
				"Effect": "Allow",
				"Action": "*",
				"Resource": "*"
			}]
		}`),
	})
	// TODO: ew
	if err != nil {
		return nil, err
	}

	_, err = a.iam.AttachRolePolicy(&iam.AttachRolePolicyInput{
		RoleName:  baseRole.Role.RoleName,
		PolicyArn: primaryPolicy.Policy.Arn,
	})
	if err != nil {
		return nil, err
	}

	return baseRole.Role.Arn, nil
}

// TODO: struct this input?
func (a *AWSProvider) CreateMediaconnectInput(name string, sourceIp string, sourceName string) (*string, error) {
	flow, err := a.mediaconnect.CreateFlow(&mediaconnect.CreateFlowInput{
		Name: aws.String(name),
		Source: &mediaconnect.SetSourceRequest{
			Name:          aws.String(sourceName),
			WhitelistCidr: aws.String(fmt.Sprintf("%s/32", sourceIp)),
			Protocol:      aws.String("srt-listener"),
		},
	})
	// TODO: ew
	if err != nil {
		return nil, err
	}

	return flow.Flow.FlowArn, nil
}

// TODO: struct this input?
func (a *AWSProvider) CreateMediaLiveInput(name string, mediaconnectArn string, roleArn string) (*string, error) {
	r, err := a.medialive.CreateInput(&medialive.CreateInputInput{
		Name:    aws.String(name),
		Type:    aws.String(medialive.InputTypeMediaconnect),
		RoleArn: aws.String(roleArn),
		MediaConnectFlows: []*medialive.MediaConnectFlowRequest{
			{
				FlowArn: aws.String(mediaconnectArn),
			},
		},
	})
	// TODO: ew
	if err != nil {
		return nil, err
	}

	return r.Input.Arn, nil
}

func (a *AWSProvider) CreateMediaLiveChannel(name string, roleArn string, inputId string) (*string, error) {
	r, err := a.medialive.CreateChannel(&medialive.CreateChannelInput{
		ChannelClass: aws.String(medialive.ChannelClassSinglePipeline),
		Name:         aws.String(name),
		RoleArn:      aws.String(roleArn),
		InputAttachments: []*medialive.InputAttachment{
			{
				InputId: aws.String(inputId),
			},
		},
		Destinations: []*medialive.OutputDestination{
			{
				Settings: []*medialive.OutputDestinationSettings{
					{
						Url: aws.String("s3ssl://<bucket-name>"),
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
	// TODO: ew
	if err != nil {
		return nil, err
	}

	return r.Channel.Arn, nil
}

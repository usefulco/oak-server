package provider

import (
	context "context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

type ProviderServer struct {
	UnimplementedProviderServiceServer
	awsSession *session.Session
}

func (s *ProviderServer) SetupAWS(ctx context.Context, r *SetupAWSRequest) (*SetupAWSResponse, error) {
	// create iam for medialive
	iamSvc := iam.New(s.awsSession)

	iamResult, err := iamSvc.CreateRole(&iam.CreateRoleInput{
		RoleName: aws.String(fmt.Sprintf("%s-media-live", r.IamPrefix)),
		AssumeRolePolicyDocument: aws.String(`{
			"Version": "2012-10-17",
			"Statement": [{
				"Effect": "Allow",
				"Principal": {
					"Service": "medialive.amazonaws.com"
				},
				"Action": ["sts:AssumeRole"]
			}]
		}`),
	})
	if err != nil {
		panic(err)
	}

	policyOne, err := iamSvc.CreatePolicy(&iam.CreatePolicyInput{
		PolicyName: aws.String(fmt.Sprintf("%s-media-live", r.IamPrefix)),
		PolicyDocument: aws.String(`{
			"Version": "2012-10-17",
			"Statement": [{
				"Effect": "Allow",
				"Action": "*",
				"Resource": "*"
			}]
		}`),
	})
	if err != nil {
		panic(err)
	}

	_, err = iamSvc.AttachRolePolicy(&iam.AttachRolePolicyInput{
		RoleName:  iamResult.Role.RoleName,
		PolicyArn: policyOne.Policy.Arn,
	})
	if err != nil {
		panic(err)
	}

	return &SetupAWSResponse{
		MediaLiveIamArn: *iamResult.Role.Arn,
	}, nil
}

func NewServer(awsSession *session.Session) ProviderServiceServer {
	return &ProviderServer{
		awsSession: awsSession,
	}
}

package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var SharedSession *session.Session

func CreateSession() {
	SharedSession, _ = session.NewSession(&aws.Config{
		Region: aws.String("us-east-2"),
	})
}

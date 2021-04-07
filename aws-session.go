package awsS3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func (s *Session) ConnectAws(AWSSession Session) (*session.Session, error) {
	sessionAWS, err := session.NewSession(
		&aws.Config{
			Region: aws.String(AWSSession.AWSRegion),
			Credentials: credentials.NewStaticCredentials(
				AWSSession.AccessKeyID,
				AWSSession.SecretAccessKey,
				"",
			),
		})

	if err != nil {
		return nil, err
	}

	return sessionAWS, nil
}

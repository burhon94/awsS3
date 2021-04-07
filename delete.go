package awsS3

import (
	"errors"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (s *Session) DeleteFile(session *session.Session, file string) error {
	if session == nil || s == nil {
		err := errors.New("aws-Session is incorrect")
		return err
	}

	if len(strings.Split(file, ".")) < 2 && len(strings.Split(file, ".")) > 2 {
		err := errors.New("File incorrect, please inter file with mimeType")
		return err
	}

	svc := s3.New(session)
	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(file),
	})
	if err != nil {
		return err
	}

	err = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(file),
	})
	if err != nil {
		return err
	}

	return nil
}

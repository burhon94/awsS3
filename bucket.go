package awsS3

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (s *Session) CreateBucket(session *session.Session, bucketName string) error {
	if session == nil || s == nil {
		err := errors.New("aws-Session is incorrect")
		return err
	}

	svc := s3.New(session)
	_, err := svc.CreateBucket(&s3.CreateBucketInput{
		ACL:    aws.String("public-read"),
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return err
	}

	err = svc.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *Session) DeleteBucket(session *session.Session, bucketName string) error {
	if session == nil || s == nil {
		err := errors.New("aws-Session is incorrect")
		return err
	}

	svc := s3.New(session)
	_, err := svc.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return err
	}

	err = svc.WaitUntilBucketNotExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return err
	}

	return nil
}

package awsS3

import (
	"errors"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
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

func (s *Session) DeleteAllFiles(session *session.Session) error {
	if session == nil || s == nil {
		err := errors.New("aws-Session is incorrect")
		return err
	}

	svc := s3.New(session)
	listIterator := s3manager.NewDeleteListIterator(svc, &s3.ListObjectsInput{
		Bucket: aws.String(s.Bucket),
	})

	if err := s3manager.NewBatchDeleteWithClient(svc).Delete(aws.BackgroundContext(), listIterator); err != nil {
		return err
	}

	return nil
}

func (s *Session) DeleteBucketSvc(session *session.Session, bucketName string) error {
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

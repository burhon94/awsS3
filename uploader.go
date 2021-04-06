package awsS3

import (
	"errors"
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func (s *Session) UploadFile(awsSession *session.Session, file multipart.File, fileHeader *multipart.FileHeader, author string) (string, error) {
	var (
		filename = fileHeader.Filename
		mimeType = fileHeader.Header.Get("Content-Type")
		err      error
	)
	if awsSession == nil {
		err = errors.New("aws-Session is incorrect")
	}

	uploader := s3manager.NewUploader(awsSession)

	up, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(s.Bucket),
		ACL:         aws.String("public-read"),
		Key:         aws.String(filename),
		ContentType: aws.String(mimeType),
		Tagging:     aws.String(fmt.Sprintf("whoUpload=%v", author)),
		Body:        file,
	})

	if err != nil {
		err = errors.New(fmt.Sprintf("uploader: %v\nerr: %v", up, err))
		return "", err
	}
	filepath := up.Location

	return filepath, nil
}

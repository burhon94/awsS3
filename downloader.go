package awsS3

import (
	"errors"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func (s *Session) DownloadFile(sess *session.Session, file *os.File, contentType, item string) (int64, error) {
	if sess == nil || s == nil {
		err := errors.New("aws-Session is incorrect")
		return 0, err
	}

	downloader := s3manager.NewDownloader(sess)

	fileBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket:              aws.String(s.Bucket),
			Key:                 aws.String(item),
			ResponseContentType: aws.String(contentType),
		})
	if err != nil {
		return 0, err
	}

	return fileBytes, nil
}

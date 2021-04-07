package awsS3

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func (s *Session) UploadFile(awsSession *session.Session, file multipart.File, fileHeader *multipart.FileHeader, author string) (string, error) {
	var (
		filename = fileHeader.Filename
		mimeType = fileHeader.Header.Get("Content-Type")
		err      error
	)

	if awsSession == nil || s == nil {
		err = errors.New("aws-Session is incorrect")
		return "", err
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

func (s *Session) GetListFiles(session *session.Session) (listFiles ListFiles, err error) {
	var (
		file      Item
		itemsList []Item
		count     int
	)

	if session == nil || s == nil {
		err = errors.New("aws-Session is incorrect")
		return listFiles, err
	}

	svc := s3.New(session)
	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(s.Bucket)})
	if err != nil {
		return listFiles, err
	}

	for _, item := range resp.Contents {
		file.Name = *item.Key
		file.Size = *item.Size
		file.LastModify = *item.LastModified
		file.StorageClass = *item.StorageClass
		file.PathURL = "https://" + s.Bucket + ".s3." + s.AWSRegion + ".amazonaws.com/" + file.Name
		itemsList = append(itemsList, file)
		count++
	}
	listFiles.Items = itemsList
	listFiles.CountFiles = count

	return listFiles, nil
}

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

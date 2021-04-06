package awsS3

import (
	"errors"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Item struct {
	Name         string    `json:"name"`
	Size         int64     `json:"size"`
	LastModify   time.Time `json:"last_modify"`
	StorageClass string    `json:"storage_class"`
}

type ListFiles struct {
	CountFiles int    `json:"count_files"`
	Items      []Item `json:"items"`
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
		itemsList = append(itemsList, file)
		count++
	}
	listFiles.Items = itemsList
	listFiles.CountFiles = count

	return listFiles, nil
}

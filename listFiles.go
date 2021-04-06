package awsS3

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type item struct {
	Name         string
	Size         int64
	LastModify   time.Time
	StorageClass string
}

type ListFiles struct {
	CountFiles int
	Items []item
}

func (s *Session) GetListFiles(session *session.Session) (listFiles ListFiles, err error) {
	var (
	 file item
	 itemsList []item
	 count int
	)

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

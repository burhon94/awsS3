package awsS3

import (
	"time"
)

type Session struct {
	AccessKeyID     string
	SecretAccessKey string
	AWSRegion       string
	Bucket          string
}

type Item struct {
	Name         string    `json:"name"`
	Size         int64     `json:"size"`
	LastModify   time.Time `json:"last_modify"`
	StorageClass string    `json:"storage_class"`
	PathURL      string    `json:"path_url"`
}

type ListFiles struct {
	CountFiles int    `json:"count_files"`
	Items      []Item `json:"items"`
}

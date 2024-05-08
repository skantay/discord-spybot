package spybot

import (
	"bytes"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func newRepo(s3 *s3.S3) botRepo {
	return botRepo{s3}
}

type botRepo struct {
	s3 *s3.S3
}

func (b botRepo) uploadFile(bucket, key string, file []byte) error {
	_, err := b.s3.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(file),
	})
	return err
}

// getDownloadURL generates a presigned URL for downloading a file from S3.
func (b botRepo) getDownloadURL(bucket, key string, expiration time.Duration) (string, error) {
	req, _ := b.s3.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	urlStr, err := req.Presign(expiration)
	if err != nil {
		return "", err
	}
	return urlStr, nil
}

// listFiles lists all files in the specified S3 bucket.
func (b botRepo) listFiles(bucket string) ([]string, error) {
	result, err := b.s3.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return nil, err
	}

	var files []string
	for _, item := range result.Contents {
		files = append(files, *item.Key)
	}
	return files, nil
}

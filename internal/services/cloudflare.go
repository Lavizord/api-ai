package services

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type R2Uploader struct {
	client *s3.Client
	ve     *R2Venvs
}

func NewR2Uploader(ve *R2Venvs) (*R2Uploader, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(ve.AccessKeyId, ve.AccessKeySecret, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load r2 default configs: %v", err.Error())
	}
	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", ve.AccountId))
	})

	return &R2Uploader{
		client: client,
		ve:     ve,
	}, nil
}

func (u *R2Uploader) UploadFile(ctx context.Context, key string, data []byte) (string, error) {
	mimeType := http.DetectContentType(data)

	_, err := u.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:        aws.String(u.ve.BucketName),
		Key:           aws.String(key),
		Body:          bytes.NewReader(data),
		ContentLength: aws.Int64(int64(len(data))),
		ContentType:   aws.String(mimeType),
	})
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://%s.r2.cloudflarestorage.com/%s/%s", u.ve.AccountId, u.ve.BucketName, key)
	return url, nil
}

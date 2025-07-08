package services

import (
	"fmt"
	"os"
)

type R2Venvs struct {
	AccountId       string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
}

func LoadR2Venv() (*R2Venvs, error) {
	env := &R2Venvs{
		AccountId:       os.Getenv("R2_ACCOUNT_ID"),
		AccessKeyId:     os.Getenv("R2_ACCESS_KEY_ID"),
		AccessKeySecret: os.Getenv("R2_ACCESS_KEY_SECRET"),
		BucketName:      os.Getenv("R2_BUCKET_NAME"),
	}

	var missing []string
	if env.AccountId == "" {
		missing = append(missing, "R2_ACCOUNT_ID")
	}
	if env.AccessKeyId == "" {
		missing = append(missing, "R2_ACCESS_KEY_ID")
	}
	if env.AccessKeySecret == "" {
		missing = append(missing, "R2_ACCESS_KEY_SECRET")
	}
	if env.BucketName == "" {
		missing = append(missing, "R2_BUCKET_NAME")
	}

	if len(missing) > 0 {
		return nil, fmt.Errorf("missing required environment variables for Cloudflare  r2 Service: %v", missing)
	}

	return env, nil
}

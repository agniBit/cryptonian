package s3

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/agniBit/cryptonian/internal/logger"
	"github.com/agniBit/cryptonian/model/cfg"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var initOnce = &sync.Once{}
var client *s3.Client

func InitS3(cfg *cfg.Config) {
	initOnce.Do(func() {
		awsCfg, err := config.LoadDefaultConfig(context.TODO(),
			config.WithRegion(cfg.S3.AwsRegion),
			config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
				Value: aws.Credentials{
					AccessKeyID:     cfg.S3.AwsAccessKeyId,
					SecretAccessKey: cfg.S3.AwsSecretAccessKey,
					Source:          "credentials from environment variables",
				},
			}),
		)

		if err != nil {
			logger.Fatal(context.Background(), "unable to load AWS config", err, nil)
		}
		client = s3.NewFromConfig(awsCfg)
	})
}

func UploadToS3(ctx context.Context, bucket, key, filePath string) error {

	file, err := os.Open(filePath)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("unable to open file %s: %v", filePath, err), nil)
		return fmt.Errorf("unable to open file %s: %v", filePath, err)
	}
	defer file.Close()

	_, err = client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("unable to upload %s to S3: %v", filePath, err), nil)
		return fmt.Errorf("unable to upload %s to S3: %v", filePath, err)
	}

	return nil
}

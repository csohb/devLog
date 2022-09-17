package image

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type AWSConfig struct {
	AccessKeyID     string
	AccessKeySecret string
	Region          string
	BucketName      string
	UploadTimeout   int
	BaseURL         string
}

func NewAWSConfig(accessKey, accessSecret, region, bucketName, baseUrl string, timeout int) *AWSConfig {
	return &AWSConfig{
		AccessKeyID:     accessKey,
		AccessKeySecret: accessSecret,
		Region:          region,
		BucketName:      bucketName,
		UploadTimeout:   timeout,
		BaseURL:         baseUrl,
	}
}

func CreateSession(awsConfig AWSConfig) *s3.S3 {
	sess := session.Must(session.NewSession(
		&aws.Config{
			Region: aws.String(awsConfig.Region),
			Credentials: credentials.NewStaticCredentials(
				awsConfig.AccessKeyID,
				awsConfig.AccessKeySecret,
				"",
			),
		},
	))
	s3Session := s3.New(sess)
	return s3Session
}

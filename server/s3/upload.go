package s3

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/sirupsen/logrus"
	"io"
	"strings"
	"time"
)

type Uploader struct {
	Bucket     string
	Region     string
	Profile    string
	Cfg        *aws.Config
	Logger     *logrus.Logger
	client     *s3.Client
	uploader   *manager.Uploader
	downloader *manager.Downloader
}

func NewS3Uploader(bucket, cfgFile, credential, profile string) *Uploader {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedCredentialsFiles([]string{credential}),
		config.WithSharedConfigFiles([]string{cfgFile}),
		config.WithSharedConfigProfile(profile),
		config.WithLogConfigurationWarnings(true),
	)
	if err != nil {
		logrus.Fatalf("failed to load configuration, %s", err)
		return nil
	}

	fmt.Println("cfg : ", cfg)

	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)
	downloader := manager.NewDownloader(client)
	return &Uploader{
		Bucket:     bucket,
		Cfg:        &cfg,
		Logger:     logrus.StandardLogger(),
		client:     client,
		uploader:   uploader,
		downloader: downloader,
		Region:     "ap-northeast-2",
	}
}

func (u *Uploader) GetBucketList() (ret []string, err error) {
	output, err := u.client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		return nil, err
	}

	for _, v := range output.Buckets {
		ret = append(ret, *v.Name)
	}

	return ret, nil
}

func (u *Uploader) UploadFile(dirName, fileName, contentType string, file io.Reader, useOrgName bool) (url string, err error) {
	if useOrgName == false {
		tmp := strings.Split(fileName, ".")
		ext := "." + tmp[len(tmp)-1]
		fileName = u.GenerateFileName(fileName) + ext
	}

	path := fmt.Sprintf("%s/%s/%s", dirName, time.Now().Format("20060102"), fileName)
	input := &s3.PutObjectInput{
		Bucket:      &u.Bucket,
		Key:         &path,
		Body:        file,
		ContentType: &contentType,
	}

	output, err := u.uploader.Upload(context.TODO(), input)
	if err != nil {
		return "", err
	}
	url = output.Location
	return url, nil
}

func (u Uploader) GenerateFileName(fileName string) string {
	src := fileName + time.Now().Format("150405")
	h := md5.New()
	defer h.Reset()
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil))
}

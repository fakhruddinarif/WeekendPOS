package service

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"mime/multipart"
)

func UploadImage(entity string, s3Client *s3.Client, fileHeader *multipart.FileHeader) (*string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buffer := bytes.NewBuffer(nil)
	if _, err := buffer.ReadFrom(file); err != nil {
		return nil, err
	}

	key := fmt.Sprintf("%s/%s-%s", entity, uuid.New().String(), fileHeader.Filename)

	_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(viper.GetString("s3.bucket")),
		Key:    aws.String(key),
		Body:   bytes.NewReader(buffer.Bytes()),
	})
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/%s", viper.GetString("s3.url"), key)
	return &url, nil
}

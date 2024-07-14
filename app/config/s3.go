package config

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/viper"
)

func NewS3Client(viper *viper.Viper) *s3.Client {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if service == s3.ServiceID {
			return aws.Endpoint{
				URL:               viper.GetString("s3.url"),
				SigningRegion:     viper.GetString("s3.region"),
				HostnameImmutable: true,
			}, nil
		}
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(viper.GetString("s3.region")),
		config.WithCredentialsProvider(aws.NewCredentialsCache(aws.CredentialsProviderFunc(func(context.Context) (aws.Credentials, error) {
			return aws.Credentials{
				AccessKeyID:     viper.GetString("s3.access-key-id"),
				SecretAccessKey: viper.GetString("s3.secret-access-key"),
			}, nil
		}))),
		config.WithEndpointResolverWithOptions(customResolver),
	)
	if err != nil {
		panic("Unable to load AWS SDK config, " + err.Error())
	}
	client := s3.NewFromConfig(cfg)

	return client
}

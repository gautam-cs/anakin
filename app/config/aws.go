package config

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/smithy-go/logging"
	"github.com/phuslu/log"
)

func awsConfig(key string, secret string, region string) (aws.Config, error) {
	clientLogMode := aws.LogRetries | aws.LogRequest | aws.LogResponse

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(key, secret, "")),
		config.WithLogConfigurationWarnings(true),
		config.WithClientLogMode(clientLogMode),
		config.WithLogger(logging.NewStandardLogger(os.Stdout)))

	if err != nil {
		log.Error().Err(err).Msg("error loading aws config")
	}

	return cfg, err
}

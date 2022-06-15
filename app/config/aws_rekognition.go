package config

import (
	"sync"

	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"go.mongodb.org/mongo-driver/mongo"
)

//RekognitionClient handle
func rekognitionConfig() *awsRekognitionConfig {
	if IsDebugEnv() {
		return appConfig.AWSRekognition
	}

	if appConfig.AWSRekognition != nil {
		return appConfig.AWSRekognition
	}

	appConfig.AWSRekognition = new(awsRekognitionConfig)
	readDecodedSecret("rekognition", appConfig.AWSRekognition)
	return appConfig.AWSRekognition
}

var rekognitionClients = map[string]*rekognition.Client{}
var rekogMutex = &sync.Mutex{}
var docDBMutex = &sync.Mutex{}
var documentDB *mongo.Database

func DefaultRekognitionClient() *rekognition.Client {
	return MakeRekognitionClient(*appConfig.Server.AwsRegion)
}

func MakeRekognitionClient(region string) *rekognition.Client {
	rekogMutex.Lock()

	defer func() {
		rekogMutex.Unlock()
	}()

	if client, ok := rekognitionClients[region]; ok {
		return client
	}

	config := rekognitionConfig()

	rekognitionAWSConfig, _ := awsConfig(
		config.AccessKey,
		config.AccessSecret,
		region)

	client := rekognition.NewFromConfig(rekognitionAWSConfig)

	rekognitionClients[region] = client

	return client
}

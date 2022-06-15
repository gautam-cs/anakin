package config

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/spf13/viper"
)

//MySQLConfig holder
type mySqlConfig struct {
	Address        string `mapstructure:"address"`
	Username       string `mapstructure:"username"`
	Password       string `mapstructure:"password"`
	Database       string `mapstructure:"database"`
	MaxConnections int    `mapstructure:"maxConnections"`
}

type awsRekognitionConfig struct {
	AccessKey    string `mapstructure:"accessKey"`
	AccessSecret string `mapstructure:"accessSecret"`
	Region       string `mapstructure:"region"`
}

type awsDocumentDBConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Address  string `mapstructure:"address"`
	Database string `mapstructure:"database"`
}

type redisConfig struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
}

type awsS3Config struct {
	AccessKey    string `mapstructure:"accessKey"`
	AccessSecret string `mapstructure:"accessSecret"`
	Bucket       string `mapstructure:"bucket"`
	Region       string `mapstructure:"region"`
}

type awsSESConfig struct {
	AccessKey    string `mapstructure:"accessKey"`
	AccessSecret string `mapstructure:"accessSecret"`
	Region       string `mapstructure:"region"`
	SenderName   string `mapstructure:"senderName"`
	SenderEmail  string `mapstructure:"senderEmail"`
}

func (config *awsSESConfig) Sender() string {
	return fmt.Sprintf("%s <%s>", config.SenderName, config.SenderEmail)
}

//ServerConfig holder
type serverConfig struct {
	Port          int     `mapstructure:"port"`
	AwsRegion     *string `mapstructure:"awsRegion"`
	TemplatesPath string  `mapstructure:"templatesPath"`
}

//AccountConfig holder
type accountConfig struct {
	JwtSecret string `mapstructure:"jwtSecret"`
	CryptKey  string `mapstructure:"cryptKey"`
}

//CommonConfig server common settings
type commonConfig struct {
	// Domain      string `mapstructure:"domain"`
	Environment string `mapstructure:"environment"`
}

type zoomConfig struct {
	ApiKey    string `mapstructure:"apiKey"`
	ApiSecret string `mapstructure:"apiSecret"`
}

type twilioConfig struct {
	AccountSID   string `mapstructure:"accountSID"`
	AuthToken    string `mapstructure:"authToken"`
	FromPhone    string `mapstructure:"fromPhone"`
	SMSFromPhone string `mapstructure:"smsFromNumber"`
}

type kafkaConfig struct {
	BootstrapServers string `mapstructure:"bootstrap.servers"`
}

type mytConfig struct {
	Server         *serverConfig         `mapstructure:"server"`
	Account        *accountConfig        `mapstructure:"account"`
	ReadMySQL      *mySqlConfig          `mapstructure:"readsqldb"`
	WriteMySQL     *mySqlConfig          `mapstructure:"writesqldb"`
	Redis          *redisConfig          `mapstructure:"redis"`
	AWSRekognition *awsRekognitionConfig `mapstructure:"rekognition"`
	AWSDocumentDB  *awsDocumentDBConfig  `mapstructure:"documentdb"`
	AWSS3          *awsS3Config          `mapstructure:"s3"`
	AWSSES         *awsSESConfig         `mapstructure:"ses"`
	Zoom           *zoomConfig           `mapstructure:"zoom"`
	Twilio         *twilioConfig         `mapstructure:"twilio"`
	Kafka          *kafkaConfig          `mapstructure:"kafka"`
}

var appConfig *mytConfig
var common *commonConfig

//IsDebugEnv env check
func IsDebugEnv() bool {
	return EnvironmentType(common.Environment) == EnvironmentTypeDebug
}

func IsProdEnv() bool {
	return EnvironmentType(common.Environment) == EnvironmentTypeProduction
}

func configLoadFailed(err error) {
	fmt.Println(err)
	panic(err)
}

//LoadConfig from file
func Load() {
	appConfig = &mytConfig{}

	if err := readFile("env"); err != nil {
		configLoadFailed(err)
	}

	common = &commonConfig{}
	if err := viper.Unmarshal(common); err != nil {
		configLoadFailed(err)
	}

	if err := readFile(common.Environment); err != nil {
		configLoadFailed(err)
	}

	if err := viper.Unmarshal(appConfig); err != nil {
		configLoadFailed(err)
	}

	fmt.Println("config:", *appConfig.Server.AwsRegion)
}

func readFile(name string) error {

	viper.SetConfigName(name)             // name of config file (without extension)
	viper.SetConfigType("json")           // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("config/secrets") // optionally look for config in the project directory

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Error reading config: missing env.json")
		} else {
			fmt.Println("Error reading config", err)
		}

		return err
	}

	return nil
}

func readSecret(configKey string) (*secretsmanager.GetSecretValueOutput, error) {
	//Create a Secrets Manager client
	awsConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(*appConfig.Server.AwsRegion))

	if err != nil {
		return nil, err
	}

	svc := secretsmanager.NewFromConfig(awsConfig)

	secretKey := fmt.Sprintf("%s/%s", common.Environment, configKey)

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretKey),
	}

	return svc.GetSecretValue(context.TODO(), input)
}

func readDecodedSecret(configKey string, v interface{}) {
	output, err := readSecret(configKey)

	if err != nil {
		configLoadFailed(err)
	}

	if err := json.Unmarshal([]byte(*output.SecretString), v); err != nil {
		configLoadFailed(err)
	}
}

func ZoomConfig() *zoomConfig {
	if IsDebugEnv() {
		return appConfig.Zoom
	}

	if appConfig.Zoom != nil {
		return appConfig.Zoom
	}

	appConfig.Zoom = new(zoomConfig)
	readDecodedSecret("zoom", appConfig.Zoom)
	return appConfig.Zoom
}

func TwilioConfig() *twilioConfig {
	if IsDebugEnv() {
		return appConfig.Twilio
	}

	if appConfig.Twilio != nil {
		return appConfig.Twilio
	}

	appConfig.Twilio = new(twilioConfig)
	readDecodedSecret("twilio", appConfig.Twilio)
	return appConfig.Twilio
}

func ListenPort() int {
	if appConfig.Server != nil {
		return appConfig.Server.Port
	}

	return 80
}

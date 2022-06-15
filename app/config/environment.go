package config

type EnvironmentType string

const (
	EnvironmentTypeDebug      EnvironmentType = "debug"
	EnvironmentTypePreProd    EnvironmentType = "preprod"
	EnvironmentTypeProduction EnvironmentType = "prod"
)

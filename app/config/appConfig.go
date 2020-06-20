// Package config reasd configurations from a YAML file and load them into a AppConfig type to save the configuration
// information for the application.
// Configuration for different environment can be saved in files with different suffix, for example [Dev], [Prod]
package config

import (
	"fmt"
	logConfig "github.com/jfeng45/glogger/config"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// AppConfig represents the application config
type AppConfig struct {
	SQLConfig       DataStoreConfig `yaml:"sqlConfig"`
	SQLConfigTx       DataStoreConfig `yaml:"sqlConfigTx"`
	CouchdbConfig   DataStoreConfig `yaml:"couchdbConfig"`
	CacheGrpcConfig DataStoreConfig `yaml:"cacheGrpcConfig"`
	UserGrpcConfig  DataStoreConfig `yaml:"userGrpcConfig"`
	ZapConfig       logConfig.Logging       `yaml:"zapConfig"`
	LorusConfig     logConfig.Logging        `yaml:"logrusConfig"`
	LogConfig       logConfig.Logging        `yaml:"logConfig"`
	UseCaseConfig   UseCaseConfig   `yaml:"useCaseConfig"`
}

// UseCaseConfig represents different use cases
type UseCaseConfig struct {
	Registration RegistrationConfig `yaml:"registration"`
	RegistrationTx RegistrationTxConfig `yaml:"registrationTx"`
	ListUser     ListUserConfig     `yaml:"listUser"`
}

// RegistrationConfig represents registration use case
type RegistrationConfig struct {
	Code           string     `yaml:"code"`
	UserDataConfig DataConfig `yaml:"userDataConfig"`
}

// RegistrationConfigTx represents registration use cases that support transaction
type RegistrationTxConfig struct {
	Code           string     `yaml:"code"`
	UserDataConfig DataConfig `yaml:"userDataConfig"`
}

// ListUserConfig represents list user use case
type ListUserConfig struct {
	Code            string     `yaml:"code"`
	UserDataConfig  DataConfig `yaml:"userDataConfig"`
	CacheDataConfig DataConfig `yaml:"cacheDataConfig"`
}

// DataConfig represents data service
type DataConfig struct {
	Code            string          `yaml:"code"`
	DataStoreConfig DataStoreConfig `yaml:"dataStoreConfig"`
}

// DataConfig represents handlers for data store. It can be a database or a gRPC connection
type DataStoreConfig struct {
	Code string `yaml:"code"`
	// Only database has a driver name, for grpc it is "tcp" ( network) for server
	DriverName string `yaml:"driverName"`
	// For database, this is datasource name; for grpc, it is target url
	UrlAddress string `yaml:"urlAddress"`
	// Only some databases need this database name
	DbName string `yaml:"dbName"`
	// To indicate whether support transaction or not. "true" means supporting transaction
	Tx bool `yaml:"tx"`
}

// LogConfig represents logger handler
// Logger has many parameters can be set or changed. Currently, only three are listed here. Can add more into it to
// fits your needs.
type LogConfig struct {
	// log library name
	Code string `yaml:"code"`
	// log level
	Level string `yaml:"level"`
	// show caller in log message
	EnableCaller bool `yaml:"enableCaller"`
}

// BuildConfig build the AppConfig
// if the filaname is not empty, then it reads the file of the filename (in the same folder) and put it into the AppConfig
func BuildConfig(filename ...string) (*AppConfig, error) {
	if len(filename) == 1 {
		return buildConfigFromFile(filename[0])
	} else {
		return BuildConfigWithoutFile()
	}
}

// BuildConfigWithoutFile create AppConfig with adhoc value
func BuildConfigWithoutFile() (*AppConfig, error) {
	return nil, nil
}

// buildConfigFromFile reads the file of the filename (in the same folder) and put it into the AppConfig
func buildConfigFromFile(filename string) (*AppConfig, error) {

	var ac AppConfig
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrap(err, "read error")
	}
	err = yaml.Unmarshal(file, &ac)

	if err != nil {
		return nil, errors.Wrap(err, "unmarshal")
	}
	err = validateConfig(ac)
	if err != nil {
		return nil, errors.Wrap(err, "validate config")
	}
	fmt.Println("appConfig:", ac)
	return &ac, nil
}



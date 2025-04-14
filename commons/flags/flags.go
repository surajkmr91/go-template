package flags

import (
	"os"

	"trendtracker/constants"

	flag "github.com/spf13/pflag"
)

var (
	port           = flag.Int(constants.PortKey, constants.PortDefaultValue, "port")
	baseConfigPath = flag.String(constants.BaseConfigPathKey, constants.BaseConfigPathDefaultValue,
		"path to folder that stores the configuration")
)

func init() {
	flag.Parse()
}

// Env is the application.yml runtime environment
func Env() string {
	env := os.Getenv(constants.EnvKey)
	if env == "" {
		return constants.EnvDefaultValue
	}
	return env
}

// Port is the application.yml port number where the process will be started
func Port() int {
	return *port
}

// BaseConfigPath is the path that holds the configuration files
func BaseConfigPath() string {
	return *baseConfigPath
}

// AWSRegion is the region where the application is running
func AWSRegion() string {
	region := os.Getenv(constants.AWSRegionKey)
	if region == "" {
		return constants.AWSRegionDefaultValue
	}
	return region
}

// AWSAccessKeyID is the access key id for aws
func AWSAccessKeyID() string {
	return os.Getenv(constants.AWSAccessKeyID)
}

// AWSSecretAccessKey is the secret access key for aws
func AWSSecretAccessKey() string {
	return os.Getenv(constants.AWSSecretAccessKey)
}

// AWSBucket is used to get the aws bucket
func AWSBucket() string {
	return os.Getenv(constants.AWSBucket)
}

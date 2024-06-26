package config

import (
	"fmt"
	"os"
)

type appConfig struct {
	Namespace   string
	ServiceName string
}

var globalAppConfig = appConfig{}

func init() {
	globalAppConfig.Namespace = os.Getenv("NAMESPACE")
	globalAppConfig.ServiceName = os.Getenv("SERVICE")

	if globalAppConfig.Namespace == "" {
		panic(fmt.Sprintln("init app config failed. missing namespace in env."))
	}
	if globalAppConfig.ServiceName == "" {
		panic(fmt.Sprintln("init app config failed. missing namespace in env."))
	}
}

func GetNamespace() string {
	return globalAppConfig.Namespace
}

func GetServiceName() string {
	return globalAppConfig.ServiceName
}

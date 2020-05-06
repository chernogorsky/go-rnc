package rnc

import (
	"fmt"
	"os"
)

const envParamStoreType string = "RNC_PARAMS_STORE"
const defEnvParamStoreType string = "env"

func getRemoteConfigEnv(paramName string) (string, error) {
	return os.Getenv(paramName), nil
}
func getRemoteConfig(paramName string) (string, error) {
	paramStoreType := os.Getenv(envParamStoreType)
	if paramStoreType == "" {
		paramStoreType = defEnvParamStoreType
	}

	switch paramStoreType {
	case "env":
		return getRemoteConfigEnv(paramName)
	default:
		return "", fmt.Errorf("getRemoteConfig: Method %s is not supported", paramStoreType)
	}
}

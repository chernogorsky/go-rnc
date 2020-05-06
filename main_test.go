package main

import (
	"os"
	"testing"
)

func TestGetRemoteConfig(t *testing.T) {
	_, err := getRemoteConfig("param_name")
	if err != nil {
		t.Errorf("getRemoteConfig returns error %v", err)
	}
}

func TestGetRemoteConfigEnv(t *testing.T) {
	paramName := "RNC_SOME_PARAM"
	paramVal := "someValue"
	if err := os.Setenv(paramName, paramVal); err != nil {
		t.Fatalf("Setenv: %v", err)
	}
	res, err := getRemoteConfig(paramName)
	if err != nil {
		t.Errorf("getRemoteConfig returns error %v", err)
	}
	if res != paramVal {
		t.Errorf("getRemoteConfig.env returns wrong value: %s instead of %s", res, paramVal)
	}
}

func TestGetRemoteConfigError(t *testing.T) {
	if err := os.Setenv(envParamStoreType, "none"); err != nil {
		t.Fatalf("Setenv: %v", err)
	}
	_, err := getRemoteConfig("param_name")
	if err == nil {
		t.Errorf("getRemoteConfig does not returns error")
	}
}

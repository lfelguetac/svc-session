package utils_test

import (
	"session-service-v2/app/utils"
	"testing"
)

func TestGetEnvStringWithoudDefined(t *testing.T) {
	envValue := utils.GetEnvString("F001", "defaultValue")
	if envValue != "defaultValue" {
		t.Errorf("GetEnvString(\"F001\", \"defaultValue\") FAILED. Expected %s, got %s\n", "defaultValue", envValue)
	} else {
		t.Logf("GetEnvString(\"F001\", \"defaultValue\") PASSED. Expected %s, got %s\n", "defaultValue", envValue)
	}
}

func TestGetEnvStringDefined(t *testing.T) {
	t.Setenv("F002", "defined")

	envValue := utils.GetEnvString("F002", "default")
	if envValue != "defined" {
		t.Errorf("GetEnvString(\"F002\", \"default\") FAILED. Expected %s, got %s\n", "defined", envValue)
	} else {
		t.Logf("GetEnvString(\"F002\", \"default\") PASSED. Expected %s, got %s\n", "defined", envValue)
	}
}

func TestGetBoolEnvWithoutDefined(t *testing.T) {
	envValue := utils.GetBoolEnv("F003", true)

	if envValue != true {
		t.Errorf("GetBoolEnv(\"F003\", true) FAILED. Expected %t, got %t\n", true, envValue)
	} else {
		t.Logf("GetBoolEnv(\"F003\", true) PASSED. Expected %t, got %t\n", true, envValue)

	}
}

func TestGetBoolEnvDefined(t *testing.T) {
	t.Setenv("F004", "false")

	envValue := utils.GetBoolEnv("F004", true)

	if envValue != false {
		t.Errorf("GetBoolEnv(\"F004\", true) FAILED. Expected %t, got %t\n", false, envValue)
	} else {
		t.Logf("GetBoolEnv(\"F004\", true) PASSED. Expected %t, got %t\n", false, envValue)
	}
}

func TestGetBoolEnvDefinedButWrong(t *testing.T) {
	t.Setenv("F005", "prueba")

	envValue := utils.GetBoolEnv("F005", true)

	if envValue != true {
		t.Errorf("GetBoolEnv(\"F005\", true) FAILED. Expected %t, got %t\n", true, envValue)
	} else {
		t.Logf("GetBoolEnv(\"F005\", true) PASSED. Expected %t, got %t\n", true, envValue)
	}
}

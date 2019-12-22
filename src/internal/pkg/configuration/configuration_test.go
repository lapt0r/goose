package configuration

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var testBytes = []byte(`
	[
		{
			"Name" : "TestRule",
			"Rule" : "password",
			"Confidence" : 0.7,
			"Severity" : 0
		},
		{
			"Name" : "LowConfidenceRule",
			"Rule" : "api",
			"Confidence" : 0.3,
			"Severity" : 5
		}
	]
	`)

func TestFileLoad(t *testing.T) {
	file, err := ioutil.TempFile("", "TestFileLoad")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())
	file.Write(testBytes)
	result := LoadConfiguration(file.Name())
	validationerror := validateConfiguration(result)
	if validationerror != nil {
		t.Error(validationerror)
	}
}

func TestConfigurationLoad(t *testing.T) {
	result := unmarshalConfiguration(testBytes)
	err := validateConfiguration(result)
	if err != nil {
		t.Error(err)
	}
}

func validateConfiguration(config []ScanRule) error {
	if len(config) != 2 {
		return fmt.Errorf("Expected to deserialize 2 rules but found %v", len(config))
	}
	if config[0].Name != "TestRule" {
		return fmt.Errorf("Expected Name to be TestRule but was %v", config[0].Name)
	}
	if config[1].Name != "LowConfidenceRule" {
		return fmt.Errorf("expected Name to be LowConfidenceRule but was %v", config[1].Name)
	}
	return nil
}

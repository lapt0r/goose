package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//ScanRule contains a friendly name, regex rule, confidence, and severity
type ScanRule struct {
	Name       string
	Rule       string
	Confidence float64
	Severity   int
}

//LoadConfiguration loads a set of ScanRules from a provided path target
func LoadConfiguration(path string) []ScanRule {

	var contents, err = ioutil.ReadFile(path)
	if err != nil {
		panic(err) //per Go documentation, io.ReadFile returns nil if successfully read to end of file, not err == EOF
	}
	return unmarshalConfiguration(contents)
}

func unmarshalConfiguration(b []byte) []ScanRule {
	var result []ScanRule
	err := json.Unmarshal(b, &result)
	if err != nil {
		fmt.Println("Error loading configuration: ", err)
	}
	return result
}

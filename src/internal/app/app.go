package app

import (
	"fmt"
	"internal/pkg/configuration"
	"internal/pkg/loader"
	"internal/pkg/regexfilter"
	"path/filepath"
)

var rules []configuration.ScanRule
var targets []loader.ScanTarget

//RuleCount : Get count of rules initialized by application.
func RuleCount() int {
	return len(rules)
}

//Init : initalizes the Goose application
func Init(configpath string, targetpath string) {
	//todo
	if configpath == "" {
		configpath = "config/goose_rules.json"
	}
	rules = configuration.LoadConfiguration(configpath)
	absoluteTargetPath, _ := filepath.Abs(targetpath)
	fmt.Printf("Initializing Goose with target [%v]..\n", absoluteTargetPath)
	targets = loader.GetTargets(absoluteTargetPath)
	fmt.Printf("\nGot [%v] targets\n", len(targets))
}

//Run : runs the Goose application
func Run() {
	var result []regexfilter.Finding
	for _, target := range targets {
		for _, rule := range rules {
			ruleChannel := make(chan regexfilter.Finding)
			go regexfilter.ScanFile(target, rule, ruleChannel)
			for f := range ruleChannel {
				//kb todo: application config for confidence threshold
				if !f.IsEmpty() && f.Confidence > 0.65 {
					result = append(result, f)
				}
			}
		}
	}
	fmt.Println("\n--- Scanning complete ---")
	fmt.Printf("[%v] results\n", len(result))
	for _, finding := range result {
		fmt.Printf("FINDING\n-----\n%v\n", finding)
	}
}

package app

import (
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"
	"sync"

	"github.com/lapt0r/goose/internal/pkg/configuration"
	"github.com/lapt0r/goose/internal/pkg/loader"
	"github.com/lapt0r/goose/internal/pkg/regexfilter"
)

var rules []configuration.ScanRule
var targets []loader.ScanTarget

//RuleCount : Get count of rules initialized by application.
func RuleCount() int {
	return len(rules)
}

//Init : initalizes the Goose application
func Init(configpath string, targetpath string, interactive bool) {
	if configpath == "" {
		configpath = "config/goose_rules.json"
	}
	//todo : application configuration tuning knobs (confidence intervals, etc)
	rules = configuration.LoadConfiguration(configpath)
	absoluteTargetPath, _ := filepath.Abs(targetpath)
	if interactive {
		fmt.Printf("Initializing Goose with target [%v]..\n", absoluteTargetPath)
	}
	targets = loader.GetTargets(absoluteTargetPath)
	if interactive {
		fmt.Printf("\nGot [%v] targets\n", len(targets))
	}
}

//Run : runs the Goose application
func Run(interactive bool) {
	var result []regexfilter.Finding
	var ruleChannel = make(chan regexfilter.Finding, 4)
	var bufferChannel = make(chan bool)
	var wg sync.WaitGroup
	//anonymous buffer thread to empty the channel to prevent deadlocking
	go func() {
		for {
			select {
			case <-bufferChannel:
				break
			case f, ok := <-ruleChannel:
				if ok && f.Confidence > 0.65 {
					result = append(result, f)
				}
			default:
				//do nothing
			}
		}
	}()
	for _, rule := range rules {
		for _, target := range targets {
			wg.Add(1)
			go regexfilter.ScanFile(target, rule, ruleChannel, &wg)
		}
		if interactive {
			fmt.Printf("Waiting for routines to finish..\n")
		}
		wg.Wait()
		bufferChannel <- false
	}

	outputResults(result, interactive)
}

func outputResults(result []regexfilter.Finding, interactive bool) {
	if interactive {
		fmt.Println("\n--- Scanning complete ---")
		fmt.Printf("[%v] results\n", len(result))
		for _, finding := range result {
			fmt.Printf("FINDING\n-----\n%v\n", finding)
		}
	} else {
		json, err := json.Marshal(result)
		if err != nil {
			//todo : should this terminate?
			log.Fatal(err)
		}
		//todo(?) : support file output target or rely on caller piping to stream?
		fmt.Printf("%v", string(json))
	}
}

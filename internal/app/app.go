package app

import (
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"sync"

	"github.com/lapt0r/goose/internal/pkg/configuration"
	"github.com/lapt0r/goose/internal/pkg/decisionfilter"
	"github.com/lapt0r/goose/internal/pkg/finding"
	"github.com/lapt0r/goose/internal/pkg/loader"
	"github.com/lapt0r/goose/internal/pkg/regexfilter"
	"github.com/lapt0r/goose/internal/pkg/report"
)

var rules []configuration.ScanRule
var targets []loader.ScanTarget

//RuleCount : Get count of rules initialized by application.
func RuleCount() int {
	return len(rules)
}

//Init : initalizes the Goose application
func Init(configpath string, targetpath string, interactive bool, commitDepth int) {
	if configpath == "" {
		configpath = "config/goose_rules.json"
	}
	//todo : application configuration tuning knobs (confidence intervals, etc)
	rules = configuration.LoadConfiguration(configpath)
	absoluteTargetPath, _ := filepath.Abs(targetpath)
	if interactive {
		fmt.Printf("Initializing Goose with target [%v]..\n", absoluteTargetPath)
	}
	targets = loader.GetTargets(absoluteTargetPath, commitDepth)
	if interactive {
		fmt.Printf("Got [%v] targets\n", len(targets))
	}
}

//Run : runs the Goose application
func Run(interactive bool, decisionTree bool, outputmode string) {
	var fChannel = make(chan []finding.Finding, 1)

	var results []finding.Finding
	var wg sync.WaitGroup
	var total = len(targets)
	var count = 0
	for _, target := range targets {
		wg.Add(1)
		if decisionTree {
			go decisionfilter.ScanFile(target, fChannel, &wg)
		} else {
			go regexfilter.ScanFile(target, &rules, fChannel, &wg)
		}
	}

	go func() {
		defer func() {
			if r := recover(); r != nil {
				println("panic:" + r.(string))
			}
		}()
		if interactive {
			fmt.Printf("Waiting for routines to finish..\n")
		}
		for {
			result, _ := <-fChannel
			//empty results are only returned by finished scan goroutines
			if len(result) == 0 {
				count++
			} else {
				results = append(results, result...)
			}
			if count == total {
				close(fChannel)
				break
			}
		}

	}()
	wg.Wait()
	outputResults(results, interactive, outputmode)
}

func outputResults(result []finding.Finding, interactive bool, outputmode string) {
	if interactive {
		fmt.Println("\n--- Scanning complete ---")
		fmt.Printf("[%v] results\n", len(result))
		for _, finding := range result {
			fmt.Printf("FINDING\n-----\n%v\n", finding)
		}
	} else {
		var jsonOut []byte
		var err error
		switch strings.ToLower(outputmode) {
		case "gitlab":
			jsonOut, err = report.SerializeFindingsToGitLab(result)
		default:
			jsonOut, err = json.Marshal(result)
		}

		if err != nil {
			log.Fatal(err)
		}
		//todo(?) : support file output target or rely on caller piping to stream?
		fmt.Printf("%v", string(jsonOut))
	}
}

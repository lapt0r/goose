package app

import (
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/lapt0r/goose/internal/pkg/configuration"
	"github.com/lapt0r/goose/internal/pkg/decisionfilter"
	"github.com/lapt0r/goose/internal/pkg/finding"
	"github.com/lapt0r/goose/internal/pkg/loader"
	"github.com/lapt0r/goose/internal/pkg/regexfilter"
	"github.com/lapt0r/goose/internal/pkg/report"
)

var rules []configuration.ScanRule
var absoluteTargetPath string
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
	absoluteTargetPath, _ = filepath.Abs(targetpath)
	if interactive {
		fmt.Printf("[%v] Initializing Goose with target [%v]..\n", time.Now().Format(time.RFC1123), absoluteTargetPath)
	}
	targets = loader.GetTargets(absoluteTargetPath, commitDepth)
	if interactive {
		fmt.Printf("[%v] Got [%v] targets\n", time.Now().Format(time.RFC1123), len(targets))
	}
}

//Run : runs the Goose application
func Run(interactive bool, decisionTree bool, outputmode string, filterPaths string) {
	//todo: configurable concurrency
	var concurrency = 400
	var fChannel = make(chan []finding.Finding, concurrency)
	var semaphore = make(chan bool, concurrency)
	var results []finding.Finding
	var wg sync.WaitGroup
	var total = len(targets)
	var count = 0
	go func() {
		defer func() {
			if r := recover(); r != nil {
				println("panic:" + r.(string))
			}
		}()
		if interactive {
			fmt.Printf("[%v] Waiting for routines to finish..\n", time.Now().Format(time.RFC1123))
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
	for _, target := range targets {
		wg.Add(1)
		semaphore <- true
		go func(t loader.ScanTarget, w *sync.WaitGroup) {
			defer func() { <-semaphore }()
			if decisionTree {
				decisionfilter.ScanFile(t, fChannel, &wg)
			} else {
				regexfilter.ScanFile(t, &rules, fChannel, &wg)
			}
		}(target, &wg)
	}
	for i := 0; i < cap(semaphore); i++ {
		semaphore <- true
	}
	wg.Wait()
	results = filterResults(results, strings.Split(filterPaths, ","))
	outputResults(results, interactive, outputmode)
}

func filterResults(findings []finding.Finding, filterPaths []string) []finding.Finding {
	var result []finding.Finding
	for _, p := range filterPaths {
		for _, f := range findings {
			isMatch, _ := regexp.MatchString(p, f.Location.Path)
			if !isMatch {
				result = append(result, f)
			}
		}
	}
	return result
}

func outputResults(result []finding.Finding, interactive bool, outputmode string) {
	if interactive {
		fmt.Println("\n--- Scanning complete ---")
		fmt.Printf("[%v] [%v] results\n", time.Now().Format(time.RFC1123), len(result))
		for _, finding := range result {
			fmt.Printf("FINDING\n-----\n%v\n", finding)
		}
	} else {
		var jsonOut []byte
		var err error
		switch strings.ToLower(outputmode) {
		case "gitlab":
			jsonOut, err = report.SerializeFindingsToGitLab(result, absoluteTargetPath)
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

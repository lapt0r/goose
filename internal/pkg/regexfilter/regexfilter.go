package regexfilter

import (
	"bufio"
	"log"
	"regexp"
	"strings"
	"sync"

	"github.com/lapt0r/goose/internal/pkg/configuration"
	"github.com/lapt0r/goose/internal/pkg/finding"
	"github.com/lapt0r/goose/internal/pkg/loader"
	"github.com/lapt0r/goose/internal/pkg/reflectorfilter"
)

//ScanFile takes a path and a scan rule and returns a slice of findings
func ScanFile(target loader.ScanTarget, rules *[]configuration.ScanRule, fchannel chan []finding.Finding, waitgroup *sync.WaitGroup) {
	defer waitgroup.Done()
	input, err := loader.GetBytesFromScanTarget(target)
	var findings []finding.Finding
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	index := 1 //lines of code start at one
	for scanner.Scan() {
		for _, rule := range *rules {
			f := evaluateRule(scanner.Text(), rule)
			if !f.IsEmpty() {
				f.Location = finding.Location{Path: target.Path, Line: index}
				findings = append(findings, f)
			}
			index++
		}
	}
	fchannel <- findings
}

func evaluateRule(line string, rule configuration.ScanRule) finding.Finding {
	//kb todo: these should be constructed somewhere else and referenced by pointer
	matcher, err := regexp.Compile(rule.Rule)
	if err != nil {
		panic(err)
	}
	match := matcher.FindString(line)
	if match != "" && reflectorfilter.IsReflected(match) == false {
		return finding.Finding{
			Match:      match,
			Location:   finding.Location{},
			Rule:       rule.Rule,
			Confidence: rule.Confidence,
			Severity:   rule.Severity}
	}
	return finding.Finding{}
}

package regexfilter

import (
	"bufio"
	"fmt"
	"internal/pkg/configuration"
	"log"
	"os"
	"regexp"
)

//Finding contains the matched line, the location of the match, and the confidence of the match
type Finding struct {
	Match      string
	Location   string
	Rule       string
	Confidence float64
}

//IsEmpty returns true if all fields are default, false otherwise
func (finding *Finding) IsEmpty() bool {
	return finding.Match == "" && finding.Confidence == 0.0 &&
		finding.Location == "" && finding.Rule == ""
}

//Stringer for Finding struct
func (finding *Finding) String() string {
	return fmt.Sprintf("Finding [%v] Location [%v] Rule [%v] Confidence [%v]", finding.Match, finding.Location, finding.Rule, finding.Confidence)
}

//ScanFile takes a path and a scan rule and returns a slice of findings
func ScanFile(targetpath string, rule configuration.ScanRule) []Finding {
	file, err := os.Open(targetpath)
	result := make([]Finding, 5) //kb todo: tune this number.  expected # of findings per file?
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	index := 0
	for scanner.Scan() {
		finding := evaluateRule(scanner.Text(), index, rule)
		if !finding.IsEmpty() {
			finding.Location = fmt.Sprintf("%v : %v", targetpath, index)
			_ = append(result, finding)
		}
		index++
	}
	return result
}

func evaluateRule(line string, index int, rule configuration.ScanRule) Finding {
	//kb todo: these should be constructed somewhere else and referenced by pointer
	matcher, err := regexp.Compile(rule.Rule)
	if err != nil {
		panic(err)
	}
	match := matcher.FindString(line)
	if match != "" {
		return Finding{match, "NOTSET", rule.Rule, rule.Confidence}
	}
	return Finding{}
}

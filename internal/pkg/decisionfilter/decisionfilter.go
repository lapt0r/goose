package decisionfilter

import (
	"bufio"
	"fmt"
	"log"
	"regexp"
	"strings"
	"sync"

	"github.com/lapt0r/goose/internal/pkg/assignment"
	"github.com/lapt0r/goose/internal/pkg/decisiontree"
	"github.com/lapt0r/goose/internal/pkg/finding"
	"github.com/lapt0r/goose/internal/pkg/loader"
)

//ScanFile scans a provided target with the decision tree scan engine
func ScanFile(target loader.ScanTarget, fchannel chan []finding.Finding, waitgroup *sync.WaitGroup) {
	defer func() {
		if r := recover(); r != nil {
			println("panic:" + r.(string))
		}
	}()
	defer waitgroup.Done()
	var findings []finding.Finding
	input, err := loader.GetBytesFromScanTarget(target)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	index := 0
	for scanner.Scan() {
		finding := evaluateRule(strings.TrimSpace(scanner.Text()))
		if !finding.IsEmpty() {
			finding.Location = fmt.Sprintf("%v : %v", target.Path, index)
			findings = append(findings, finding)
		}
		index++
	}
	//an empty finding signals that we are done
	fchannel <- findings
}

func evaluateRule(input string) finding.Finding {
	var assignments = generateAssignments(input)
	var filtered = filterAssignments(assignments)
	if len(filtered) > 0 {
		return finding.Finding{
			Match:      input,
			Location:   "NOTSET",
			Rule:       "DecisionTree",
			Confidence: 0.7, //todo: decision tree rules?
			Severity:   1}
	}
	return finding.Finding{}
}

func tokenize(input string) []string {
	return tokenizeWithSeparator(input, " ")
}

func tokenizeWithSeparator(input string, separator string) []string {
	return strings.Split(input, separator)
}

func filterAssignments(slice []assignment.Assignment) []assignment.Assignment {
	result := slice[:0]
	for _, x := range slice {
		if x.IsSecret() || x.IsURLCredential() || x.HasKnownTokenPrefix() {
			if !x.IsReflected() {
				result = append(result, x)
			}
		}
	}
	return result
}

//slick: https://github.com/golang/go/wiki/SliceTricks#filtering-without-allocating
func filterXMLTags(tokens []string) []string {
	result := tokens[:0]
	for _, x := range tokens {
		if !containsXMLTag(x) {
			result = append(result, x)
		}
	}
	return result
}

func containsXMLTag(token string) bool {
	var regex, err = regexp.Compile("[<>]")
	if err != nil {
		panic(err)
	}
	return regex.MatchString(token)
}

func generateAssignments(input string) []assignment.Assignment {
	//todo: some kind of normalization
	var tokens = tokenize(input)
	var tree = decisiontree.ConstructTree(tokens)
	return generateAssignmentsRecursive(tree)
}

func generateAssignmentsRecursive(node *decisiontree.Node) []assignment.Assignment {
	//oh boy here we go
	var result = make([]assignment.Assignment, 1)
	var item = &result[0]
	var current = node
	for {
		if current.IsType() {
			item.Type = current.Value
		} else if current.IsAssignmentOperator() {
			item.Separator = current.Value
		} else if current.IsStringAssignment() || item.Separator != "" {
			item.Value = current.Value
			if current.HasNext() {
				result = append(result, generateAssignmentsRecursive(current.Next)...)
				//recursive call will catch downstream stuff, break
				break
			}
		} else {
			//do nothing (?)
			item.Name = current.Value
		}
		if !current.HasNext() {
			break
		} else {
			current = current.Next
		}
	}
	return result
}

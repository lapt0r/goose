package regexfilter

import (
	"testing"

	"github.com/lapt0r/goose/internal/pkg/configuration"
	"github.com/lapt0r/goose/internal/pkg/finding"
)

var testrule = configuration.ScanRule{Name: "TestRule", Rule: "password = \\w{4,}", Confidence: 0.5, Severity: 1}
var flexrule = configuration.ScanRule{Name: "Generic 8+ byte rule", Rule: "", Confidence: 0.9, Severity: 1}

func TestEmptyFinding(t *testing.T) {
	testFinding := finding.Finding{}
	if !testFinding.IsEmpty() {
		t.Errorf("Expected empty finding but got %v", testFinding)
	}
}

func TestEvaluateRuleMatch(t *testing.T) {
	teststring := "password = foobar"
	result := evaluateRule(teststring, testrule)
	if result.IsEmpty() {
		t.Errorf("Expected 1 result but got an empty result")
	}
	if result.Match != teststring {
		t.Errorf("Expected match to be 'password' but was '%v'", result.Match)
	}
	if result.Confidence != testrule.Confidence ||
		result.Rule != testrule.Rule {
		t.Errorf("Expected confidence 0.5 and severity 1 but found confidence %v and severity %v", result.Confidence, result.Severity)
	}
}

func TestEvaluateRuleMatchWithReflect(t *testing.T) {
	teststring := "password = \"password123\""
	result := evaluateRule(teststring, testrule)
	if !result.IsEmpty() {
		t.Errorf("Expected no result but got 1 result")
	}
}

func TestEvaluateRuleNoMatch(t *testing.T) {
	teststring := "somestring = foobar"
	result := evaluateRule(teststring, testrule)
	if !result.IsEmpty() {
		t.Errorf("Expected empty result but got a non-empty result")
	}
}

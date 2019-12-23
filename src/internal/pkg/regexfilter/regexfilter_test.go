package regexfilter

import (
	"internal/pkg/configuration"
	"testing"
)

var testrule = configuration.ScanRule{Name: "TestRule", Rule: "password", Confidence: 0.5, Severity: 1}
var flexrule = configuration.ScanRule{Name: "Generic 8+ byte rule", Rule: "", Confidence: 0.9, Severity: 1}

func TestEmptyFinding(t *testing.T) {
	testFinding := Finding{}
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
	if result.Match != "password" {
		t.Errorf("Expected match to be 'password' but was '%v'", result.Match)
	}
	if result.Confidence != testrule.Confidence ||
		result.Rule != testrule.Rule {
		t.Errorf("Expected confidence 0.5 and severity 1 but found confidence %v and severity %v", result.Confidence, result.Severity)
	}
}

func TestEvaluateRuleNoMatch(t *testing.T) {
	teststring := "somestring = foobar"
	result := evaluateRule(teststring, testrule)
	if !result.IsEmpty() {
		t.Errorf("Expected empty result but got a non-empty result")
	}
}

package regexfilter

import (
	"internal/pkg/configuration"
	"testing"
)

var testrule = configuration.ScanRule{"TestRule", "password", 0.5, 1}

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

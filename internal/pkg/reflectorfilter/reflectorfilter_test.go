package reflectorfilter

import (
	"testing"
)

var reflectedFinding = "password = \"password\""
var partialReflectedFinding = "password = \"password123\""
var valueSubsetFinding = "private static final String PASSWORD = \"password\";"
var nonReflectedFinding = "token = \"AKIA032as876tcbvc\""
var malformedReflectedFinding = " = setsecret('akiaxxxxxxxxx')"

func TestIsReflectedReturnsTrueOnReflectedString(t *testing.T) {
	result := IsReflected(reflectedFinding)
	if !result {
		t.Errorf("Expected result to be true but was %v for %v", result, reflectedFinding)
	}
}

func TestIsReflectedReturnsTrueOnPartialReflectedString(t *testing.T) {
	result := IsReflected(partialReflectedFinding)
	if !result {
		t.Errorf("Expected result to be true but was %v for %v", result, partialReflectedFinding)
	}
}

func TestIsReflectedReturnsFalseOnNonReflectedString(t *testing.T) {
	result := IsReflected(nonReflectedFinding)
	if result {
		t.Errorf("Expected result to be false but was %v for %v", result, nonReflectedFinding)
	}
}

func TestIsReflectedReturnsTrueOnSubsetString(t *testing.T) {
	result := IsReflected(valueSubsetFinding)
	if !result {
		t.Errorf("Expected result to be true but was %v for %v", result, nonReflectedFinding)
	}
}

func TestIsReflectedReturnsFalseOnMalformedString(t *testing.T) {
	result := IsReflected(malformedReflectedFinding)
	if result {
		t.Errorf("Expected result to be false but was %v for %v", result, nonReflectedFinding)
	}
}

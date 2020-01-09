package reflectorfilter

import (
	"testing"
)

var reflectedFinding = "password = \"password\""
var partialReflectedFinding = "password = \"password123\""
var nonReflectedFinding = "token = \"AKIA032as876tcbvc\""

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

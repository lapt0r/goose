package entropy

import (
	"math/rand"
	"testing"
	"math"
)

func TestGetCharacterCountOneOfEach(t *testing.T) {
	var testString = "abcde"
	result := GetCharacterCount(testString)
	if len(result) != 5 {
		t.Errorf("Expected a map with 5 keys but got %v", len(result))
	}
	for _,r := range testString {
		if result[r] != 1 {
			t.Errorf("Character %v expected count 1, got count %v", r, result[r])
		}
	}
}

func TestGetCharacterCountRepeatLetter(t *testing.T) {
	var testString = "aaaaa"
	result := GetCharacterCount(testString)
	if len(result) != 1 {
		t.Errorf("Expected a map with 1 key but got %v", len(result))
	}
	for _,r := range testString {
		if result[r] != 5 {
			t.Errorf("Character %v expected count 5, got count %v", r, result[r])
		}
	}
}

func TestGetPValuesOneOfEach(t *testing.T) {
	var testString = "ABCDE"
	result := GetPValues(testString)
	if len(result) != 5 {
		t.Errorf("Expected to get 5 P values but got %v", len(result))
	}
	for val := range result {
		if result[val] != float64(0.2) {
			t.Errorf("Expected P value of 0.2 but got %v", val)
		}
	}
}

func TestGetPValuesRepeatLetter(t *testing.T) {
	var testString = "AAAAA"
	result := GetPValues(testString)
	if len(result) != 1 {
		t.Errorf("Expected 1 P value but found %v", len(result))
	}
	if result[0] != float64(1) {
		t.Errorf("Expected P value of 1 but found %v", result[0])
	}
}

func TestGetShannonEntropyZeroEntropy(t *testing.T) {
	var testString = "AAAAAA"
	result := GetShannonEntropy(testString)
	if result != float64(0) {
		t.Errorf("Expected entropy value of 0 but got %v", result)
	}
}

func TestGetShannonEntropyMaxEntropy(t *testing.T) {
	var testString = "abcde"
	var maximum = math.Log(float64(len(testString)))
	var epsilon = math.Exp2(-float64(rand.Intn(10)))
	result := GetShannonEntropy(testString)
	//shannon entropy is a numerical approximation that converges to ln(n) where n is the length of the string
	//we are using the squeeze theorem here to generate an arbitrary small epsilon and confirm that the result is within +/- epsilon of the maximum
	if !(result < maximum + epsilon  && maximum - epsilon < result) {
		t.Errorf("Expected entropy value of %v but got %v", maximum, result)
	}
}
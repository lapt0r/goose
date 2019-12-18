package entropy

import "testing"

func TestGetCharacterCountOneOfEach(t *testing.T) {
	var testString = "abcde"
	result := GetCharacterCount(testString)
	if len(result) != 5 {
		t.Errorf("Expected a map with 5 keys but got %v", len(result))
	}
	for char := range testString {
		if result[string(char)] != 1 {
			t.Errorf("Character %v expected count 1, got count %v", string(char), result[string(char)])
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
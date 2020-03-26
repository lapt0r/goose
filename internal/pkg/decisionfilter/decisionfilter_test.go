package decisionfilter

import "testing"

func TestTokenize(t *testing.T) {
	var testString = "a b c"
	var result = tokenize(testString)

	if len(result) != 3 {
		t.Errorf("Expected 3 items but found %v", len(result))
	}

	if result[0] != "a" {
		t.Errorf("Expected first element to be 'a'  but was '%v'", result[0])
	}
	if result[1] != "b" {
		t.Errorf("Expected second element to be 'b'  but was '%v'", result[0])
	}
	if result[2] != "c" {
		t.Errorf("Expected third element to be 'c'  but was '%v'", result[0])
	}
}

func TestTokenizeWithSeparator(t *testing.T) {
	var testString = "a;b;c"
	var result = tokenizeWithSeparator(testString, ";")

	if len(result) != 3 {
		t.Errorf("Expected 3 items but found %v", len(result))
	}

	if result[0] != "a" {
		t.Errorf("Expected first element to be 'a'  but was '%v'", result[0])
	}
	if result[1] != "b" {
		t.Errorf("Expected second element to be 'b'  but was '%v'", result[0])
	}
	if result[2] != "c" {
		t.Errorf("Expected third element to be 'c'  but was '%v'", result[0])
	}
}

func TestContainsXMLTagHasTag(t *testing.T) {
	var testString = "<xml/>"
	var result = containsXMLTag(testString)
	if !result {
		t.Errorf("expected [%v] to contain an xml tag.", testString)
	}
}

func TestContainsXMLTagNoTag(t *testing.T) {
	var testString = "blah blah something xml blah/!@#$%"
	var result = containsXMLTag(testString)
	if result {
		t.Errorf("expected [%v] not to contain an xml tag.", testString)
	}
}

func TestXMLFilter(t *testing.T) {
	var testString = "<?xml version=\"1.0\" encoding=\"utf-8\" ?>"
	var result = filterXMLTags(tokenize(testString))
	if len(result) != 2 {
		t.Errorf("Expected 2 results but found %v", len(result))
	}
	if result[0] != "version=\"1.0\"" {
		t.Errorf("Expected first result to be [version=\"1.0\"]  but was [%v]", result[0])
	}
	if result[1] != "encoding=\"utf-8\"" {
		t.Errorf("Expected first result to be [encoding=\"utf-8\"]  but was [%v]", result[1])
	}
}

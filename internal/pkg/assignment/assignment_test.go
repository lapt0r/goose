package assignment

import "testing"

func TestArrayLiteral(t *testing.T) {
	testAssignment := Assignment{
		Type:      "var",
		Name:      "myarray",
		Separator: "=",
		Value:     "[1,2,3,4,5]"}
	if !testAssignment.IsArrayAssignment() {
		t.Errorf("expected array to be an array assignment but was not an array assignment.")
	}
}

func TestDictLiteral(t *testing.T) {
	testAssignment := Assignment{
		Type:      "var",
		Name:      "mydict",
		Separator: "=",
		Value:     "{1,2,3,4,5}"}
	if !testAssignment.IsDictAssignment() {
		t.Errorf("expected dict to be a dict assignment but was not a dict assignment.")
	}
}

func TestInvalidValue(t *testing.T) {
	testAssignment := Assignment{
		Type:      "var",
		Name:      "secret",
		Separator: "=",
		Value:     "0;"}
	if testAssignment.IsValidValue() {
		t.Errorf("Value %v was incorrectly considered valid.", testAssignment.Value)
	}
}

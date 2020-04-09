package decisionfilter

import (
	"testing"

	"github.com/lapt0r/goose/internal/pkg/decisiontree"
)

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

func TestGenerateAssignmentsRecursiveDoubleQuote(t *testing.T) {
	var testString = "foo = \"bar\""
	var tree = decisiontree.ConstructTree(tokenize(testString))
	var result = generateAssignmentsRecursive(tree)
	if len(result) != 1 {
		t.Errorf("Expected 1 result in collection but got %v", len(result))
	}
	var assignment = result[0]
	if assignment.Name != "foo" {
		t.Errorf("expected name [foo] but found [%v]", assignment.Name)
	}
	if assignment.Value != "\"bar\"" {
		t.Errorf("expected value [\"bar\"] but found [%v]", assignment.Value)
	}
}

func TestGenerateAssignmentsRecursiveSingleQuote(t *testing.T) {
	var testString = "foo = 'bar'"
	var tree = decisiontree.ConstructTree(tokenize(testString))
	var result = generateAssignmentsRecursive(tree)
	if len(result) != 1 {
		t.Errorf("Expected 1 result in collection but got %v", len(result))
	}
	var assignment = result[0]
	if assignment.Name != "foo" {
		t.Errorf("expected name [foo] but found [%v]", assignment.Name)
	}
	if assignment.Value != "'bar'" {
		t.Errorf("expected value ['bar'] but found [%v]", assignment.Value)
	}
}

func TestGenerateAssignmentsRecursiveNoQuote(t *testing.T) {
	var testString = "foo = bar"
	var tree = decisiontree.ConstructTree(tokenize(testString))
	var result = generateAssignmentsRecursive(tree)
	if len(result) != 1 {
		t.Errorf("Expected 1 result in collection but got %v", len(result))
	}
	var assignment = result[0]
	if assignment.Name != "foo" {
		t.Errorf("expected name [foo] but found [%v]", assignment.Name)
	}
	if assignment.Value != "bar" {
		t.Errorf("expected value [bar] but found [%v]", assignment.Value)
	}
}

func TestGenerateAssignmentsRecursiveMultiAssignment(t *testing.T) {
	var testString = "foo = bar biz : baz fizz := buzz"
	var tree = decisiontree.ConstructTree(tokenize(testString))
	var result = generateAssignmentsRecursive(tree)
	if len(result) != 3 {
		t.Errorf("Expected 3 results in collection but got %v", len(result))
	}
	var assignment1 = result[0]
	if assignment1.Name != "foo" {
		t.Errorf("expected name [foo] but found [%v]", assignment1.Name)
	}
	if assignment1.Value != "bar" {
		t.Errorf("expected value [bar] but found [%v]", assignment1.Value)
	}
	var assignment2 = result[1]
	if assignment2.Name != "biz" {
		t.Errorf("expected name [biz] but found [%v]", assignment2.Name)
	}
	if assignment2.Value != "baz" {
		t.Errorf("expected value [baz] but found [%v]", assignment2.Value)
	}
	var assignment3 = result[2]
	if assignment3.Name != "fizz" {
		t.Errorf("expected name [fizz] but found [%v]", assignment3.Name)
	}
	if assignment3.Value != "buzz" {
		t.Errorf("expected value [buzz] but found [%v]", assignment3.Value)
	}
}

func TestGenerateAssignmentsMultiAssignment(t *testing.T) {
	var testString = "foo = bar biz : baz"
	var result = generateAssignments(testString)
	if len(result) != 2 {
		t.Errorf("Expected 2 results in collection but got %v", len(result))
	}
	var assignment1 = result[0]
	if assignment1.Name != "foo" {
		t.Errorf("expected name [foo] but found [%v]", assignment1.Name)
	}
	if assignment1.Value != "bar" {
		t.Errorf("expected value [bar] but found [%v]", assignment1.Value)
	}
	var assignment2 = result[1]
	if assignment2.Name != "biz" {
		t.Errorf("expected name [biz] but found [%v]", assignment2.Name)
	}
	if assignment2.Value != "baz" {
		t.Errorf("expected value [baz] but found [%v]", assignment2.Value)
	}
}

func TestEvaluateRuleMatch(t *testing.T) {
	teststring := "password = foobar"
	result := evaluateRule(teststring)
	if result.IsEmpty() {
		t.Errorf("Expected 1 result but got an empty result")
	}
	if result.Match != teststring {
		t.Errorf("Expected match to be 'password' but was '%v'", result.Match)
	}
	if result.Confidence != 0.7 ||
		result.Rule != "DecisionTree" {
		t.Errorf("Expected confidence 0.5 and severity 1 but found confidence %v and severity %v", result.Confidence, result.Severity)
	}
}

func TestEvaluateRuleURLCredential(t *testing.T) {
	teststring := "admin:password123@example.com"
	result := evaluateRule(teststring)
	if result.IsEmpty() {
		t.Errorf("expected 1 result but got an empty result")
	}
	if result.Match != teststring {
		t.Errorf("Expected match to be like admin:password123 but was %v", result.Match)
	}
}

func TestEvaluateRuleMethodAssignmentEdgeCase(t *testing.T) {
	teststring := "setAmazonKeys ( 'AKIAxxxxxxxxxxxxx' , 'ePvh4kxxxxxxxxxxxxxxxxxxxx');"
	result := evaluateRule(teststring)
	if result.IsEmpty() {
		t.Errorf("expected 1 result but got an empty result")
	}
	if result.Match != teststring {
		t.Errorf("Expected match to be like %v but was %v", teststring, result.Match)
	}
}

func TestEvaluateRuleSecretNominal(t *testing.T) {
	teststring := "var client_secret = 'foobarbiz'"
	result := evaluateRule(teststring)
	if result.IsEmpty() {
		t.Errorf("expected 1 result but got an empty result")
	}
	if result.Match != teststring {
		t.Errorf("Expected match to be like %v but was %v", teststring, result.Match)
	}
}

func TestEvaluateRuleCompositeSecret(t *testing.T) {
	teststring := "secretkey = 'foobarbiz'"
	result := evaluateRule(teststring)
	if result.IsEmpty() {
		t.Errorf("expected 1 result but got an empty result")
	}
	if result.Match != teststring {
		t.Errorf("Expected match to be like %v but was %v", teststring, result.Match)
	}
}

func TestEvaluateRuleSecretMethodCall(t *testing.T) {
	//to investigate: how can we catch this in a more abstract fashion with the parser?
	teststring := "setsecret('AKIAxxxxxxxxx')"
	result := evaluateRule(teststring)
	if result.IsEmpty() {
		t.Errorf("expected 1 result but got an empty result")
	}
	if result.Match != teststring {
		t.Errorf("Expected match to be like %v but was %v", teststring, result.Match)
	}
}

func TestEvaluateRuleFalsePositiveBookTitle(t *testing.T) {
	teststrings := [...]string{"The Secret Kingdom: Stones of Ravenglass", "9780963146403,6778761,\"Quality Secret : The Right Way to Manage\",0,"}
	for _, teststring := range teststrings {
		result := evaluateRule(teststring)
		if !result.IsEmpty() {
			t.Errorf("expected no results for %v", teststring)
		}
	}
}

func TestEvaluateRuleFalsePositiveMapKeyAssignment(t *testing.T) {
	teststring := "key = key.map( jQuery.camelCase );"
	result := evaluateRule(teststring)
	if !result.IsEmpty() {
		t.Errorf("expected no results")
	}
}

func TestEvaluateRuleFalsePositiveIntegerSwitch(t *testing.T) {
	teststring := "int forcePasswordChange = 0;"
	result := evaluateRule(teststring)
	if !result.IsEmpty() {
		t.Errorf("expected no results")
	}
}

func TestEvaluateRuleFalsePositiveTestCode(t *testing.T) {
	teststring := "$password = 'testing';"
	result := evaluateRule(teststring)
	if !result.IsEmpty() {
		t.Errorf("expected no results")
	}
}

func TestEvaluateRuleFalsePositiveArgumentValueIndexing(t *testing.T) {
	teststring := "$password = argv[0];"
	result := evaluateRule(teststring)
	if !result.IsEmpty() {
		t.Errorf("expected no results")
	}
}

func TestEvaluateRuleFalsePositiveFunctionDefinitionPasswordNullDefault(t *testing.T) {
	teststring := "public function withUserInfo($user, $password = null);"
	result := evaluateRule(teststring)
	if !result.IsEmpty() {
		t.Errorf("expected no results")
	}
}

func TestEvaluateRuleFalsePositiveFunctionDefinitionImageBlob(t *testing.T) {
	teststring := "        'vYRjtYRjrXtjpXtjlGNje2tazoxazoRaxoxaxoRavYRatYRatX'."
	result := evaluateRule(teststring)
	if !result.IsEmpty() {
		t.Errorf("expected no results")
	}
}

func TestEvaluateRuleFalsePositiveReadRequestData(t *testing.T) {
	teststrings := [...]string{"$admin_password = isset($_POST['admin_password']) ? $_POST['admin_password'] : '';"}
	for _, teststring := range teststrings {
		result := evaluateRule(teststring)
		if !result.IsEmpty() {
			t.Errorf("expected no results for %v", teststring)
		}
	}
}

func TestEvaluateRuleFalsePositiveReflected(t *testing.T) {
	teststrings := [...]string{"private static final String PASSWORD = \"password\";"}
	for _, teststring := range teststrings {
		result := evaluateRule(teststring)
		if !result.IsEmpty() {
			t.Errorf("expected no results for %v", teststring)
		}
	}
}

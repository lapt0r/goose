package decisiontree

import "testing"

func TestNodeIsType(t *testing.T) {
	var node = Node{Value: "string", Next: nil}
	if !node.IsType() {
		t.Errorf("Expected node to be a type.")
	}
}

func TestNodeIsAssignmentOperator(t *testing.T) {
	var node = Node{Value: "=", Next: nil}
	if !node.IsAssignmentOperator() {
		t.Errorf("expected [%v] to be an assignment", node.Value)
	}
}

func TestNodeIsAssignmentOperatorAssignmentString(t *testing.T) {
	var node = Node{Value: "foo=bar", Next: nil}
	if node.IsAssignmentOperator() {
		t.Errorf("expected [%v] to not be an assignment", node.Value)
	}
}

func TestConstructTree(t *testing.T) {
	var testTokens = []string{"foo", "bar", "biz"}
	var this = ConstructTree(testTokens)
	count := 0
	for {
		count++
		//count panic button in case of accidental circular linked list
		if count > 3 || !this.HasNext() {
			break
		} else {
			this = this.Next
		}
	}
	if count != 3 {
		t.Errorf("expected 3 elements but found %v", count)
	}
}

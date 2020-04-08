package decisiontree

import (
	"regexp"
)

var byteRegex = "0x[a-fA-F0-9]{2}"
var doubleQuoteStringRegex = "\"[^\"]*\""
var singleQuoteStringRegex = "'[^']*'"

//Node is a singly-linked list element in the decision tree
type Node struct {
	Value string
	Next  *Node
}

//IsType returns whether or not the node is a type declaration
func (node *Node) IsType() bool {
	regex, _ := regexp.Compile("(?i)(^var$|byte|string|int\\d{2}?|float|double)")
	return regex.MatchString(node.Value)
}

//IsAssignmentOperator returns whether or not the node is an assignment operator
func (node *Node) IsAssignmentOperator() bool {
	regex, _ := regexp.Compile("^(=|<-|:=|:)$")
	return regex.MatchString(node.Value)
}

//IsStringAssignment returns whether or not the value matches a single or double quoted string
func (node *Node) IsStringAssignment() bool {
	singleQuoteRegex, _ := regexp.Compile(singleQuoteStringRegex)
	doubleQuoteRegex, _ := regexp.Compile(doubleQuoteStringRegex)
	return singleQuoteRegex.MatchString(node.Value) || doubleQuoteRegex.MatchString(node.Value)
}

//HasNext returns whether or not there is a next node in the list
func (node *Node) HasNext() bool {
	return node.Next != nil
}

//IsEmpty returns whether or not the node is empty
func (node *Node) IsEmpty() bool {
	return node.Value == "" && node.Next == nil
}

//ConstructTree generates a Node list from a token stream
func ConstructTree(tokens []string) *Node {
	var head *Node
	var last *Node
	for _, token := range tokens {
		this := Node{Value: token, Next: nil}
		if head == nil {
			head = &this
		} else {
			last.Next = &this
		}
		last = &this
	}
	return head
}

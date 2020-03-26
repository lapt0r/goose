package decisionfilter

import (
	"regexp"
	"strings"

	"github.com/lapt0r/goose/internal/pkg/decisiontree"
)

//Assignment is an externally visible struct to aid in deconstructing complex assignment statements into their conceptual base form
type Assignment struct {
	Type      string
	Name      string
	Separator string
	Value     string
}

func tokenize(input string) []string {
	return tokenizeWithSeparator(input, " ")
}

func tokenizeWithSeparator(input string, separator string) []string {
	return strings.Split(input, separator)
}

//slick: https://github.com/golang/go/wiki/SliceTricks#filtering-without-allocating
func filterXMLTags(tokens []string) []string {
	result := tokens[:0]
	for _, x := range tokens {
		if !containsXMLTag(x) {
			result = append(result, x)
		}
	}
	return result
}

func containsXMLTag(token string) bool {
	var regex, err = regexp.Compile("[<>]")
	if err != nil {
		panic(err)
	}
	return regex.MatchString(token)
}

func generateAssignments(input string) []Assignment {
	//todo: some kind of normalization
	var tokens = tokenize(input)
	var tree = decisiontree.ConstructTree(tokens)
}

func generateAssignmentsRecursive(node Node) []Assignment {
	//oh boy here we go
	for {
		if
	}
}

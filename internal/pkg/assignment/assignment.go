package assignment

import (
	"fmt"
	"regexp"

	"github.com/lapt0r/goose/internal/pkg/reflectorfilter"
)

//Assignment is an externally visible struct to aid in deconstructing complex assignment statements into their conceptual base form
type Assignment struct {
	Type      string
	Name      string
	Separator string
	Value     string
}

//IsSecret returns whether or not the Assignment is considered to be a secret assignment
func (assignment *Assignment) IsSecret() bool {
	secretregex, _ := regexp.Compile("(?i)(secret($|_|[()]|key)|password$|(api|access)_?key|connection[a-z0-9\\-_]*string)")
	return assignment.IsValidValue() && secretregex.MatchString(assignment.Name) && assignment.Separator != "" && assignment.IsKnownSecretAssignmentType() && !secretregex.MatchString(assignment.Value)
}

//IsConfigAssignment returns whether or not the assignment matches a configuration value
func (assignment *Assignment) IsConfigAssignment() bool {
	return assignment.Name != "" && assignment.Separator != "" && (assignment.IsStringLiteral() || assignment.IsUnquotedString())
}

//IsKnownSecretAssignmentType returns whether or not the assignment is a known secret assignment type
func (assignment *Assignment) IsKnownSecretAssignmentType() bool {
	return assignment.IsArrayAssignment() || assignment.IsDictAssignment() || assignment.IsStringLiteral() || assignment.IsConfigAssignment()
}

//IsURLCredential returns whether or not the assignment matches a url credential
func (assignment *Assignment) IsURLCredential() bool {
	urlcredregex, _ := regexp.Compile("(?i)\\w+:\\w+@[a-z0-9\\-]+\\.[a-z]{2,5}($|[\"']|/)")
	mailto, _ := regexp.Compile("mailto")
	//thought: this is either a config value (likely unquoted) or a hardcoded string
	nameIsURLCredential := urlcredregex.MatchString(assignment.Name) && !mailto.MatchString(assignment.Name)
	valueIsURLCredential := urlcredregex.MatchString(assignment.Value) && !mailto.MatchString(assignment.Value)
	return nameIsURLCredential || valueIsURLCredential
}

//HasKnownTokenPrefix returns whether or not the value has a known token prefix
func (assignment *Assignment) HasKnownTokenPrefix() bool {
	//amazon, facebook, google, slack
	knownprefixregex, _ := regexp.Compile("(^|[\"'])(AKIA|EAACEdEose0cBA|AIza|xox[pboa])[a-zA-Z0-9/_\\-]+")
	return knownprefixregex.MatchString(assignment.Value)
}

//IsTokenAssignment checks that the name matches 'token' and the value is a string literal
func (assignment *Assignment) IsTokenAssignment() bool {
	tokenregex, _ := regexp.Compile("(?i)token")
	return tokenregex.MatchString(assignment.Name) && assignment.IsStringLiteral()
}

//IsStringLiteral checks if the assignment is a string literal
func (assignment *Assignment) IsStringLiteral() bool {
	stringliteral, _ := regexp.Compile("[\"'][^\"']'+[\"']")
	return stringliteral.MatchString(assignment.Value)
}

//IsUnquotedString checks if the assignment is an unquoted string
func (assignment *Assignment) IsUnquotedString() bool {
	unquotedstring, _ := regexp.Compile("\\S+")
	return unquotedstring.MatchString(assignment.Value)
}

//IsArrayAssignment checks if the assignment value is an array
func (assignment *Assignment) IsArrayAssignment() bool {
	arrayliteral, _ := regexp.Compile("\\[[^\\]\\[]+\\]")
	return arrayliteral.MatchString(assignment.Value)
}

//IsDictAssignment checks if the assignment is a dictionary
func (assignment *Assignment) IsDictAssignment() bool {
	dictliteral, _ := regexp.Compile("{[^{}]+}")
	return dictliteral.MatchString(assignment.Value)
}

//IsValidValue returns whether or not the value is invalid (true, false, or empty object/string)
func (assignment *Assignment) IsValidValue() bool {
	invalidvalue, _ := regexp.Compile("(?i)(true|false|\\[\\]|\\{\\}|''|\"\"|[01];?$|\\[\\d+\\]|test|null\\)?;?$|get|post)")
	//valid secrets are 6 or more characters in 2020 systems.  Less than that is trivially brute-forceable.
	return !invalidvalue.MatchString(assignment.Value) && len(assignment.Value) > 5
}

//IsReflected returns whether or not the assignment is reflected
func (assignment *Assignment) IsReflected() bool {
	return reflectorfilter.IsReflected(fmt.Sprintf("%v = %v", assignment.Name, assignment.Value))
}

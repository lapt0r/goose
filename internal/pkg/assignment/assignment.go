package assignment

import "regexp"

//Assignment is an externally visible struct to aid in deconstructing complex assignment statements into their conceptual base form
type Assignment struct {
	Type      string
	Name      string
	Separator string
	Value     string
}

//IsSecret returns whether or not the Assignment is considered to be a secret assignment
func (assignment *Assignment) IsSecret() bool {
	secretregex, _ := regexp.Compile("(?i)(secret|password|key|connection[a-z0-9\\-_]*string)")
	return secretregex.MatchString(assignment.Name) && assignment.Separator != "" && assignment.Value != ""
}

//IsURLCredential returns whether or not the assignment matches a url credential
func (assignment *Assignment) IsURLCredential() bool {
	urlcredregex, _ := regexp.Compile("(?i)\\w+:\\w+@[a-z0-9\\-]+\\.[a-z]{2,5}")
	//thought: this is either a config value (likely unquoted) or a hardcoded string
	return urlcredregex.MatchString(assignment.Name) || urlcredregex.MatchString(assignment.Value)
}

//HasKnownTokenPrefix returns whether or not the value has a known token prefix
func (assignment *Assignment) HasKnownTokenPrefix() bool {
	//amazon, facebook, google, slack
	knownprefixregex, _ := regexp.Compile("(?i)(AKIA|EAACEdEose0cBA|AIza|xox[pboa])")
	return knownprefixregex.MatchString(assignment.Value)
}

//IsTokenAssignment checks that the name matches 'token' and the value is a string literal
func (assignment *Assignment) IsTokenAssignment() bool {
	tokenregex, _ := regexp.Compile("(?i)token")
	stringliteral, _ := regexp.Compile("[\"'][^\"]+[\"']")
	return tokenregex.MatchString(assignment.Name) && !stringliteral.MatchString(assignment.Value)
}

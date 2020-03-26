package finding

import "fmt"

//Finding contains the matched line, the location of the match, and the confidence of the match
type Finding struct {
	Match      string
	Location   string
	Rule       string
	Confidence float64
	Severity   int
}

//IsEmpty returns true if all fields are default, false otherwise
func (finding *Finding) IsEmpty() bool {
	return finding.Match == "" && finding.Confidence == 0.0 &&
		finding.Location == "" && finding.Rule == "" && finding.Severity == 0
}

//Stringer for Finding struct
func (finding Finding) String() string {
	return fmt.Sprintf("Finding [%v] Location [%v] Rule [%v] Confidence [%v]", finding.Match, finding.Location, finding.Rule, finding.Confidence)
}

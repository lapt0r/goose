package finding

import (
	"fmt"

	gli "gitlab.com/gitlab-org/security-products/analyzers/common/v2/issue"
)

//Finding contains the matched line, the location of the match, and the confidence of the match
type Finding struct {
	Match      string
	Location   Location
	Rule       string
	Confidence float64
	Severity   string
}

//Location wraps the path/line data for a finding
type Location struct {
	Path string
	Line int
}

//IsEmpty returns true if all fields are default, false otherwise
func (finding *Finding) IsEmpty() bool {
	return finding.Match == "" && finding.Confidence == 0.0 &&
		finding.Location == Location{} && finding.Rule == "" && finding.Severity == ""
}

//ConvertSeverityToGitLab converts severity to GitLab-compatible int.
func (finding *Finding) ConvertSeverityToGitLab() gli.SeverityLevel {
	return gli.ParseSeverityLevel(finding.Severity)
}

//Stringer for Finding struct
func (finding Finding) String() string {
	return fmt.Sprintf("Finding [%v] Location [%v] Rule [%v] Confidence [%v]", finding.Match, finding.Location, finding.Rule, finding.Confidence)
}

//Stringer for Location struct
func (location Location) String() string {
	return fmt.Sprintf("%v : %v", location.Path, location.Line)
}

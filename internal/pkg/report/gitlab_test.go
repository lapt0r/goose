package report

import (
	"encoding/json"
	"testing"

	"github.com/lapt0r/goose/internal/pkg/finding"
	gli "gitlab.com/gitlab-org/security-products/analyzers/common/v2/issue"
)

func TestSerializeGitLabReport(t *testing.T) {
	l := finding.Location{
		Path: "/foo/bar/biz.php",
		Line: 10}
	f := finding.Finding{
		Match:      "password = s00persekr1t",
		Location:   l,
		Rule:       "DecisionTree",
		Confidence: 0.7,
		Severity:   "high",
	}
	findings := []finding.Finding{f}
	report, err := SerializeFindingsToGitLab(findings, "/foo/")
	if err != nil {
		t.Errorf("Error while serializing json: %v", err)
	}
	var r gli.Report
	json.Unmarshal(report, &r)
	if len(r.Vulnerabilities) == 0 {
		t.Errorf("Expected 1 vulnerability but found 0")
	}
	//this is a test of GitLab path-relative output
	if r.Vulnerabilities[0].Location.File != "bar/biz.php" {
		t.Errorf("Path from finding %v did not match path from report %v", f.Location.Path, r.Vulnerabilities[0].Location.File)
	}
}

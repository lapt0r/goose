package report

import (
	"encoding/json"
	"path/filepath"

	"github.com/lapt0r/goose/internal/pkg/finding"
	gli "gitlab.com/gitlab-org/security-products/analyzers/common/v2/issue"
)

//SerializeFindingsToGitLab serializes Goose internals to GitLab-consumable format
func SerializeFindingsToGitLab(findings []finding.Finding, basepath string) ([]byte, error) {
	report := gli.NewReport()
	for _, finding := range findings {
		relativePath, _ := filepath.Rel(basepath, finding.Location.Path)
		issue := gli.Issue{
			Category:    gli.CategorySast,
			Name:        "Secret in source control",
			Message:     "Secret exposed in source control",
			Description: "Secrets committed to source control can allow users with access to source to potentially perform actions in the context of the account associated with the secret.  Contractors and other non-employees often get temporary access to source; it should not be treated as a trust boundary!",
			Severity:    gli.SeverityLevelCritical,
			Location: gli.Location{
				File:      relativePath,
				LineStart: finding.Location.Line},
			Scanner: gli.Scanner{
				ID:   "goose",
				Name: "Goose"}}
		report.Vulnerabilities = append(report.Vulnerabilities, issue)
	}
	bytes, err := json.Marshal(report)
	return bytes, err
}

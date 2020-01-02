package loader

import (
	"testing"
)

var firstCommitString = "17ed9ccdcc40b4704a2eb4df57afeb344932362b"
var parentDirectory = "../../../../"
var expectedContent = `# Goose
It's a lovely day in source control, and you are a horrible goose.
`

func TestGitCommitLoad(t *testing.T) {
	content := getGitBytesImpl(parentDirectory, firstCommitString, "README.md")
	result := string(content)
	if result != expectedContent {
		t.Errorf("expected %v but found %v", expectedContent, result)
	}
}

func TestGitTargetEnumeration(t *testing.T) {
	targets, enumErr := EnumerateRepositoryCommits(parentDirectory)
	t.Logf("Found %v targets", len(targets))
	if enumErr != nil {
		t.Errorf("Error enumerating repository : %v", enumErr)
	}
	if len(targets) < 15 {
		t.Errorf("Expected at least 15 commits but found %v", len(targets))
	}
}

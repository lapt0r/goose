package loader

import "testing"

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

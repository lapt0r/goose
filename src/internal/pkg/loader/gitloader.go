package loader

import (
	"fmt"
	"io"
	"log"
	"strings"

	git "gopkg.in/src-d/go-git.v4"
	gitplumbing "gopkg.in/src-d/go-git.v4/plumbing"
	gitobject "gopkg.in/src-d/go-git.v4/plumbing/object"
)

//LoadGitTargets : Loads content from a provided git hash and parent directory
func LoadGitTargets(commit string, parent string) ([]ScanTarget, error) {
	var result []ScanTarget
	repository, openerr := git.PlainOpen(parent)
	if openerr != nil {
		fmt.Printf("Error opening repository %v : %v", parent, openerr)
	}

	//kb todo: enumerate all commits from repository and refactor this out to a method call to append targets to result
	commithash := gitplumbing.NewHash(commit)
	commitContent, err := repository.CommitObject(commithash)
	fileiterator, _ := commitContent.Files()
	for {
		file, err := fileiterator.Next()
		if err == io.EOF {
			break
		}
		result = append(result, ScanTarget{Path: fmt.Sprintf("%v:%v:%v", parent, commit, file.Name), Type: "git"})
	}
	return result, err
}

func getCommitData(file *gitobject.File) []byte {
	//do something with file arg
	isBinary, bErr := file.IsBinary()
	if bErr == nil && isBinary == false {
		contents, err := file.Contents()
		if err == nil {
			return []byte(contents)
		}
	}
	return nil
}

func rehydratePathString(path string) (string, string, string) {
	components := strings.Split(path, ":")
	return components[0], components[1], components[2]
}

func getGitBytes(path string) []byte {
	parent, commit, filename := rehydratePathString(path)
	return getGitBytesImpl(parent, commit, filename)
}

func getGitBytesImpl(parent string, commit string, filename string) []byte {
	repository, openerr := git.PlainOpen(parent)
	if openerr != nil {
		fmt.Printf("Error opening repository %v : %v", parent, openerr)
	}
	commithash := gitplumbing.NewHash(commit)
	commitContent, err := repository.CommitObject(commithash)
	if err != nil {
		log.Fatal(err)
	}
	file, fErr := commitContent.File(filename)
	if fErr != nil {
		log.Fatal(fErr)
	}
	return getCommitData(file)
}

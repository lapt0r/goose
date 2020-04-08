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

//EnumerateRepositoryFileChanges : Enumerate the commmits in a target repository and return the assembled ScanTargets
func EnumerateRepositoryFileChanges(parent string, commitDepth int) ([]ScanTarget, error) {
	var result []ScanTarget
	repository, err := git.PlainOpen(parent)
	if err == nil {
		options := git.LogOptions{Order: git.LogOrderCommitterTime}
		commits, _ := repository.Log(&options)
		for i := 0; i < commitDepth; i++ {
			next, err := commits.Next()
			if err != nil {
				//this is an EOF error which tells us to terminate the loop
				break
			}
			targets, err := LoadGitTargets(next, parent)
			result = append(result, targets...)
		}
	}
	return result, err
}

//LoadGitTargets : Loads content from a provided git hash and parent directory
func LoadGitTargets(commit *gitobject.Commit, parent string) ([]ScanTarget, error) {
	var result []ScanTarget
	fileiterator, err := commit.Files()
	for {
		file, fileErr := fileiterator.Next()
		if fileErr == io.EOF {
			break
		}
		result = append(result, ScanTarget{Path: fmt.Sprintf("%v|%v|%v", parent, commit.Hash, file.Name), Type: "git"})
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
	//kb todo: return error here?
	return nil
}

func rehydratePathString(path string) (string, string, string) {
	components := strings.Split(path, "|")
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

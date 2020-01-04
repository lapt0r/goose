package loader

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

//ScanTarget : Struct to abstract loading of filesystem and git commit history targets
type ScanTarget struct {
	Path string
	Type string
}

//GetTargets returns all files that are children of the target path
func GetTargets(parent string) []ScanTarget {
	var files []ScanTarget
	//kb todo: this inline function could use its own unit test, to refactor out
	err := filepath.Walk(parent, func(path string, info os.FileInfo, err error) error {
		info, infoerr := os.Stat(path)
		if infoerr != nil {
			fmt.Printf("error loading %v : %v", path, infoerr)
		} else {
			if !info.IsDir() {
				if ValidateContent(path) {
					files = append(files, ScanTarget{Path: path, Type: "filesystem"})
				}
			} else {
				//kb todo: git repository handling
			}
		}
		return nil
	})
	if err != nil {
		//kb note : currently, walk function does not actually return an error.  This block will never be hit in normal operation
		log.Fatal(err)
	}
	//load git targets
	//kb todo: walk subdirectories of target and enumerate repository commits of every detected repo
	gitTargets, gitErr := EnumerateRepositoryCommits(parent)
	if gitErr == nil {
		files = append(files, gitTargets...)
	}
	return files
}

//GetBytesFromScanTarget : Receives a ScanTarget struct, returns a byte slice of file contents (filesystem) or patch data (git)
func GetBytesFromScanTarget(target ScanTarget) ([]byte, error) {
	switch target.Type {
	case "filesystem":
		return getFileSystemBytes(target.Path), nil
	case "git":
		return getGitBytes(target.Path), nil
	default:
		return nil, fmt.Errorf("Failed to load bytes from %v.  Unknown type %v", target.Path, target.Type)
	}
}

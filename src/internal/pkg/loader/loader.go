package loader

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gogs/chardet"
)

//GetTargets returns all files that are children of the target path
func GetTargets(parent string) []string {
	var files []string
	err := filepath.Walk(parent, func(path string, info os.FileInfo, err error) error {
		//todo: validate file format
		if ValidateContent(path) {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		//kb todo - handle errors gracefully when file processing
		panic(err)
	}
	return files
}

//ValidateContent ensures that the contents of the text are valid binary encodings
func ValidateContent(path string) bool {
	var validBytes bool = false
	var file, err = os.Open(path)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	var headerBytes = make([]byte, 1024)
	var bytesRead, headerReadError = file.Read(headerBytes)
	if headerReadError != nil {
		validBytes = false
		fmt.Printf("Error reading file [%v] : %v", path, headerReadError)
	}
	//kb note: important to slice to the actual number of bytes read
	var charSet = getByteCharset(headerBytes[0:bytesRead])
	//kb note: should support other encodings other than ISO-8859-1
	if charSet.Charset == "ISO-8859-1" {
		validBytes = true
	}
	return validBytes
}

func getByteCharset(b []byte) *chardet.Result {
	var detector = chardet.NewTextDetector()
	var result, err = detector.DetectBest(b)
	if err != nil {
		//kb todo - graceful error handling here? This will kill file parsing as is
		panic(err)
	}
	return result
}

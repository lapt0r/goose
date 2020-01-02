package loader

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gogs/chardet"
)

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
	if charSet == nil {
		//parsing failure.  Return false
		return false
	}
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
		fmt.Printf("Error getting character bytes: %v", err)
	}
	return result
}

//stream abstraction for filesystem ScanTarget
func getFileSystemBytes(path string) []byte {
	bytes, err := ioutil.ReadFile(path)
	if err == nil {
		return bytes
	}
	return nil
}

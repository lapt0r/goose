package loader

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

	"github.com/gogs/chardet"
)

//ValidateContent ensures that the contents of the text are valid binary encodings
func ValidateContent(path string) bool {
	gitpath, _ := regexp.MatchString("\\.git", path)
	if gitpath {
		return false
	}
	var file, err = os.Open(path)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	var headerBytes = make([]byte, 512)
	var bytesRead, headerReadError = file.Read(headerBytes)

	// empty files are valid
	if headerReadError == io.EOF {
		return true
	}

	if headerReadError != nil {
		fmt.Printf("Error reading file [%v] : %v", path, headerReadError)
	}

	//first, do http content type check
	contentType := http.DetectContentType(headerBytes[0:bytesRead])
	isImage, _ := regexp.MatchString("img", contentType)
	if contentType == "application/octet-stream" || isImage {
		return false
	}
	//charset detection for Go is pretty bad currently, until a port of UDE is available this will have to do
	return true
}

func getByteCharset(b []byte) *chardet.Result {
	var detector = chardet.NewTextDetector()
	//note : detector will return an error if a character set is not detected.  Unknown charsets should be ignored
	var result, _ = detector.DetectBest(b)
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

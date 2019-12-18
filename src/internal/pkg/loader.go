package loader

import
(
	"github.com/gogs/chardet"
	"os"
	"path/filepath"
)

func GetTargets(parent string) []string {
	var files []string
	err := filepath.Walk(parent, func(path string, info os.FileInfo, err error) error {
		//todo: validate file format
        files = append(files, path)
        return nil
    })
    if err != nil {
        panic(err)
	}
	return files
}

func ValidateContent(path string) bool {
	var file = open(path, 'r')
	defer close(file)

}

func GetByteCharset(b []byte) chardet.Result {
	var detector = chardet.NewTextDetector()
	var result, err = detector.DetectBest(b)
	if err != nil {
		panic(err)
	}
}
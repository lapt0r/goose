package loader
import(
	"testing"
	"crypto/rand"
	"fmt"
)

func TestGetByteCharsetValidASCIIString(t *testing.T) {
	var testbytes = []byte("My Super Cool Test String")
	var testResult = GetByteCharset(testbytes)
	if testResult.Charset != "ISO-8859-1" {
		t.Errorf("Chardet mismatch.  Expected ASCII but got %v", testResult.Charset)
	}
}

func TestGetByteCharSetRandomBytes(t *testing.T) {
	var b = make([]byte,100)
	var _, err = rand.Read(b)
	if err != nil {
		panic(err)
	}
	var testResult = GetByteCharset(b)
	if testResult.Charset == "ISO-8859-1" {
		t.Errorf("Chardet mismatch.  Expected unknown charset but got %v", testResult.Charset)
	}
}
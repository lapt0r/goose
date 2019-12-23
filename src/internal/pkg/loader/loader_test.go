package loader

import (
	"crypto/rand"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var testbytes = []byte("My Super Cool Test String")
var longtestbytes = []byte(`
import system;

namespace HelloWorld
{
    public static class HelloWorld
    {
        public static void Main(string[] args){
            string mysupersecretpassword = "hacktheplanet1337!";
            Console.WriteLine("My super secret password is {0}", mysupersecretpassword);
            string my_database_secret = "nobodywillfindthissupersecretvalue!";
        }
    }
}
`)

func TestGetByteCharsetValidASCIIString(t *testing.T) {
	var testResult = getByteCharset(testbytes)
	if testResult.Charset != "ISO-8859-1" {
		t.Errorf("Chardet mismatch.  Expected ISO-8859-1 but got %v", testResult.Charset)
	}
}

func TestGetByteCharsetValidLongString(t *testing.T) {
	var testResult = getByteCharset(longtestbytes)
	if testResult.Charset != "ISO-8859-1" {
		t.Errorf("Chardet mismatch.  Expected ISO-8859-1 but got %v", testResult.Charset)
	}
}

func TestLoaderFileLongString(t *testing.T) {
	file, err := ioutil.TempFile("", "LoaderTestFileLongString")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())
	file.Write(longtestbytes)
	var testResult = ValidateContent(file.Name())
	if !testResult {
		t.Errorf("Expected result to be true but was %v", testResult)
	}
}

func TestGetByteCharSetRandomBytes(t *testing.T) {
	var b = make([]byte, 100)
	var _, err = rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	var testResult = getByteCharset(b)
	if testResult.Charset == "ISO-8859-1" {
		t.Errorf("Chardet mismatch.  Expected unknown charset but got %v", testResult.Charset)
	}
}

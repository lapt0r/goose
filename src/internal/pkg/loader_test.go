package loader

func TestGetByteCharset(t *testing.T) {
	var testbytes = []byte("My Super Cool Test String")
	var testResult = loader.GetByteCharset(testbytes)
	if testResult.Charset != "ASCII" {
		t.Errorf("Chardet mismatch.  Expected ASCII but got %v", testResult.Charset)
	}
	return
}
package streams

import (
	"fmt"
	"os"
	"testing"
)

var testUrl string = "https://raw.githubusercontent.com/hellgate75/general_utils/master/test/resources/sample.txt?token=AAG2EEJA7T2LYJKQ6WVHIOS5ALJFS"
var testfile string = fmt.Sprintf("%s%c%s%c%s%c%s", "..", os.PathSeparator, "test", os.PathSeparator, "resources", os.PathSeparator, "test.txt.yml")

func TestDownloadFile(t *testing.T) {
	err := DownloadFile(testfile, testUrl)
	if err != nil {
		t.Fatal(fmt.Sprintf("DownloadFile::Error Downloading file : %s", err.Error()))
	}

	if _, ferr := os.Stat(testfile); ferr != nil {
		t.Fatal(fmt.Sprintf("DownloadFile::Error Saving Downloaded file : %s", ferr.Error()))
	}
	os.RemoveAll(testfile)
}

func TestDownloadFileAsByteArray(t *testing.T) {
	bytes, err := DownloadFileAsByteArray(testUrl)
	if err != nil {
		t.Fatal(fmt.Sprintf("TestDownloadFileAsByteArray::Error Downloading file : %s", err.Error()))
	}

	if len(bytes) != 11 {
		t.Fatal(fmt.Sprintf("TestDownloadFileAsByteArray::Error Saving Downloaded file - size : %d", len(bytes)))
	}
}

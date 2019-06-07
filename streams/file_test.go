package streams

import (
	"fmt"
	"os"
	"testing"
)

var testLoadFile string = fmt.Sprintf("%s%c%s%c%s%c%s", "..", os.PathSeparator, "test", os.PathSeparator, "resources", os.PathSeparator, "sample.txt")

func TestGetCurrentPath(t *testing.T) {
	path := GetCurrentPath()
	pwd, err := os.Getwd()
	if err != nil {
		t.Fatal(fmt.Sprintf("TestGetCurrentPath::error Error recovering current path. Error : <%s>!!", err.Error()))
	}
	if path != pwd {
		t.Fatal(fmt.Sprintf("TestGetCurrentPath::error Wrong path. Expected <%s> Given <%s>!!", pwd, path))
	}
}

func TestLoadFileContent(t *testing.T) {
	content, err := LoadFileContent(testLoadFile)
	var expected string = "Sample file"
	if err != nil {
		t.Fatal(fmt.Sprintf("TestLoadFileContent::error Error recovering content from file %s. Error : <%s>!!", testLoadFile, err.Error()))
	}
	if content != expected {
		t.Fatal(fmt.Sprintf("TestLoadFileContent::error Error recovering content from file %s. Expected <%s> Given <%s>!!", testLoadFile, expected, content))
	}
}

func TestLoadFileBytes(t *testing.T) {
	bytes, err := LoadFileBytes(testLoadFile)
	if err != nil {
		t.Fatal(fmt.Sprintf("TestLoadFileContent::error Error recovering content from file %s. Error : <%s>!!", testLoadFile, err.Error()))
	}
	if len(bytes) != 11 {
		t.Fatal(fmt.Sprintf("TestLoadFileContent::error Error recovering content from file %s!!", testLoadFile))
	}
}

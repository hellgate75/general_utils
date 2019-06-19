package parsers

import (
	"fmt"
	"os"
	"testing"
)

var testFolderBase64 = fmt.Sprintf("%s%c%s%c%s%c%s", "..", os.PathSeparator, "test", os.PathSeparator, "sample", os.PathSeparator, "base64")
var testFileBase64 = fmt.Sprintf("%s%c%s", testFolderBase64, os.PathSeparator, "sample_parser_file.b64")

type base64SampleStruct struct {
	Name    string
	Surname string
	Age     int
}

var base64SampleData base64SampleStruct = base64SampleStruct{
	"Fabrizio",
	"Torelli",
	44,
}

var base64Parser Parser
var base64ParserErr error
var base64ParserInited bool = false

func initBase64Parser() {
	if base64ParserInited {
		return
	}
	base64Parser, base64ParserErr = New(BASE64)
	base64ParserInited = true
}

func TestBase64SerializeToFileAndDeserializeFromFile(t *testing.T) {
	initBase64Parser()
	if base64ParserErr != nil {
		t.Fatal(fmt.Sprintf("Parser Creation - Arisen unexpected error : %s", base64ParserErr.Error()))
	}
	os.MkdirAll(testFolderBase64, 666)
	defer func(testFolder string) {
		//os.RemoveAll(testFolder)
	}(testFolderBase64)
	err := base64Parser.SerializeToFile(testFileBase64, base64SampleData)
	if err != nil {
		t.Fatal(fmt.Sprintf("Parsing to File - Arisen unexpected error : %s", err.Error()))
	}
	_, err = os.Stat(testFileBase64)
	if err != nil {
		t.Fatal(fmt.Sprintf("Serializing of File - Arisen unexpected error : %s", err.Error()))
	}
	value := base64SampleStruct{}
	err = base64Parser.DeserializeFromFile(testFileBase64, &value)
	if err != nil {
		t.Fatal(fmt.Sprintf("Parsing from File - Arisen unexpected error : %s", err.Error()))
	}
	if value.Name != base64SampleData.Name || value.Surname != base64SampleData.Surname ||
		value.Age != base64SampleData.Age {
		t.Fatal(fmt.Sprintf("TestBase64SerializeToFileAndDeserializeFromFile::error : Expected <%v> but Given <%v>", base64SampleData, value))
	}
}

func TestBase64SerializeToByesAndDeserializeFromBytes(t *testing.T) {
	initBase64Parser()
	if base64ParserErr != nil {
		t.Fatal(fmt.Sprintf("Parser Creation - Arisen unexpected error : %s", base64ParserErr.Error()))
	}
	os.MkdirAll(testFolderBase64, 666)
	byteArray, err := base64Parser.SerializeToBytes(base64SampleData)
	if err != nil {
		t.Fatal(fmt.Sprintf("Parsing to Bytes - Arisen unexpected error : %s", err.Error()))
	}
	if len(byteArray) == 0 {
		t.Fatal("Parsing to Bytes - Found empty buffer")
	}
	value := base64SampleStruct{}
	err = base64Parser.DeserializeFromBytes(byteArray, &value)
	if err != nil {
		t.Fatal(fmt.Sprintf("Parsing from Bytes - Arisen unexpected error : %s", err.Error()))
	}
	if value.Name != base64SampleData.Name || value.Surname != base64SampleData.Surname ||
		value.Age != base64SampleData.Age {
		t.Fatal(fmt.Sprintf("TestBase64SerializeToByesAndDeserializeFromBytes::error : Expected <%v> but Given <%v>", base64SampleData, value))
	}
}

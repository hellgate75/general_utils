package parsers

import (
	"fmt"
	"os"
	"testing"
)

var testFolderGolang = fmt.Sprintf("%s%c%s%c%s%c%s", "..", os.PathSeparator, "test", os.PathSeparator, "sample", os.PathSeparator, "golang")
var testFileGolang = fmt.Sprintf("%s%c%s", testFolderGolang, os.PathSeparator, "sample_parser_file.golang")

type golangSampleStruct struct {
	Name    string `golang:"name,omitempty"`
	Surname string `golang:"surname,omitempty"`
	Age     int    `golang:"age,omitempty"`
}

var golangSampleData golangSampleStruct = golangSampleStruct{
	"Fabrizio",
	"Torelli",
	44,
}

var golangParser Parser
var golangParserErr error
var golangParserInited bool = false

func initGolangParser() {
	if golangParserInited {
		return
	}
	golangParser, golangParserErr = New(GOLANG)
	golangParserInited = true
}

func TestGolangSerializeToFileAndDeserializeFromFile(t *testing.T) {
	initGolangParser()
	if golangParserErr != nil {
		t.Fatal(fmt.Sprintf("Parser Creation - Arisen unexpected error : %s", golangParserErr.Error()))
	}
	os.MkdirAll(testFolderGolang, 0777)
	defer func(testFolder string) {
		os.RemoveAll(testFolder)
	}(testFolderGolang)
	err := golangParser.SerializeToFile(testFileGolang, golangSampleData)
	if err != nil {
		t.Fatal(fmt.Sprintf("Parsing to File - Arisen unexpected error : %s", err.Error()))
	}
	_, err = os.Stat(testFileGolang)
	if err != nil {
		t.Fatal(fmt.Sprintf("Serializing of File - Arisen unexpected error : %s", err.Error()))
	}
	value := golangSampleStruct{}
	err = golangParser.DeserializeFromFile(testFileGolang, &value)
	if err != nil {
		t.Fatal(fmt.Sprintf("Parsing from File - Arisen unexpected error : %s", err.Error()))
	}
	if value.Name != golangSampleData.Name || value.Surname != golangSampleData.Surname ||
		value.Age != golangSampleData.Age {
		t.Fatal(fmt.Sprintf("TestGolangSerializeToFileAndDeserializeFromFile::error : Expected <%v> but Given <%v>", golangSampleData, value))
	}
}

func TestGolangSerializeToByesAndDeserializeFromBytes(t *testing.T) {
	initGolangParser()
	if golangParserErr != nil {
		t.Fatal(fmt.Sprintf("Parser Creation - Arisen unexpected error : %s", golangParserErr.Error()))
	}
	os.MkdirAll(testFolderGolang, 0777)
	defer func(testFolder string) {
		os.RemoveAll(testFolder)
	}(testFolderGolang)
	bytes, err := golangParser.SerializeToBytes(golangSampleData)
	if err != nil {
		t.Fatal(fmt.Sprintf("Parsing to Bytes - Arisen unexpected error : %s", err.Error()))
	}
	if len(bytes) == 0 {
		t.Fatal("Parsing to Bytes - Found emty buffer")
	}
	value := golangSampleStruct{}
	err = golangParser.DeserializeFromBytes(bytes, &value)
	if err != nil {
		t.Fatal(fmt.Sprintf("Parsing from File - Arisen unexpected error : %s", err.Error()))
	}
	if value.Name != golangSampleData.Name || value.Surname != golangSampleData.Surname ||
		value.Age != golangSampleData.Age {
		t.Fatal(fmt.Sprintf("TestGolangSerializeToByesAndDeserializeFromBytes::error : Expected <%v> but Given <%v>", golangSampleData, value))
	}
}

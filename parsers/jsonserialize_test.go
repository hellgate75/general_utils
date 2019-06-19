package parsers

import (
	"fmt"
	"os"
	"testing"
)

var testFolderJson = fmt.Sprintf("%s%c%s%c%s%c%s", "..", os.PathSeparator, "test", os.PathSeparator, "sample", os.PathSeparator, "json")
var testFileJson = fmt.Sprintf("%s%c%s", testFolderJson, os.PathSeparator, "sample_parser_file.json")

type jsonSampleStruct struct {
	Name    string `json:"name,omitempty"`
	Surname string `json:"surname,omitempty"`
	Age     int    `json:"age,omitempty"`
}

var jsonSampleData jsonSampleStruct = jsonSampleStruct{
	"Fabrizio",
	"Torelli",
	44,
}

var jsonParser Parser
var jsonParserErr error
var jsonParserInited bool = false

func initJsonParser() {
	if jsonParserInited {
		return
	}
	jsonParser, jsonParserErr = New(JSON)
	jsonParserInited = true
}

func TestJsonSerializeToFileAndDeserializeFromFile(t *testing.T) {
	initJsonParser()
	if jsonParserErr != nil {
		t.Fatal(fmt.Sprintf("Parser Creation - Arisen unexpected error : %s", jsonParserErr.Error()))
	}
	os.MkdirAll(testFolderJson, 666)
	defer func(testFolder string) {
		os.RemoveAll(testFolder)
	}(testFolderJson)
	err := jsonParser.SerializeToFile(testFileJson, jsonSampleData)
	if err != nil {
		t.Fatal(fmt.Sprintf("Parsing to File - Arisen unexpected error : %s", err.Error()))
	}
	_, err = os.Stat(testFileJson)
	if err != nil {
		t.Fatal(fmt.Sprintf("Serializing of File - Arisen unexpected error : %s", err.Error()))
	}
	value := jsonSampleStruct{}
	err = jsonParser.DeserializeFromFile(testFileJson, &value)
	if err != nil {
		t.Fatal(fmt.Sprintf("Parsing from File - Arisen unexpected error : %s", err.Error()))
	}
	if value.Name != jsonSampleData.Name || value.Surname != jsonSampleData.Surname ||
		value.Age != jsonSampleData.Age {
		t.Fatal(fmt.Sprintf("TestJsonSerializeToFileAndDeserializeFromFile::error : Expected <%v> but Given <%v>", jsonSampleData, value))
	}
}

func TestJsonSerializeToByesAndDeserializeFromBytes(t *testing.T) {
	initJsonParser()
	if jsonParserErr != nil {
		t.Fatal(fmt.Sprintf("Parser Creation - Arisen unexpected error : %s", jsonParserErr.Error()))
	}
	os.MkdirAll(testFolderJson, 666)
	defer func(testFolder string) {
		os.RemoveAll(testFolder)
	}(testFolderJson)
	bytes, err := jsonParser.SerializeToBytes(jsonSampleData)
	if err != nil {
		t.Fatal(fmt.Sprintf("Parsing to Bytes - Arisen unexpected error : %s", err.Error()))
	}
	if len(bytes) == 0 {
		t.Fatal("Parsing to Bytes - Found emty buffer")
	}
	value := jsonSampleStruct{}
	err = jsonParser.DeserializeFromBytes(bytes, &value)
	if err != nil {
		t.Fatal(fmt.Sprintf("Parsing from File - Arisen unexpected error : %s", err.Error()))
	}
	if value.Name != jsonSampleData.Name || value.Surname != jsonSampleData.Surname ||
		value.Age != jsonSampleData.Age {
		t.Fatal(fmt.Sprintf("TestJsonSerializeToByesAndDeserializeFromBytes::error : Expected <%v> but Given <%v>", jsonSampleData, value))
	}
}

package parsers

import (
	"fmt"
	"os"
	"testing"
)

var testFolderYaml = fmt.Sprintf("%s%c%s%c%s%c%s", "..", os.PathSeparator, "test", os.PathSeparator, "sample", os.PathSeparator, "yaml")
var testFileYaml = fmt.Sprintf("%s%c%s", testFolderYaml, os.PathSeparator, "sample_parser_file.yaml")

type yamlSampleStruct struct {
	Name    string `yaml:"name,omitempty"`
	Surname string `yaml:"surname,omitempty"`
	Age     int    `yaml:"age,omitempty"`
}

var yamlSampleData yamlSampleStruct = yamlSampleStruct{
	"Fabrizio",
	"Torelli",
	44,
}

var yamlParser Parser
var yamlParserErr error
var yamlParserInited bool = false

func initYamlParser() {
	if yamlParserInited {
		return
	}
	yamlParser, yamlParserErr = New(YAML)
	yamlParserInited = true
}

func TestYamlSerializeToFileAndDeserializeFromFile(t *testing.T) {
	initYamlParser()
	if yamlParserErr != nil {
		t.Fatal(fmt.Sprintf("Parser Creation - Arisen unexpected error : %s", yamlParserErr.Error()))
	}
	os.MkdirAll(testFolderYaml, 0777)
	defer func(testFolder string) {
		os.RemoveAll(testFolder)
	}(testFolderYaml)
	err := yamlParser.SerializeToFile(testFileYaml, yamlSampleData)
	if err != nil {
		t.Fatal(fmt.Sprintf("Parsing to File - Arisen unexpected error : %s", err.Error()))
	}
	_, err = os.Stat(testFileYaml)
	if err != nil {
		t.Fatal(fmt.Sprintf("Serializing of File - Arisen unexpected error : %s", err.Error()))
	}
	value := yamlSampleStruct{}
	err = yamlParser.DeserializeFromFile(testFileYaml, &value)
	if err != nil {
		t.Fatal(fmt.Sprintf("Parsing from File - Arisen unexpected error : %s", err.Error()))
	}
	if value.Name != yamlSampleData.Name || value.Surname != yamlSampleData.Surname ||
		value.Age != yamlSampleData.Age {
		t.Fatal(fmt.Sprintf("TestYamlSerializeToFileAndDeserializeFromFile::error : Expected <%v> but Given <%v>", yamlSampleData, value))
	}
}

func TestYamlSerializeToByesAndDeserializeFromBytes(t *testing.T) {
	initYamlParser()
	if yamlParserErr != nil {
		t.Fatal(fmt.Sprintf("Parser Creation - Arisen unexpected error : %s", yamlParserErr.Error()))
	}
	os.MkdirAll(testFolderYaml, 0777)
	defer func(testFolder string) {
		os.RemoveAll(testFolder)
	}(testFolderYaml)
	bytes, err := yamlParser.SerializeToBytes(yamlSampleData)
	if err != nil {
		t.Fatal(fmt.Sprintf("Parsing to Bytes - Arisen unexpected error : %s", err.Error()))
	}
	if len(bytes) == 0 {
		t.Fatal("Parsing to Bytes - Found emty buffer")
	}
	value := yamlSampleStruct{}
	err = yamlParser.DeserializeFromBytes(bytes, &value)
	if err != nil {
		t.Fatal(fmt.Sprintf("Parsing from File - Arisen unexpected error : %s", err.Error()))
	}
	if value.Name != yamlSampleData.Name || value.Surname != yamlSampleData.Surname ||
		value.Age != yamlSampleData.Age {
		t.Fatal(fmt.Sprintf("TestYamlSerializeToByesAndDeserializeFromBytes::error : Expected <%v> but Given <%v>", yamlSampleData, value))
	}
}

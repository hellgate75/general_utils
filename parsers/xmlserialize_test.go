package parsers

import (
	"fmt"
	"os"
	"testing"
)

var testFolderXml = fmt.Sprintf("%s%c%s%c%s%c%s", "..", os.PathSeparator, "test", os.PathSeparator, "sample", os.PathSeparator, "xml")
var testFileXml = fmt.Sprintf("%s%c%s", testFolderXml, os.PathSeparator, "sample_parser_file.xml")

type xmlSampleStruct struct {
	Name    string `xml:"name,omitempty"`
	Surname string `xml:"surname,omitempty"`
	Age     int    `xml:"age,omitempty"`
}

var xmlSampleData xmlSampleStruct = xmlSampleStruct{
	"Fabrizio",
	"Torelli",
	44,
}

var xmlParser Parser
var xmlParserErr error
var xmlParserInited bool = false

func initXmlParser() {
	if xmlParserInited {
		return
	}
	xmlParser, xmlParserErr = New(XML)
	xmlParserInited = true
}

func TestXmlSerializeToFileAndDeserializeFromFile(t *testing.T) {
	initXmlParser()
	if xmlParserErr != nil {
		t.Fatal(fmt.Sprintf("Parser Creation - Arisen unexpected error : %s", xmlParserErr.Error()))
	}
	os.MkdirAll(testFolderXml, 0777)
	defer func(testFolder string) {
		os.RemoveAll(testFolder)
	}(testFolderXml)
	err := xmlParser.SerializeToFile(testFileXml, xmlSampleData)
	if err != nil {
		t.Fatal(fmt.Sprintf("Parsing to File - Arisen unexpected error : %s", err.Error()))
	}
	_, err = os.Stat(testFileXml)
	if err != nil {
		t.Fatal(fmt.Sprintf("Serializing of File - Arisen unexpected error : %s", err.Error()))
	}
	value := xmlSampleStruct{}
	err = xmlParser.DeserializeFromFile(testFileXml, &value)
	if err != nil {
		t.Fatal(fmt.Sprintf("Parsing from File - Arisen unexpected error : %s", err.Error()))
	}
	if value.Name != xmlSampleData.Name || value.Surname != xmlSampleData.Surname ||
		value.Age != xmlSampleData.Age {
		t.Fatal(fmt.Sprintf("TestXmlSerializeToFileAndDeserializeFromFile::error : Expected <%v> but Given <%v>", xmlSampleData, value))
	}
}

func TestXmlSerializeToByesAndDeserializeFromBytes(t *testing.T) {
	initXmlParser()
	if xmlParserErr != nil {
		t.Fatal(fmt.Sprintf("Parser Creation - Arisen unexpected error : %s", xmlParserErr.Error()))
	}
	os.MkdirAll(testFolderXml, 0777)
	defer func(testFolder string) {
		os.RemoveAll(testFolder)
	}(testFolderXml)
	bytes, err := xmlParser.SerializeToBytes(xmlSampleData)
	if err != nil {
		t.Fatal(fmt.Sprintf("Parsing to Bytes - Arisen unexpected error : %s", err.Error()))
	}
	if len(bytes) == 0 {
		t.Fatal("Parsing to Bytes - Found emty buffer")
	}
	value := xmlSampleStruct{}
	err = xmlParser.DeserializeFromBytes(bytes, &value)
	if err != nil {
		t.Fatal(fmt.Sprintf("Parsing from File - Arisen unexpected error : %s", err.Error()))
	}
	if value.Name != xmlSampleData.Name || value.Surname != xmlSampleData.Surname ||
		value.Age != xmlSampleData.Age {
		t.Fatal(fmt.Sprintf("TestXmlSerializeToByesAndDeserializeFromBytes::error : Expected <%v> but Given <%v>", xmlSampleData, value))
	}
}

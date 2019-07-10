package parsers

import (
	"fmt"
	"os"
	"testing"
)

var testFolderBinary = fmt.Sprintf("%s%c%s%c%s%c%s", "..", os.PathSeparator, "test", os.PathSeparator, "sample", os.PathSeparator, "binary")
var testFileBinary = fmt.Sprintf("%s%c%s", testFolderBinary, os.PathSeparator, "sample_parser_file.bin")

type binarySampleStruct struct {
	Id     int64
	Type   int
	Lenght int64
	Rate   float64
}

var binarySampleData binarySampleStruct = binarySampleStruct{
	11,
	45,
	1009987,
	7.65,
}

var binaryParser Parser
var binaryParserErr error
var binaryParserInited bool = false

func initBinaryParser() {
	if binaryParserInited {
		return
	}
	binaryParser, binaryParserErr = New(YAML)
	binaryParserInited = true
}

func TestBinarySerializeToFileAndDeserializeFromFile(t *testing.T) {
	initBinaryParser()
	if binaryParserErr != nil {
		t.Fatal(fmt.Sprintf("Parser Creation - Arisen unexpected error : %s", binaryParserErr.Error()))
	}
	os.MkdirAll(testFolderBinary, 0777)
	defer func(testFolder string) {
		os.RemoveAll(testFolder)
	}(testFolderBinary)
	err := binaryParser.SerializeToFile(testFileBinary, binarySampleData)
	if err != nil {
		t.Fatal(fmt.Sprintf("Parsing to File - Arisen unexpected error : %s", err.Error()))
	}
	_, err = os.Stat(testFileBinary)
	if err != nil {
		t.Fatal(fmt.Sprintf("Serializing of File - Arisen unexpected error : %s", err.Error()))
	}
	value := binarySampleStruct{}
	err = binaryParser.DeserializeFromFile(testFileBinary, &value)
	if err != nil {
		t.Fatal(fmt.Sprintf("Parsing from File - Arisen unexpected error : %s", err.Error()))
	}
	if value.Id != binarySampleData.Id || value.Type != binarySampleData.Type ||
		value.Lenght != binarySampleData.Lenght || value.Rate != binarySampleData.Rate {
		t.Fatal(fmt.Sprintf("TestBinarySerializeToFileAndDeserializeFromFile::error : Expected <%v> but Given <%v>", binarySampleData, value))
	}
}

func TestBinarySerializeToByesAndDeserializeFromBytes(t *testing.T) {
	initBinaryParser()
	if binaryParserErr != nil {
		t.Fatal(fmt.Sprintf("Parser Creation - Arisen unexpected error : %s", binaryParserErr.Error()))
	}
	os.MkdirAll(testFolderBinary, 0777)
	defer func(testFolder string) {
		os.RemoveAll(testFolder)
	}(testFolderBinary)
	bytes, err := binaryParser.SerializeToBytes(binarySampleData)
	if err != nil {
		t.Fatal(fmt.Sprintf("Parsing to Bytes - Arisen unexpected error : %s", err.Error()))
	}
	if len(bytes) == 0 {
		t.Fatal("Parsing to Bytes - Found emty buffer")
	}
	value := binarySampleStruct{}
	err = binaryParser.DeserializeFromBytes(bytes, &value)
	if err != nil {
		t.Fatal(fmt.Sprintf("Parsing from File - Arisen unexpected error : %s", err.Error()))
	}
	if value.Id != binarySampleData.Id || value.Type != binarySampleData.Type ||
		value.Lenght != binarySampleData.Lenght || value.Rate != binarySampleData.Rate {
		t.Fatal(fmt.Sprintf("TestBinarySerializeToByesAndDeserializeFromBytes::error : Expected <%v> but Given <%v>", binarySampleData, value))
	}
}

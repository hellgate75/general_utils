package parser

import (
	"fmt"
	"testing"
)

type testInterface struct {
	Text   string
	Number int64
}

type testInterface2 struct {
	Number1 float64
	Number2 int64
}

var testValue testInterface = testInterface{
	Text:   "Test Value",
	Number: 1234,
}

var testValue2 testInterface2 = testInterface2{
	Number1: 123.45,
	Number2: 1234,
}

func arrayEquals(ar1 []byte, ar2 []byte) bool {
	if len(ar1) != len(ar2) {
		return false
	}
	for i := 0; i < len(ar1); i++ {
		if ar1[i] != ar2[i] {
			return false
		}
	}
	return true
}

func TestConvertToPlainText(t *testing.T) {
	var expected []byte = []byte{123, 84, 101, 115, 116, 32, 86, 97, 108, 117, 101, 32, 49, 50, 51, 52, 125}
	bytes, err := _convertToPlainText(testValue)
	if err != nil {
		t.Fatal(fmt.Sprintf("Error converting interface : message -> %s", err.Error()))
	}
	if !arrayEquals(expected, bytes) {
		t.Fatal(fmt.Sprintf("Error converting interface : Expected <%v> Given <%v>", expected, bytes))
	}
}

func TestConvertToJson(t *testing.T) {
	var expected []byte = []byte{123, 34, 84, 101, 120, 116, 34, 58, 34, 84, 101, 115, 116, 32, 86, 97, 108, 117, 101, 34, 44, 34, 78, 117, 109, 98, 101, 114, 34, 58, 49, 50, 51, 52, 125}
	bytes, err := _convertToJson(testValue)
	if err != nil {
		t.Fatal(fmt.Sprintf("Error converting interface : message -> %s", err.Error()))
	}
	if !arrayEquals(expected, bytes) {
		t.Fatal(fmt.Sprintf("Error converting interface : Expected <%v> Given <%v>", expected, bytes))
	}
}

func TestConvertToXml(t *testing.T) {
	var expected []byte = []byte{60, 116, 101, 115, 116, 73, 110, 116, 101, 114, 102, 97, 99, 101, 62, 60, 84, 101, 120, 116, 62, 84, 101, 115, 116, 32, 86, 97, 108, 117, 101, 60, 47, 84, 101, 120, 116, 62, 60, 78, 117, 109, 98, 101, 114, 62, 49, 50, 51, 52, 60, 47, 78, 117, 109, 98, 101, 114, 62, 60, 47, 116, 101, 115, 116, 73, 110, 116, 101, 114, 102, 97, 99, 101, 62}
	bytes, err := _convertToXml(testValue)
	if err != nil {
		t.Fatal(fmt.Sprintf("Error converting interface : message -> %s", err.Error()))
	}
	if !arrayEquals(expected, bytes) {
		t.Fatal(fmt.Sprintf("Error converting interface : Expected <%v> Given <%v>", expected, bytes))
	}
}

func TestConvertToYaml(t *testing.T) {
	var expected []byte = []byte{116, 101, 120, 116, 58, 32, 84, 101, 115, 116, 32, 86, 97, 108, 117, 101, 10, 110, 117, 109, 98, 101, 114, 58, 32, 49, 50, 51, 52, 10}
	bytes, err := _convertToYaml(testValue)
	if err != nil {
		t.Fatal(fmt.Sprintf("Error converting interface : message -> %s", err.Error()))
	}
	if !arrayEquals(expected, bytes) {
		t.Fatal(fmt.Sprintf("Error converting interface : Expected <%v> Given <%v>", expected, bytes))
	}
}

func TestConvertToBinary(t *testing.T) {
	var expected []byte = []byte{64, 94, 220, 204, 204, 204, 204, 205, 0, 0, 0, 0, 0, 0, 4, 210}
	bytes, err := _convertToBinary(testValue2)
	if err != nil {
		t.Fatal(fmt.Sprintf("Error converting interface : message -> %s", err.Error()))
	}
	if !arrayEquals(expected, bytes) {
		t.Fatal(fmt.Sprintf("Error converting interface : Expected <%v> Given <%v>", expected, bytes))
	}
	//	fmt.Println(bytes)
}

func TestConvertToBase64(t *testing.T) {
	var expected []byte = []byte{101, 49, 82, 108, 99, 51, 81, 103, 86, 109, 70, 115, 100, 87, 85, 103, 77, 84, 73, 122}
	bytes, err := _convertToBase64(testValue)
	if err != nil {
		t.Fatal(fmt.Sprintf("Error converting interface : message -> %s", err.Error()))
	}
	if !arrayEquals(expected, bytes) {
		t.Fatal(fmt.Sprintf("Error converting interface : Expected <%v> Given <%v>", expected, bytes))
	}
	//	fmt.Println(bytes)
}

func TestConvertToGoFormat(t *testing.T) {
	var expected []byte = []byte{47, 255, 129, 3, 1, 1, 13, 116, 101, 115, 116, 73, 110, 116, 101, 114, 102, 97, 99, 101, 1, 255, 130, 0, 1, 2, 1, 4, 84, 101, 120, 116, 1, 12, 0, 1, 6, 78, 117, 109, 98, 101, 114, 1, 4, 0, 0, 0, 19, 255, 130, 1, 10, 84, 101, 115, 116, 32, 86, 97, 108, 117, 101, 1, 254, 9, 164, 0}
	bytes, err := _convertToGoFormat(testValue)
	if err != nil {
		t.Fatal(fmt.Sprintf("Error converting interface : message -> %s", err.Error()))
	}
	if !arrayEquals(expected, bytes) {
		t.Fatal(fmt.Sprintf("Error converting interface : Expected <%v> Given <%v>", expected, bytes))
	}
}
func TestNewLocalWriter(t *testing.T) {
	var expected []byte = []byte{47, 255, 129, 3, 1, 1, 13, 116, 101, 115, 116, 73, 110, 116, 101, 114, 102, 97, 99, 101, 1, 255, 130, 0, 1, 2, 1, 4, 84, 101, 120, 116, 1, 12, 0, 1, 6, 78, 117, 109, 98, 101, 114, 1, 4, 0, 0, 0, 19, 255, 130, 1, 10, 84, 101, 115, 116, 32, 86, 97, 108, 117, 101, 1, 254, 9, 164, 0}
	writer := NewLocalWriter()
	num, err := writer.Write(expected)
	if err != nil {
		t.Fatal(fmt.Sprintf("Error writing bytes : message -> %s", err.Error()))
	}
	if num == 0 {
		t.Fatal(fmt.Sprintf("Error writing bytes : No bytes written"))
	}
	bytes, err2 := writer.GetBytes()
	if err2 != nil {
		t.Fatal(fmt.Sprintf("Error reading bytes : message -> %s", err2.Error()))
	}
	if !arrayEquals(expected, bytes) {
		t.Fatal(fmt.Sprintf("Error converting interface : Expected <%v> Given <%v>", expected, bytes))
	}
}

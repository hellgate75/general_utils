package parsers

import (
	"fmt"
	"testing"
)

func TestEncodingFromString(t *testing.T) {
	enc, err := EncodingFromString("json")
	if err != nil {
		t.Fatal(fmt.Sprintf("Arisen unexpected error : %s", err.Error()))
	}
	if enc != JSON {
		t.Fatal(fmt.Sprintf("JSON - Wrong result in EncodingFromString, Expected <%v> but Given <%v>", JSON, enc))
	}
	enc, err = EncodingFromString("xml")
	if err != nil {
		t.Fatal(fmt.Sprintf("Arisen unexpected error : %s", err.Error()))
	}
	if enc != XML {
		t.Fatal(fmt.Sprintf("XML - Wrong result in EncodingFromString, Expected <%v> but Given <%v>", XML, enc))
	}
	enc, err = EncodingFromString("yaml")
	if err != nil {
		t.Fatal(fmt.Sprintf("Arisen unexpected error : %s", err.Error()))
	}
	if enc != YAML {
		t.Fatal(fmt.Sprintf("YAML - Wrong result in EncodingFromString, Expected <%v> but Given <%v>", YAML, enc))
	}
	enc, err = EncodingFromString("gof")
	if err != nil {
		t.Fatal(fmt.Sprintf("Arisen unexpected error : %s", err.Error()))
	}
	if enc != GOLANG {
		t.Fatal(fmt.Sprintf("YAML - Wrong result in EncodingFromString, Expected <%v> but Given <%v>", GOLANG, enc))
	}
	enc, err = EncodingFromString("B64")
	if err != nil {
		t.Fatal(fmt.Sprintf("Arisen unexpected error : %s", err.Error()))
	}
	if enc != BASE64 {
		t.Fatal(fmt.Sprintf("YAML - Wrong result in EncodingFromString, Expected <%v> but Given <%v>", BASE64, enc))
	}
	enc, err = EncodingFromString("bin")
	if err != nil {
		t.Fatal(fmt.Sprintf("Arisen unexpected error : %s", err.Error()))
	}
	if enc != BINARY {
		t.Fatal(fmt.Sprintf("YAML - Wrong result in EncodingFromString, Expected <%v> but Given <%v>", BINARY, enc))
	}
	enc, err = EncodingFromString("i-dunno")
	if err != nil {
		t.Fatal(fmt.Sprintf("Arisen unexpected error : %s", err.Error()))
	}
	if enc != JSON {
		t.Fatal(fmt.Sprintf("UNKNOWN - Wrong result in EncodingFromString, Expected <%v> but Given <%v>", JSON, enc))
	}
	enc, err = EncodingFromString("   ")
	if err == nil {
		t.Fatal("Expected error but not arisen!!")
	}
}

func TestNewParser(t *testing.T) {
	parser, err := New(JSON)
	if err != nil {
		t.Fatal(fmt.Sprintf("Arisen unexpected error : %s", err.Error()))
	}
	if parser.GetEncoding() != JSON {
		t.Fatal(fmt.Sprintf("JSON - Wrong result in New Parser, Expected <%v> but Given <%v>", JSON, parser.GetEncoding()))
	}
	parser, err = New(XML)
	if err != nil {
		t.Fatal(fmt.Sprintf("Arisen unexpected error : %s", err.Error()))
	}
	if parser.GetEncoding() != XML {
		t.Fatal(fmt.Sprintf("XML - Wrong result in New Parser, Expected <%v> but Given <%v>", XML, parser.GetEncoding()))
	}
	parser, err = New(YAML)
	if err != nil {
		t.Fatal(fmt.Sprintf("Arisen unexpected error : %s", err.Error()))
	}
	if parser.GetEncoding() != YAML {
		t.Fatal(fmt.Sprintf("YAML - Wrong result in New Parser, Expected <%v> but Given <%v>", YAML, parser.GetEncoding()))
	}
	parser, err = New(GOLANG)
	if err != nil {
		t.Fatal(fmt.Sprintf("Arisen unexpected error : %s", err.Error()))
	}
	if parser.GetEncoding() != GOLANG {
		t.Fatal(fmt.Sprintf("YAML - Wrong result in New Parser, Expected <%v> but Given <%v>", GOLANG, parser.GetEncoding()))
	}
	parser, err = New(BASE64)
	if err != nil {
		t.Fatal(fmt.Sprintf("Arisen unexpected error : %s", err.Error()))
	}
	if parser.GetEncoding() != BASE64 {
		t.Fatal(fmt.Sprintf("YAML - Wrong result in New Parser, Expected <%v> but Given <%v>", BASE64, parser.GetEncoding()))
	}
	parser, err = New(BINARY)
	if err != nil {
		t.Fatal(fmt.Sprintf("Arisen unexpected error : %s", err.Error()))
	}
	if parser.GetEncoding() != BINARY {
		t.Fatal(fmt.Sprintf("YAML - Wrong result in New Parser, Expected <%v> but Given <%v>", BINARY, parser.GetEncoding()))
	}
	parser, err = New(Encoding(0))
	if err == nil {
		t.Fatal("Expected error but not arisen!!")
	}
}

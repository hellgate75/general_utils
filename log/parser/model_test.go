package parser

import (
	"fmt"
	"github.com/hellgate75/general_utils/common"
	"testing"
)

func TestWriterTypeToLogStreamType(t *testing.T) {
	lst, err := WriterTypeToLogStreamType(common.StdOutWriter)
	if err != nil {
		t.Fatal(fmt.Sprintf("Unable to convert WriterType - error : %s", err.Error()))
	}
	if lst != StdOutStreamType {
		t.Fatal(fmt.Sprintf("Unable to convert WriterType - Expected %v Given %v", StdOutStreamType, lst))
	}
	lst, err = WriterTypeToLogStreamType(common.FileWriter)
	if err != nil {
		t.Fatal(fmt.Sprintf("Unable to convert WriterType - error : %s", err.Error()))
	}
	if lst != FileStreamType {
		t.Fatal(fmt.Sprintf("Unable to convert WriterType - Expected %v Given %v", FileStreamType, lst))
	}
	lst, err = WriterTypeToLogStreamType(common.UrlWriter)
	if err != nil {
		t.Fatal(fmt.Sprintf("Unable to convert WriterType - error : %s", err.Error()))
	}
	if lst != UrlStreamType {
		t.Fatal(fmt.Sprintf("Unable to convert WriterType - Expected %v Given %v", UrlStreamType, lst))
	}

}

func TestNewLogStream(t *testing.T) {

}

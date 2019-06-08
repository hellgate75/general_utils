package parser

import (
	"fmt"
	"github.com/hellgate75/general_utils/common"
	"testing"
	"time"
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

func executeStreamTest(logType LogStreamType, format common.StreamInOutFormat, t *testing.T, label string) {
	stdWriter, err1 := New(logType, format)
	if err1 != nil {
		t.Fatal(fmt.Sprintf("Unable to instantiate %s Stream Writer - error : %s", label, err1.Error()))
	}
	err1 = stdWriter.Open()
	if err1 != nil {
		t.Fatal(fmt.Sprintf("Unable to open %s Stream Writer - error : %s", label, err1.Error()))
	}
	var expected interface{} = "Test Value"
	var outcome interface{}
	go func() {
		time.Sleep(250 * time.Millisecond)
		err1 = stdWriter.Write(expected)
		if err1 != nil {
			t.Fatal(fmt.Sprintf("Unable to write into %s Stream Writer - error : %s", label, err1.Error()))
		}
	}()
	select {
	case outcome = <-testChan:
		if outcome != expected {
			t.Fatal(fmt.Sprintf("Unable write into %s Stream Writer - Expected %v Given %v", label, expected, outcome))
		}
	case <-time.After(time.Second):
		t.Fatal(fmt.Sprintf("Timeout waiting write feedback from %s Stream Writer!!", label))
	}

}

func TestNewLogStream(t *testing.T) {
	testFlag = true
	testChan = make(chan interface{})
	executeStreamTest(StdOutStreamType, common.PlainTextFormat, t, "StdOut")
	//TODO: Complete and place tests for File and Url Stream Types
	testFlag = false
	close(testChan)
}

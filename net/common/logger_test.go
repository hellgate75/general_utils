package common

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestLogLevels(t *testing.T) {
	if TRACE != LogLevel(1) {
		t.Fatal(fmt.Sprintf("Wrong Values - Expected: <%v> but Given: <%v>", LogLevel(1), TRACE))
	}
	if DEBUG != LogLevel(2) {
		t.Fatal(fmt.Sprintf("Wrong Values - Expected: <%v> but Given: <%v>", LogLevel(2), DEBUG))
	}
	if INFO != LogLevel(3) {
		t.Fatal(fmt.Sprintf("Wrong Values - Expected: <%v> but Given: <%v>", LogLevel(3), INFO))
	}
	if WARNING != LogLevel(4) {
		t.Fatal(fmt.Sprintf("Wrong Values - Expected: <%v> but Given: <%v>", LogLevel(4), WARNING))
	}
	if ERROR != LogLevel(5) {
		t.Fatal(fmt.Sprintf("Wrong Values - Expected: <%v> but Given: <%v>", LogLevel(5), ERROR))
	}
	if FATAL != LogLevel(6) {
		t.Fatal(fmt.Sprintf("Wrong Values - Expected: <%v> but Given: <%v>", LogLevel(6), FATAL))
	}
	if NO_LOG != LogLevel(0) {
		t.Fatal(fmt.Sprintf("Wrong Values - Expected: <%v> but Given: <%v>", LogLevel(0), NO_LOG))
	}
}

func TestNewServerLogger(t *testing.T) {
	serverLogger := NewServerLogger(DEBUG)
	var outChan chan interface{} = make(chan interface{})
	defer func() {
		close(outChan)
	}()
	serverLogger.Open()
	serverLogger.AddOutChannel(&outChan)
	var message interface{} = "Test Message"
	var incoming interface{} = ""
	var err error
	go func() {
		time.Sleep(1 * time.Second)
		err = serverLogger.Log(INFO, message)
	}()
	select {
	case incoming = <-outChan:
	case <-time.After(5 * time.Second):
		err = errors.New("Time out!!")
	}
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error : %s", err.Error()))
	}
	if message != incoming {
		t.Fatal(fmt.Sprintf("Wrong Values - Expected: <%v> but Given: <%v>", message, incoming))
	}
	serverLogger.Close()
	err = serverLogger.Log(INFO, message)
	if err == nil {
		t.Fatal("Unexpected Nil error message when Logger is Closed and you try to Log!!!")
	}
	err = serverLogger.Open()
	err = serverLogger.Log(INFO, message)
	if err == nil {
		t.Fatal("Unexpected Nil error message when Logger is Closed and you try to Log!!!")
	}
	outChan = make(chan interface{})
	serverLogger.AddOutChannel(&outChan)
	err = serverLogger.Log(TRACE, message)
	if err == nil {
		t.Fatal("Unexpected Nil error message Log Verbosity is Higher!!!")
	}
}

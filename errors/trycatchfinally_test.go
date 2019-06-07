package errors

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestThrow(t *testing.T) {
	var err error
	go func() {
		defer func() {
			msg := recover()
			err = msg.(error)
		}()
		Throw(errors.New("Test error"))
	}()
	time.Sleep(500 * time.Millisecond)
	if err == nil {
		t.Fatal("No error arisen...")
	}
	if err.Error() != "Test error" {
		t.Fatal(fmt.Sprintf("Wrong error arisen from the system. Expected message <%s> Given message <%s>...", "Test error", err.Error()))
	}
}

func TestTry(t *testing.T) {
	var tryOut, catchOut, finallyOut bool
	var err Exception
	var newErr = errors.New("Try Catch Finally Error")
	Try(func(...interface{}) interface{} {
		tryOut = true
		Throw(newErr)
		return nil
	}).Catch(func(e Exception) {
		catchOut = true
		err = e
	}).Finally(func() {
		finallyOut = true
	}).Do()
	time.Sleep(500 * time.Millisecond)
	if !tryOut {
		t.Fatal("Try Function has not been Called!!")
	}
	if !catchOut {
		t.Fatal("Catch Function has not been Called!!")
	}
	if !finallyOut {
		t.Fatal("Finally Function has not been Called!!")
	}
	if err == nil {
		t.Fatal("No error arisen!!")
	}
	if err != newErr {
		t.Fatal(fmt.Sprintf("Wrong error arisen from the system. Expected message <%v> Given message <%v>...", newErr, err))
	}
}

func TestTryReturn(t *testing.T) {
	var tryOut, finallyOut bool
	var expected interface{} = "Try Outcome"
	var value interface{}
	value = Try(func(...interface{}) interface{} {
		tryOut = true
		return expected
	}).Finally(func() {
		finallyOut = true
	}).Do()
	time.Sleep(500 * time.Millisecond)
	if !tryOut {
		t.Fatal("Try Function has not been Called!!")
	}
	if expected != value {
		t.Fatal(fmt.Sprintf("Try Function has returned worg value. Expected <%v> Given <%v>!!", expected, value))
	}
	if !finallyOut {
		t.Fatal("Finally Function has not been Called!!")
	}
}

func TestMultiTry(t *testing.T) {
	var tryOut, catchOut, finallyOut bool
	var err Exception
	var newErr = errors.New("Try Catch Finally Error")
	MultiTry(func(...interface{}) interface{} {
		tryOut = true
		Throw(newErr)
		return nil
	}).Catch(func(e Exception) {
		catchOut = true
		err = e
	}).Finally(func() {
		finallyOut = true
	}).Do()
	time.Sleep(500 * time.Millisecond)
	if !tryOut {
		t.Fatal("Try Function has not been Called!!")
	}
	if !catchOut {
		t.Fatal("Catch Function has not been Called!!")
	}
	if !finallyOut {
		t.Fatal("Finally Function has not been Called!!")
	}
	if err == nil {
		t.Fatal("No error arisen!!")
	}
	if err != newErr {
		t.Fatal(fmt.Sprintf("Wrong error arisen from the system. Expected message <%v> Given message <%v>...", newErr, err))
	}
}

func TestMultiTryReturn(t *testing.T) {
	var tryOut, finallyOut bool
	var expected interface{} = "Try Outcome"
	var value []interface{}
	value = MultiTry(func(...interface{}) interface{} {
		tryOut = true
		return expected
	}).Finally(func() {
		finallyOut = true
	}).Do()
	time.Sleep(500 * time.Millisecond)
	if !tryOut {
		t.Fatal("Try Function has not been Called!!")
	}
	if expected != value[0] {
		t.Fatal(fmt.Sprintf("Try Function has returned worg value. Expected <%v> Given <%v>!!", expected, value[0]))
	}
	if !finallyOut {
		t.Fatal("Finally Function has not been Called!!")
	}
}

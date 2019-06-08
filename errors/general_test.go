package errors

import (
	"errors"
	"fmt"
	"testing"
)

func TestGeneralException(T *testing.T) {
	text := "test exception"
	exc := New(text)
	if exc.Error() != text {
		T.Fatal(fmt.Sprintf("Base Exception instance failed - Expected <%s> but Given <%s>", text, exc.Error()))
	}
}

func TestIsError(T *testing.T) {
	text := "test exception"
	flag := IsError(text)
	if flag {
		T.Fatal(fmt.Sprintf("String is not and error - Expected <%t> but Given <%t>", false, flag))
	}
	err := errors.New("My Custom Error")
	flag = IsError(err)
	if !flag {
		T.Fatal(fmt.Sprintf("Err is and error - Expected <%t> but Given <%t>", true, flag))
	}
}

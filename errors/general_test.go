package errors

import (
	"testing"
)

func TestGeneralException(T *testing.T) {
	text := "test exception"
	exc := New(text)
	if exc.Error() != text {
		T.Fatal("Base Exception instance failed - Expected <" + text + "> but Given <" + exc.Error() + ">")
	}
}

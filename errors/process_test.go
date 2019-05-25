package errors

import (
	"fmt"
	utils "github.com/hellgate75/general_utils/utils"
	"testing"
)

func TestProcessException(T *testing.T) {
	text := "test exception"
	var pid utils.PID = 1234
	var state utils.ProcessState = utils.ERROR

	exc := NewProcess(text, pid, state)
	if exc.Error() != text {
		T.Fatal("Base Exception instance failed - Expected <" + text + "> but Given <" + exc.Error() + ">")
	}

	if exc.Pid() != pid {
		T.Fatal("Base Exception instance failed - Expected <" + fmt.Sprintf("%v", pid) + "> but Given <" + fmt.Sprintf("%v", exc.Pid()) + ">")
	}

	if exc.State() != state {
		T.Fatal("Base Exception instance failed - Expected <" + fmt.Sprintf("%v", state) + "> but Given <" + fmt.Sprintf("%v", exc.State()) + ">")
	}
}

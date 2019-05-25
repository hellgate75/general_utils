package errors

import (
	utils "github.com/hellgate75/general_utils/utils"
)

type processError interface {
	error
	Pid() utils.PID
	State() utils.ProcessState
}

type processErrorStruct struct {
	text  string
	pid   utils.PID
	state utils.ProcessState
}

func (e *processErrorStruct) Error() string {
	return e.text
}

func (e *processErrorStruct) Pid() utils.PID {
	return e.pid
}

func (e *processErrorStruct) State() utils.ProcessState {
	return e.state
}

func NewProcess(text string, pid utils.PID, processState utils.ProcessState) processError {
	return &processErrorStruct{
		text,
		pid,
		processState,
	}
}

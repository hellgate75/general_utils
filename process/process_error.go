package process

import ()

type processError interface {
	error
	Pid() PID
	State() ProcessState
}

type processErrorStruct struct {
	text  string
	pid   PID
	state ProcessState
}

func (e *processErrorStruct) Error() string {
	return e.text
}

func (e *processErrorStruct) Pid() PID {
	return e.pid
}

func (e *processErrorStruct) State() ProcessState {
	return e.state
}

func NewProcess(text string, pid PID, processState ProcessState) processError {
	return &processErrorStruct{
		text:  text,
		pid:   pid,
		state: processState,
	}
}

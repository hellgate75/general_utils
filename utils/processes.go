package utils

import (
//	"fmt"
//	"time"
)

type ProcessState int
type ProcessChain chan struct{}

type ProcessFunction func(*ProcessManager, ProcessChain, ProcessChain)

const (
	DEFAULT ProcessState = iota
	STARTING
	RUNNING
	PAUSED
	RESUMING
	STOP
	DEAD
	ERROR
)

type ProcessManager struct {
	Pid			int64
	InChain		ProcessChain
	OutChain	ProcessChain
	_state		ProcessState
	_progess	bool
	_function	ProcessFunction
	
}

func (m ProcessManager) Start() {
	m._progess = false
	m.startProcess()
}

func (m ProcessManager) Status() bool {
	return m._state == RUNNING
}

func (m ProcessManager) startProcess() {
	defer func() {
		if m._progess {
			m._state = RUNNING
		} else {
			m._state = ERROR
		}
	}	
	m._state = STARTING
}

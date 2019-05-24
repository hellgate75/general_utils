package utils

import (
//	"fmt"
//	"time"
)

type ProcessState int
type ProcessChain chan struct{}

type ProcessFunction func(*ProcessManager, *ProcessChain, *ProcessChain)
type PID int64

var REG_PID PID = 0

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
	Pid       int64
	InChain   ProcessChain
	OutChain  ProcessChain
	_state    ProcessState
	_progess  bool
	_function ProcessFunction
}

func (m ProcessManager) Init(function ProcessFunction) {
	m._function = function
	m.InChain = make(ProcessChain)
	m.OutChain = make(ProcessChain)
}

func (m ProcessManager) Start() {
	if m._function == nil {
		panic("ProcessManager::error : Please provide function to process manager")
	}
	m._progess = false
	m.startProcess()
}

func (m ProcessManager) Status() bool {
	return m._state == RUNNING
}

func (m ProcessManager) startProcess() {
	defer func() {
		if m._progess {
			m._state = STOP
		} else {
			m._state = ERROR
		}
		m._progess = false
	}()
	REG_PID++
	m.Pid = REG_PID
	m._state = STARTING
	m._state = RUNNING
	m._function(&m, &m.InChain, &m.OutChain)
	m._progess = true

}

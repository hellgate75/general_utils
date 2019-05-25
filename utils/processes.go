package utils

import (
	//	"fmt"
	"time"
)

type ProcessState int
type ProcessChannel chan interface{}

type ProcessFunction func(ProcessManager, *ProcessChannel, *ProcessChannel)
type PID int64

var REG_PID PID = 0

const (
	DEFAULT  ProcessState = 1
	STARTING ProcessState = 2
	RUNNING  ProcessState = 3
	PAUSED   ProcessState = 4
	RESUMING ProcessState = 5
	DONE     ProcessState = 6
	ERROR    ProcessState = 7
)

type ProcessManager interface {
	Start()
	Running() bool
	Status() ProcessState
	GetInChannel() *ProcessChannel
	GetOutChannel() *ProcessChannel
}

type processManagerStruct struct {
	Pid       PID
	_inChan   ProcessChannel
	_outChan  ProcessChannel
	_state    ProcessState
	_progess  bool
	_function ProcessFunction
}

func (m *processManagerStruct) setState(s ProcessState) {
	m._state = s
}

func (m *processManagerStruct) init(function ProcessFunction) *processManagerStruct {

	m.setState(DEFAULT)
	m._function = function
	m._inChan = make(ProcessChannel)
	m._outChan = make(ProcessChannel)
	return m
}

func (m *processManagerStruct) GetInChannel() *ProcessChannel {
	return &m._inChan
}

func (m *processManagerStruct) GetOutChannel() *ProcessChannel {
	return &m._outChan
}

func (m *processManagerStruct) Start() {
	if m._function == nil {
		panic("ProcessManager::error : Please call Init before Start of Process Manager")
	}
	m._progess = false
	go m.startProcess()
}

func (m *processManagerStruct) Running() bool {
	return m._state > DEFAULT && m._state < DONE
}

func (m *processManagerStruct) Status() ProcessState {
	return m._state
}

func (m *processManagerStruct) startProcess() {
	//fmt.Println("Init Start ...")
	defer func() {
		time.Sleep(1 * time.Second)
		m.stopProcess()
	}()
	//fmt.Println("Preparing to Start ...")
	REG_PID++
	m.Pid = REG_PID
	m._state = STARTING
	//fmt.Println("Starting ...")
	time.Sleep(1 * time.Second)
	m._state = RUNNING
	m._function(m, &m._inChan, &m._outChan)
	//fmt.Println("Done!!!")
	m._progess = true
}

func (m *processManagerStruct) stopProcess() {
	defer close(m._inChan)
	defer close(m._outChan)
	if m._progess {
		m._state = DONE
	} else {
		m._state = ERROR
	}
	m._progess = false
}

func NewProcessManager(fn ProcessFunction) ProcessManager {
	var pms processManagerStruct = processManagerStruct{
		Pid: 0,
	}
	return pms.init(fn)
}

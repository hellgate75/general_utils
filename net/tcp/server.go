package tcp

import (
	"errors"
	"fmt"
	errs "github.com/hellgate75/general_utils/errors"
	"github.com/hellgate75/general_utils/net/common"
	"net"
	"strconv"
	"sync"
	"time"
)

const (
	CONN_TYPE = "tcp"
)

// Tcp Endpoint handler function
type TcpEndpoint func(net.Conn, *chan interface{}, common.NetContext) error

// Call closure function
type TcpCycleClosureFunc func(net.Conn, common.ServerLogger) error

var __defaultClosureFunc TcpCycleClosureFunc = func(conn net.Conn, logger common.ServerLogger) error {
	err := conn.Close()
	if err != nil {
		logger.Log(common.ERROR, fmt.Sprintf("Error closing the client connection, Error is %s", err.Error()))
	}
	return err
}

// Tcp Endpoint Map type
type TcpEntriesList []TcpEndpoint

type __tcpServerStruct struct {
	EntriesList    TcpEntriesList
	ListeningHost  string
	ListeningPort  common.Port
	ClosureFunc    TcpCycleClosureFunc
	__listener     net.Listener
	__lock         sync.RWMutex
	__globalCache  common.NetCache
	__serviceCache map[string]common.NetCache
	__globalEnv    common.NetEnvironment
	__serviceEnv   map[string]common.NetEnvironment
	__serverPipe   chan interface{}
	__running      bool
	__logger       common.ServerLogger
}

func (tcpServer *__tcpServerStruct) __baseTcpEntryActionFunc(conn net.Conn) {
	var err error
	defer func() {
		itf := recover()
		if errs.IsError(itf) {
			err = itf.(error)
			tcpServer.__logger.Log(common.ERROR, fmt.Sprintf("Error executing tcp server handler : %s", err.Error()))
		} else {
			err = errors.New(fmt.Sprintf("%v", itf))
			tcpServer.__logger.Log(common.ERROR, fmt.Sprintf("Error executing tcp server handler : %v", itf))
		}
	}()
	tcpServer.__lock.RLock()
	var counter int = 0
	for index, tcpEndPoint := range tcpServer.EntriesList {
		counter++
		go func() {
			defer func() {
				recover()
				counter--
			}()
			err = tcpEndPoint(conn, &tcpServer.__serverPipe, tcpServer)
			if err != nil {
				tcpServer.__logger.Log(common.ERROR, fmt.Sprintf("Error executing tcp server handler # %x : %s", index, err.Error()))
			}
		}()
	}
	tcpServer.__lock.RUnlock()

	for counter > 0 {
		time.Sleep(500 * time.Millisecond)
	}

	if tcpServer.ClosureFunc != nil {
		err = tcpServer.ClosureFunc(conn, tcpServer.__logger)
	} else {
		err = conn.Close()
	}
	if err != nil {
		tcpServer.__logger.Log(common.ERROR, fmt.Sprintf("Error executing tcp server : %s", err.Error()))
	}
}

func (tcpServer *__tcpServerStruct) Open() error {
	var err error
	defer func() {
		itf := recover()
		if errs.IsError(itf) {
			err = itf.(error)
			tcpServer.__logger.Log(common.ERROR, fmt.Sprintf("Error executing tcp server Open : %s", err.Error()))
		} else {
			err = errors.New(fmt.Sprintf("%v", itf))
			tcpServer.__logger.Log(common.ERROR, fmt.Sprintf("Error executing tcp server Open : %v", itf))
		}
	}()
	tcpServer.__running = true
	tcpServer.__logger.Open()
	listener, errC := net.Listen(CONN_TYPE, fmt.Sprintf("%s:%v", tcpServer.ListeningHost, tcpServer.ListeningPort))
	if errC != nil {
		return errC
	}
	tcpServer.__listener = listener
	go func() {
		for tcpServer.__running {
			conn, err := tcpServer.__listener.Accept()
			if err != nil {
				tcpServer.__logger.Log(common.ERROR, fmt.Sprintf("Error acquiring server client : %s", err.Error()))
			} else {
				go func() {
					tcpServer.__baseTcpEntryActionFunc(conn)
				}()
			}
		}
	}()
	tcpServer.__serverPipe = make(chan interface{})
	tcpServer.__globalCache = make(common.NetCache)
	tcpServer.__globalEnv = make(common.NetEnvironment)
	if tcpServer.__listener != nil {
		tcpServer.__globalEnv["ServerAddress"] = tcpServer.__listener.Addr()
	}
	tcpServer.__globalEnv["ServerPort"] = tcpServer.ListeningPort
	tcpServer.__globalEnv["ServerHost"] = tcpServer.ListeningHost
	tcpServer.__serviceCache = make(map[string]common.NetCache)
	tcpServer.__serviceEnv = make(map[string]common.NetEnvironment)
	for k, _ := range tcpServer.EntriesList {
		tcpServer.__serviceCache[strconv.Itoa(k)] = make(common.NetCache)
		tcpServer.__serviceEnv[strconv.Itoa(k)] = make(common.NetEnvironment)
	}
	return err
}
func (tcpServer *__tcpServerStruct) Close() error {
	var err error
	defer func() {
		itf := recover()
		if errs.IsError(itf) {
			err = itf.(error)
			tcpServer.__logger.Log(common.ERROR, fmt.Sprintf("Error executing tcp server Close : %s", err.Error()))
		} else {
			err = errors.New(fmt.Sprintf("%v", itf))
			tcpServer.__logger.Log(common.ERROR, fmt.Sprintf("Error executing tcp server Close : %v", itf))
		}
	}()
	tcpServer.__logger.Close()
	tcpServer.__running = false
	if tcpServer.__listener != nil {
		err = tcpServer.__listener.Close()
	}
	close(tcpServer.__serverPipe)
	return err
}
func (tcpServer *__tcpServerStruct) IsListening() bool {
	return tcpServer.__running
}
func (tcpServer *__tcpServerStruct) IsRunning() bool {
	return tcpServer.__running
}
func (tcpServer *__tcpServerStruct) Stream(interface{}) error {
	return errors.New("Not Implemeted!!")

}
func (tcpServer *__tcpServerStruct) Receive() (interface{}, error) {
	return nil, errors.New("Not Implemeted!!")
}
func (tcpServer *__tcpServerStruct) HandleConnOn(hf common.ServerHablerFunc) error {
	return errors.New("Not Implemeted!!")
}
func (tcpServer *__tcpServerStruct) HandlingFuncs() []common.ServerHablerFunc {
	return []common.ServerHablerFunc{}
}
func (tcpServer *__tcpServerStruct) RevokeHandler(hf common.ServerHablerFunc) error {
	return errors.New("Not Implemeted!!")
}
func (tcpServer *__tcpServerStruct) Destroy() error {
	var err error
	defer func() {
		itf := recover()
		if errs.IsError(itf) {
			err = itf.(error)
			if tcpServer.__logger != nil {
				tcpServer.__logger.Log(common.ERROR, fmt.Sprintf("Error executing tcp server Destroy : %s", err.Error()))
			}
		} else {
			err = errors.New(fmt.Sprintf("%v", itf))
			if tcpServer.__logger != nil {
				tcpServer.__logger.Log(common.ERROR, fmt.Sprintf("Error executing tcp server Destroy : %v", err))
			}
		}
	}()
	if tcpServer.__logger != nil {
		tcpServer.__logger.Close()
		tcpServer.__logger = nil
	}
	return err
}
func (tcpServer *__tcpServerStruct) WaitFor() error {
	var err error
	defer func() {
		itf := recover()
		if errs.IsError(itf) {
			err = itf.(error)
			tcpServer.__logger.Log(common.ERROR, fmt.Sprintf("Error executing tcp server WaitFor : %s", err.Error()))
		} else {
			err = errors.New(fmt.Sprintf("%v", itf))
			tcpServer.__logger.Log(common.ERROR, fmt.Sprintf("Error executing tcp server WaitFor : %v", err))
		}
	}()
	for tcpServer.__running {
		time.Sleep(1 * time.Second)
	}
	return err
}
func (tcpServer *__tcpServerStruct) Clients() []common.ClientRef {
	return []common.ClientRef{}
}
func (tcpServer *__tcpServerStruct) LogOn(channel *chan interface{}) error {
	return tcpServer.__logger.AddOutChannel(channel)
}
func (tcpServer *__tcpServerStruct) Logger() common.ServerLogger {
	return tcpServer.__logger
}

func (tcpServer *__tcpServerStruct) GetCacheEntries() *common.NetCache {
	return &tcpServer.__globalCache
}
func (tcpServer *__tcpServerStruct) GetServiceCacheEntries(keyCode string) *common.NetCache {
	cache, ok := tcpServer.__serviceCache[keyCode]
	if !ok {
		return nil
	}
	return &cache
}
func (tcpServer *__tcpServerStruct) GetServerEnv() *common.NetEnvironment {
	return &tcpServer.__globalEnv
}
func (tcpServer *__tcpServerStruct) GetServiceEnv(keyCode string) *common.NetEnvironment {
	env, ok := tcpServer.__serviceEnv[keyCode]
	if !ok {
		return nil
	}
	return &env
}

// Define Tcp Rest server based on inpuit parameters
// Paarameters:
//    logLevel (common.LogLevel) Server Log level verbosity
//    host (string) Server Listening Host/IP address
//    port (common.Port) Server Listening port
//    entriesList (TcpEntriesList) Tcp Endpoint entries list
//    closure (TcpCycleClosureFunc) Function invoked at the end of any acceptance of client and handler execution
// Returns:
//    common.Server Tcp Server Instance
func NewTcpServer(logLevel common.LogLevel, host string, port common.Port, entriesList TcpEntriesList, closure TcpCycleClosureFunc) common.Server {
	logger := common.NewServerLogger(logLevel)
	var closureFunc TcpCycleClosureFunc = closure
	if closureFunc == nil {
		closureFunc = __defaultClosureFunc
	}

	return &__tcpServerStruct{
		ListeningPort: port,
		ListeningHost: host,
		__logger:      logger,
		EntriesList:   entriesList,
		ClosureFunc:   closure,
	}
}

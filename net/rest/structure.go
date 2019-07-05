package rest

import (
	"errors"
	"fmt"
	errs "github.com/hellgate75/general_utils/errors"
	"github.com/hellgate75/general_utils/net/common"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// Rest Endpoint handler function
type RestEndpoint common.RestAction

// Rest Endpoint Map type
type RestEntriesMap map[string]RestEndpoint

type __restServerStruct struct {
	EntriesMap     RestEntriesMap
	StateHandler   common.HttpStateHandler
	ListeningPort  common.Port
	__lock         sync.Locker
	__globalCache  common.RestCache
	__serviceCache map[string]common.RestCache
	__globalEnv    common.RestEnvironment
	__serviceEnv   map[string]common.RestEnvironment
	__serverPipe   chan interface{}
	__running      bool
	__logger       common.ServerLogger
}

func (rest *__restServerStruct) __baseRestEntryActionFunc(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		itf := recover()
		if errs.IsError(itf) {
			err = itf.(error)
			rest.__logger.Log(common.ERROR, fmt.Sprintf("Error executing rest : %s", err.Error()))
		} else {
			rest.__logger.Log(common.ERROR, fmt.Sprintf("Error executing rest : %v", err))
		}
	}()
	var path string = r.URL.Path
	var query url.Values = r.URL.Query()
	if r.Method != "GET" && r.Method != "DELETE" {
		query = r.PostForm
	}
	rest.__lock.Lock()
	if restEndPoint, ok := rest.EntriesMap[path]; ok {
		err = restEndPoint(rest.StateHandler, query, &rest.__serverPipe, w, r, rest)
		if err != nil {
			rest.__logger.Log(common.ERROR, fmt.Sprintf("Error executing rest : %s", err.Error()))
		}
	} else {
		err = errors.New(fmt.Sprintf("Rest point not found : %s, Error 404"))
		action, errAct := rest.StateHandler.GetActionByCode(404)
		if errAct != nil {
			rest.__logger.Log(common.ERROR, fmt.Sprintf("Error executing rest : %s", errAct.Error()))
			rest.__lock.Unlock()
			return
		}
		err = action(w, r)
		if err != nil {
			rest.__logger.Log(common.ERROR, fmt.Sprintf("Error executing status http action : %s", err.Error()))
		}
	}
	rest.__lock.Unlock()
	if err != nil {
		rest.__logger.Log(common.ERROR, fmt.Sprintf("Error executing rest : %s", err.Error()))
	}
}

func (rest *__restServerStruct) Open() error {
	var err error
	defer func() {
		itf := recover()
		if errs.IsError(itf) {
			err = itf.(error)
		}
	}()
	rest.__running = true
	rest.__logger.Open()
	http.HandleFunc("/", rest.__baseRestEntryActionFunc)
	http.ListenAndServe(fmt.Sprintf(":%v", rest.ListeningPort), nil)
	rest.__serverPipe = make(chan interface{})
	rest.__globalCache = make(common.RestCache)
	rest.__globalEnv = make(common.RestEnvironment)
	rest.__serviceCache = make(map[string]common.RestCache)
	rest.__serviceEnv = make(map[string]common.RestEnvironment)
	for k, _ := range rest.EntriesMap {
		rest.__serviceCache[k] = make(common.RestCache)
		rest.__serviceEnv[k] = make(common.RestEnvironment)
	}
	return err
}
func (rest *__restServerStruct) Close() error {
	var err error
	defer func() {
		itf := recover()
		if errs.IsError(itf) {
			err = itf.(error)
		}
	}()
	rest.__logger.Close()
	rest.__running = false
	close(rest.__serverPipe)
	return err
}
func (rest *__restServerStruct) IsListening() bool {
	return rest.__running
}
func (rest *__restServerStruct) IsRunning() bool {
	return rest.__running
}
func (rest *__restServerStruct) Stream(interface{}) error {
	return errors.New("Not Implemeted!!")

}
func (rest *__restServerStruct) Receive() (interface{}, error) {
	return nil, errors.New("Not Implemeted!!")
}
func (rest *__restServerStruct) HandleConnOn(hf common.ServerHablerFunc) error {
	return errors.New("Not Implemeted!!")
}
func (rest *__restServerStruct) HandlingFuncs() []common.ServerHablerFunc {
	return []common.ServerHablerFunc{}
}
func (rest *__restServerStruct) RevokeHandler(hf common.ServerHablerFunc) error {
	return errors.New("Not Implemeted!!")
}
func (rest *__restServerStruct) Destroy() error {
	var err error
	defer func() {
		itf := recover()
		if errs.IsError(itf) {
			err = itf.(error)
		}
	}()
	rest.__logger = nil
	return err
}
func (rest *__restServerStruct) WaitFor() error {
	var err error
	defer func() {
		itf := recover()
		if errs.IsError(itf) {
			err = itf.(error)
		}
	}()
	for rest.__running {
		time.Sleep(1 * time.Second)
	}
	return err
}
func (rest *__restServerStruct) Clients() []common.ClientRef {
	return []common.ClientRef{}
}
func (rest *__restServerStruct) LogOn(channel *chan interface{}) error {
	return rest.__logger.AddOutChannel(channel)
}
func (rest *__restServerStruct) Logger() common.ServerLogger {
	return rest.__logger
}

func (rest *__restServerStruct) GetCacheEntries() *common.RestCache {
	return &rest.__globalCache
}
func (rest *__restServerStruct) GetServiceCacheEntries(restPath string) *common.RestCache {
	cache, ok := rest.__serviceCache[restPath]
	if !ok {
		return nil
	}
	return &cache
}
func (rest *__restServerStruct) GetServerEnv() *common.RestEnvironment {
	return &rest.__globalEnv
}
func (rest *__restServerStruct) GetServiceEnv(restPath string) *common.RestEnvironment {
	env, ok := rest.__serviceEnv[restPath]
	if !ok {
		return nil
	}
	return &env
}

// Define New Rest server based on inpuit parameters
// Paarameters:
//    logLevel (common.LogLevel) Server Log level verbosity
//    port (common.Port) Server Listening port
//    stateHandler (common.HttpStateHandler) Http State Handler, passed to Rest service (see common.NewHttpStateHandler function)
//     entriesMap (RestEntriesMap) Rest Endpoint entries map
// Returns:
//    common.Server Rest Server Instance
func NewRestServer(logLevel common.LogLevel, port common.Port, stateHandler common.HttpStateHandler, entriesMap RestEntriesMap) common.Server {
	logger := common.NewServerLogger(logLevel)
	return &__restServerStruct{
		ListeningPort: port,
		__logger:      logger,
		StateHandler:  stateHandler,
		EntriesMap:    entriesMap,
	}
}

package common

import (
	"errors"
	"fmt"
	errs "github.com/hellgate75/general_utils/errors"
	"net/http"
	"net/url"
)

type HTTPAction func(http.ResponseWriter, *http.Request) error

func __defaulManageNotFound(w http.ResponseWriter, r *http.Request) error {
	var err error
	defer func() {
		itf := recover()
		if errs.IsError(itf) {
			err = itf.(error)
		}
	}()
	http.NotFound(w, r)
	return err
}

func __defaulManageInternalError(w http.ResponseWriter, r *http.Request) error {
	var err error
	defer func() {
		itf := recover()
		if errs.IsError(itf) {
			err = itf.(error)
		}
	}()
	http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	return err
}

// Create HTTPAction based on text message and http.Status.... status code value
// Parameters:
//    message (string) Http state message
//    code (int) http.Status.... status code value
// Returns:
//    HTTPAction Early defined HTTPAction function
func CreateHttpActionByStatus(message string, code int) HTTPAction {
	return func(w http.ResponseWriter, r *http.Request) error {
		var err error
		defer func() {
			itf := recover()
			if errs.IsError(itf) {
				err = itf.(error)
			}
		}()
		http.Error(w, message, code)
		return err
	}
}

// Default HTTPActon Handling 404 Not found Message
var DefaultNotFoundAction HTTPAction = __defaulManageNotFound

// Default HTTPActon Handling 500 Internal Server Error Message
var DefaultInternalServerErrorAction HTTPAction = __defaulManageInternalError

// Interface that defines the default http status habdler functions
type HttpStateHandler interface {
	// Create HTTPAction based on text message and http.Status.... status code value
	// Parameters:
	//    message (string) Http state message
	// Returns:
	//   HTTPAction Early defined HTTPAction function
	//    error Any error that can occurs during computation
	GetActionByCode(code int) (HTTPAction, error)
	// Create HTTPAction based on text message and http.Status.... status code value
	// Parameters:
	//    message (string) Http state message
	//    code (INT) http.Status.... status code value
	// Returns:
	//    error Any error that can occurs during computation
	AddActionByCode(message string, code int) error
	// Create HTTPAction based on text message and http.Status.... status code value
	// Parameters:
	//    message (string) Http state message
	//    code (INT) http.Status.... status code value
	// Returns:
	//    error Any error that can occurs during computation
	AddActionsByMap(map[int]HTTPAction) error
}

type __httpStateHandlerStruct struct {
	__stateHandlerMap map[int]HTTPAction
}

func (eh *__httpStateHandlerStruct) GetActionByCode(code int) (HTTPAction, error) {
	var err error
	var action HTTPAction
	defer func() {
		itf := recover()
		if errs.IsError(itf) {
			err = itf.(error)
		}
	}()
	if act, ok := eh.__stateHandlerMap[code]; ok {
		action = act
	} else {
		err = errors.New(fmt.Sprintf("Give definition does not exist for code : %x", code))
	}
	return action, err
}
func (eh *__httpStateHandlerStruct) AddActionByCode(message string, code int) error {
	var err error
	defer func() {
		itf := recover()
		if errs.IsError(itf) {
			err = itf.(error)
		}
	}()
	if _, ok := eh.__stateHandlerMap[code]; ok {
		err = errors.New(fmt.Sprintf("Give definition have overwritten another in the registry : %x", code))
	}
	eh.__stateHandlerMap[code] = CreateHttpActionByStatus(message, code)
	return err
}
func (eh *__httpStateHandlerStruct) AddActionsByMap(newErrorMap map[int]HTTPAction) error {
	var err error
	defer func() {
		itf := recover()
		if errs.IsError(itf) {
			err = itf.(error)
		}
	}()
	var overwriteList []int
	for k, v := range newErrorMap {
		if _, ok := eh.__stateHandlerMap[k]; ok {
			overwriteList = append(overwriteList, k)
		}
		eh.__stateHandlerMap[k] = v
	}
	if len(overwriteList) > 0 {
		err = errors.New(fmt.Sprintf("Some definitions have overwritten others in the registry : %v", overwriteList))
	}
	return err
}

// Create HTTPAction based on text message and http.Status.... status code value
// Parameters:
//    message (string) Http state message
//    code (INT) http.Status.... status code value
// Returns:
//    HTTPAction Early defined HTTPAction function
//    error Any error that can occurs during computation
func NewHttpStateHandler(errorMap map[int]HTTPAction) (HttpStateHandler, error) {
	var err error
	if len(errorMap) == 0 {
		errorMap[500] = __defaulManageInternalError
		errorMap[404] = __defaulManageNotFound
	} else {
		if _, ok := errorMap[404]; !ok {
			errorMap[404] = __defaulManageNotFound
		}
		if _, ok := errorMap[500]; !ok {
			errorMap[500] = __defaulManageInternalError
		}
	}
	val := &__httpStateHandlerStruct{
		__stateHandlerMap: errorMap,
	}
	return val, err
}

// Rest Action function type
type RestAction func(HttpStateHandler, url.Values, *chan interface{}, http.ResponseWriter, *http.Request, NetContext) error

// Network Cache Map type
type NetCache map[interface{}]interface{}

// Network Environment Map type
type NetEnvironment map[interface{}]interface{}

//Network Service Context
type NetContext interface {
	// Get Global Rest Server Cache Map
	// Rerturns:
	//    *NetCache Pointer to Rest Service Cache Map
	GetCacheEntries() *NetCache
	// Get Rest Endpoint Cache Map
	// Parameters:
	//    restPath (string) Rest Service path, kay for map storage
	// Rerturns:
	//    *NetCache Pointer to Endpoint Cache Map
	GetServiceCacheEntries(restPath string) *NetCache
	// Get Global Rest Server Environment Map
	// Rerturns:
	//    *NetEnvironment Pointer to Rest Service Environment Map
	GetServerEnv() *NetEnvironment
	// Get Rest Endpoint Environment Map
	// Parameters:
	//    restPath (string) Rest Service path, kay for map storage
	// Rerturns:
	//    *NetEnvironment Pointer to Endpoint Environment Map
	GetServiceEnv(restPath string) *NetEnvironment
}

package net

import (
	"errors"
	"fmt"
	errs "github.com/hellgate75/general_utils/errors"
	"github.com/hellgate75/general_utils/log"
	"github.com/hellgate75/general_utils/net/common"
	"github.com/hellgate75/general_utils/net/rest"
	"github.com/hellgate75/general_utils/net/tcp"
	"github.com/hellgate75/general_utils/net/web"
)

//Create a New Server Accordingly to Server Type Parameters
//  Parameters:
//    serverType (common.ServerType) Server Type, to be built and returned
//    args (interface{} array variadic) Arguments, related to the server configuration as follow
//  Assegnment for Arguments:
//  *  TCP Server: (common.TCP)
//       logLevel (common.LogLevel) Server Log level verbosity
//       host (string) Server Listening Host/IP address
//       port (common.Port) Server Listening port
//       entriesList (tcp.TcpEntriesList) Tcp Endpoint entries list
//       closure (tcp.TcpCycleClosureFunc) Function invoked at the end of any acceptance of client and handler execution
//  *  REST Server: (common.REST)
//       logLevel (common.LogLevel) Server Log level verbosity
//       host (string) Server Listening Host/IP address
//       port (common.Port) Server Listening port
//       stateHandler (common.HttpStateHandler) Http State Handler, replaced with default structure is case it's passed nil value
//       entriesMap (rest.RestEntriesMap) Rest Endpoint entries map
//  *  Web Server: (common.CONTENT)
//       baseFolder (string) Base FS folder that container content sources
//       config (common.ServerConfig) Server initialization parameters group
//       env (map[string]interface{}) Map of environment items
//       validator (*web.WebServerValidator) Pointer to Web Server Service/Call Validator
//       notFoundAction (*common.HTTPAction) Pointer to 404 (Not found) state http handler
// Returns:
//    (common.Server built server instance,
//     error Any error that occurs during computation)
func NewServer(serverType common.ServerType, args ...interface{}) (common.Server, error) {
	var err error
	var server common.Server
	defer func() {
		itf := recover()
		if errs.IsError(itf) {
			err = itf.(error)
			if logger != nil {
				logger.Log(log.ERROR, fmt.Sprintf("Error executing tcp server Destroy : %s", err.Error()))
			}
		} else {
			err = errors.New(fmt.Sprintf("%v", itf))
			if logger != nil {
				logger.Log(log.ERROR, fmt.Sprintf("Error executing tcp server Destroy : %v", itf))
			}
		}
	}()
	if len(args) == 0 {
		return nil, errors.New(fmt.Sprintf("Insufficient arguments number <%x> for the server type <%v>!!", len(args), serverType))
	}
	switch serverType {
	case common.TCP:
		if len(args) < 5 {
			return nil, errors.New(fmt.Sprintf("Insufficient arguments number <%x> (minimum : 5) for the TCP Server Type!!", len(args)))
		}
		server = tcp.NewTcpServer(args[0].(common.LogLevel), args[1].(string), args[2].(common.Port), args[3].(tcp.TcpEntriesList), args[4].(tcp.TcpCycleClosureFunc))
	case common.REST:
		if len(args) < 5 {
			return nil, errors.New(fmt.Sprintf("Insufficient arguments number <%x> (minimum : 5) for the REST Server Type!!", len(args)))
		}
		server = rest.NewRestServer(args[0].(common.LogLevel), args[1].(string), args[2].(common.Port), args[3].(common.HttpStateHandler), args[3].(rest.RestEntriesMap))
	case common.CONTENT:
		if len(args) < 5 {
			return nil, errors.New(fmt.Sprintf("Insufficient arguments number <%x> (minimum : 5) for the Web Server Type!!", len(args)))
		}
		server = web.NewWebServer(args[0].(string), args[1].(common.ServerConfig), args[2].(map[string]interface{}), args[3].(*web.WebServerValidator), args[4].(*common.HTTPAction))
	default:
		return nil, errors.New(fmt.Sprintf("Server type <%v> not implemented!!", serverType))
	}
	return server, err
}

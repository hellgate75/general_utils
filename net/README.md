## package net  import "github.com/hellgate75/general_utils/net"

##### Package network implements networking operations in type.

### FUNCTIONS

#### func NewServer(serverType common.ServerType, args ...interface{}) (common.Server, error) {
Create a New Server Accordingly to Server Type Parameters
#####  Parameters:
    serverType (common.ServerType) Server Type, to be built and returned
    args (interface{} array variadic) Arguments, related to the server configuration as follow
#####  Assegnment for Arguments:
#####  *  TCP Server: (common.TCP)
       logLevel (common.LogLevel) Server Log level verbosity
       host (string) Server Listening Host/IP address
       port (common.Port) Server Listening port
       entriesList (tcp.TcpEntriesList) Tcp Endpoint entries list
       closure (tcp.TcpCycleClosureFunc) Function invoked at the end of any acceptance of client and handler execution
#####  *  REST Server: (common.REST)
       logLevel (common.LogLevel) Server Log level verbosity
       host (string) Server Listening Host/IP address
       port (common.Port) Server Listening port
       stateHandler (common.HttpStateHandler) Http State Handler, replaced with default structure is case it's passed nil value
       entriesMap (rest.RestEntriesMap) Rest Endpoint entries map
#####  *  Web Server: (common.CONTENT)
       baseFolder (string) Base FS folder that container content sources
       config (common.ServerConfig) Server initialization parameters group
       env (map[string]interface{}) Map of environment items
       validator (*web.WebServerValidator) Pointer to Web Server Service/Call Validator
       notFoundAction (*common.HTTPAction) Pointer to 404 (Not found) state http handler
##### Returns:
    (common.Server built server instance,
     error Any error that occurs during computation)
        
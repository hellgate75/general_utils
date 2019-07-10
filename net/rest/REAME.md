## package rest // import "github.com/hellgate75/general_utils/net/rest"


### FUNCTIONS

#### func InitLogger()
     Initialize package logger if not started
     
#### func NewRestServer(logLevel common.LogLevel, host string, port common.Port, stateHandler common.HttpStateHandler, entriesMap RestEntriesMap) common.Server
    Define New Rest server based on inpuit parameters 
#####    Parameters:

       logLevel (common.LogLevel) Server Log level verbosity
       host (string) Server Listening Host/IP address
       port (common.Port) Server Listening port
       entriesMap (RestEntriesMap) Rest Endpoint entries map
#####    Returns:
       common.Server Rest Server Instance


### TYPES

##### type RestEndpoint common.RestAction
    Rest Endpoint handler function

##### type RestEntriesMap map[string]RestEndpoint
    Rest Endpoint Map type


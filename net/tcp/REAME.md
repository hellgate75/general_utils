## package tcp // import "github.com/hellgate75/general_utils/net/tcp"


### CONSTANTS

##### const (
##### 	CONN_TYPE = "tcp"
##### )

### FUNCTIONS

#### func InitLogger()
     Initialize package logger if not started

#### func NewTcpServer(logLevel common.LogLevel, host string, port common.Port, entriesList TcpEntriesList, closure TcpCycleClosureFunc) common.Server
    Define Tcp Rest server based on inpuit parameters 
#####    Parameters:
       logLevel (common.LogLevel) Server Log level verbosity
       host (string) Server Listening Host/IP address
       port (common.Port) Server Listening port
       entriesList (TcpEntriesList) Tcp Endpoint entries list
       closure (TcpCycleClosureFunc) Function invoked at the end of any acceptance of client and handler execution
#####    Returns:
       common.Server Tcp Server Instance


### TYPES

##### type TcpCycleClosureFunc func(net.Conn, common.ServerLogger) error
    Call closure function

##### type TcpEndpoint func(net.Conn, *chan interface{}, common.NetContext) error
    Tcp Endpoint handler function

##### type TcpEntriesList []TcpEndpoint
    Tcp Endpoint Map type


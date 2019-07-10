## package general_utils // import "github.com/hellgate75/general_utils"


### FUNCTIONS

#### func DecomposeFilePath(filePath string) (string, string, []string, error)
    Decompose a file path to recover most relevant information 
#####    Parameters:
       filePath (string) file absolute path (with extension)
       fileName (string) file name without extension
       extensions ([]string) list of suitable file extensions
#####    Returns:
      ( string absolute file folder path, string absolute file path, []stirng list of suitable extensions, error Any error arisen during computation )

#### func DestroyLoggerEngine()
    Destroys and unregister Logger Engine

#### func DisableLogChanMode()
    Disables Log to default Channel and restore log to target configuration

#### func EnableLogChanMode() *chan interface{}
    Enables Log to default Channel instead of target configuration

#### func FindFileInPath(folder string, fileName string, extensions []string) (string, string, error)
    Finds a file into a folder and look forward a file matching multiple file
    extensions
#####    Parameters:
    folder (string) folder containing files
    fileName (string) file name without extension
    extensions ([]string) list of suitable file extensions
#####    Returns:
    ( string absolute file path, stirng file extension, error Any error arisen during computation )

#### func InitCustomLoggerEngine(filePath string) error
    Initializes and registers full Logger from custom configuration file.
#####    Parameters:
    filePath (string) Absolute file path
#####    Returns:
    error Any suitable error risen during code execution

#### func InitDeviceLoggerEngine() error
    Initializes and registers full Logger from base configuration file.
#####    Returns:
    error Any suitable error risen during code execution

#### func InitSimpleLoggerEngine(verbosity log.LogLevel)
    Initializes and registers Logger for StdOut.
#####    Parameters:
    verbosity (log.LogLevel) Log verbosity level

#### func InitializeLoggers()
    Initialize loggers for all sub packages in global_utils. To run after the
    logger initialization only

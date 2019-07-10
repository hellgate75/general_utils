package logger // import "github.com/hellgate75/general_utils/parsers/logger"


FUNCTIONS

func GetLogConfig() log.LogConfig
    Retrieve loaded log configuration.

    Returns: log.LogConfig Loaded Log Configuration

func LoadLoggerConfigFromFile(encoding parsers.Encoding, filePath string) (*log.LogConfig, error)
    Provides read access to log configuration and parse from a File.

    Parameters:

    encoding (parsers.Encoding) File Encoding type (JSON, XML, YAML)
    filePath (string) File full path

    Returns: error Any suitable error risen during code execution

func SetLogConfig(config log.LogConfig)
    Set loaded log configuration.

    Paraeters: config (log.LogConfig) Loaded Log Configuration


## package config // import "github.com/hellgate75/general_utils/config"


### FUNCTIONS

#### func InitDatabaseConfig(arFilePath string, mask *common.Type, deserializer func([]byte, *common.Type) error) error
    Initialize Database Configuration from file path, using deserialization function
#####    Parameters:
    arFilePath (string) .ar Database file path
    mask (*common.Type) Pointer to target configuration structure
    deserializer (func([]byte, *common.Type) error) error) Deserializer function to seek into the parser package
#####    Returns:
    error Any suitable error risen during code execution

#### func LoadDatabaseFromURL(arFileName string, url string, mask *common.Type, deserializer func([]byte, *common.Type) error) error
    Load database from URL using deserialization function
#####    Parameters:
    arFileName (string) .ar Remote URL path
    mask (*common.Type) Pointer to target configuration structure
    deserializer (func([]byte, *common.Type) error) error) Deserializer function to seek into the parser package
#####    Returns:
    error Any suitable error risen during code execution

#### func ReadDatabaseConfig(arFileName string) (error, map\[string\][]byte)
    Read .ar database data and save into an output map
#####    Parameters:
    arFileName (string) .ar file path
#####    Returns:
    error Any suitable error risen during code execution
     map[string][]byte Map containing database data

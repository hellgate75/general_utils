## package common // import "github.com/hellgate75/general_utils/common"


### FUNCTIONS

#### func StorageUnitBytesFactor(unit StorageUnit) (int64, error)
    Returns Storage Unit representin factor in number of bytes.
#####    Parameters:
       unit (common.StorageUnit) Reference Storage Unit
#####    Returns: 
       int64 Factor of conversion to bytes or 0 anyway error Any suitable
       error risen during code execution

#### func StorageUnitToString(unit StorageUnit) (string, error)
    Transform Storage Unit in representing text.

    Parameters:

    unit (common.StorageUnit) Storage Unit

    Returns: string Text representing the Storage Unit or 0 anyway error Any
    suitable error risen during code execution

#### func StreamInOutFormatToString(format StreamInOutFormat) (string, error)
    Transform Stream Input/Output format in representing text.
#####    Parameters:
      format (common.StreamInOutFormat) Stream Input/Output format
#####    Returns: 
      string Text representing the Stream Input/Output format or 0 anyway
      error Any suitable error risen during code execution

#### func WriterTypeToString(wType WriterType) (string, error)
    Transform Writer Type in representing text.
#####    Parameters:
    wType (common.WriterType) Writer Type
#####    Returns: 
    string Text representing the WriterType or 0 anyway error Any
    suitable error risen during code execution

#### func StringToStorageUnit(text string) (StorageUnit, error)
    Transform text in Storage Unit.
#####     Parameters:
       text (string) Text to convert
#####     Returns: 
       common.StorageUnit Storage Unit representing the text data or 0
       anyway error Any suitable error risen during code execution

#### func StringToStreamInOutFormat(text string) (StreamInOutFormat, error)
    Transform text in Stream Input/Output format.
#####     Parameters:
       text (string) Text to convert

#####     Returns: 
       common.StreamInOutFormat Stream Input/Output format representing the text data or 0 anyway 
       error Any suitable error risen during code execution

#### func StringToWriterType(text string) (WriterType, error)
    Transform text in Writer Type.
#####     Parameters:
       text (string) Text to convert

#####    Returns: 
       common.WriterType Writer Type representing the text data or 0 anyway 
       error Any suitable error risen during code execution


### TYPES

##### type Message Type
    Generic Message

##### type StorageUnit int
    Storage Unit representing type

##### const (
##### 	//bytes = 1b
##### 	Bt StorageUnit = 1
##### 	//KiloBytes = 1024b
##### 	Kbt StorageUnit = 2
##### 	//MegaBytes = 1048576b
##### 	Mbt StorageUnit = 3
##### 	//GigaBytes = 1073741824b
##### 	Gbt StorageUnit = 4
##### 	//TeraBytes = 1099511627776b
##### 	Tbt StorageUnit = 5
##### )

##### type StreamInOutFormat int
    Storage Unit representing type

##### const (
##### 	//Plain text input/output format
##### 	PlainTextFormat StreamInOutFormat = 11
##### 	//Json input/output format
##### 	JsonFormat StreamInOutFormat = 12
##### 	//Yaml input/output format
##### 	YamlFormat StreamInOutFormat = 13
##### 	//Xml input/output format
##### 	XmlFormat StreamInOutFormat = 14
##### 	//Base64 encrypted text input/output format
##### 	EncryptedFormat StreamInOutFormat = 15
##### 	//Text encrypted input/output format
##### 	BinaryFormat StreamInOutFormat = 16
##### 	//Text Go Format input/output format
##### 	GoStructFormat StreamInOutFormat = 17
##### )

##### type Type interface{}
    Generic Type

##### type WriterType string
    Type WriterType describe any Writer Option in the cofiguration

##### const (
##### 	// Stadard Output Writer Type
##### 	StdOutWriter WriterType = "StdOut"
##### 	// File Output Writer Type
##### 	FileWriter WriterType = "File"
##### 	// URL Output Writer Type
##### 	UrlWriter WriterType = "Url"
##### )


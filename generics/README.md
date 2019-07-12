## package generics // import "github.com/hellgate75/general_utils/generics"


### FUNCTIONS

#### func GetAsMapImpl(yourInterface interface{}) map[string]interface{}
    Implement Serializable#GetAsMap standard solution
#####     Parameters:
    yourInterface (interface{} NOT POINTER) recover information from your interface and return the attributes map
#####     Returns:
    map[string]interface{} That represent the attributes of the interface as map

#### func InitLogger()
     Initialize package logger if not started

#### func PrepareFileStructMap(element interface{}) (map[string]FieldConfig, error)
     Prepare element structure map
#####     Parameters:
    element (interface{}) input element
#####     Returns:
    (map[string]FieldConfig That represent the attributes of the interface as map,
     error Any arror that can occur during the computation)

#### func SetFromMapImpl(yourInterface interface{}, attributes map[string]interface{}) (interface{}, error)
    Implement Serializable#SetFromMap standard solution
#####     Parameters:
    yourInterface (interface{} pointer) set attributes into your interface
    attributes (map[string]interface{}) Map of attribute names and values to set up into the interface
#####     Returns:
    error Any arror that can occur during the computation


TYPES

##### type FieldConfig struct {
##### 	Name           string
##### 	Type           reflect.Type
##### 	Value          reflect.Value
##### 	IsNil          bool
##### 	FieldStructure reflect.StructField
##### }
    Type containing most relevant reflection information for any of the
    attributes in your interface

#### func PrepareFileStructList(element interface{}) ([]FieldConfig, error)
     Prepare element structure list
#####     Parameters:
    element (interface{}) input element
#####     Returns:
    ([]FieldConfig That represent the attributes of the interface as list,
     error Any arror that can occur during the computation)

##### type Serializable interface {
##### 	//Retrive attributes map
##### 	GetAsMap() map[string]interface{}
##### 	//Set attributes from attributes map map
##### 	SetFromMap(map[string]interface{}) (interface{}, error)
##### }
    Generic Serializable Type

#### func NewSerializable() Serializable
    Returns a new Emty serializable object instance
#####     Returns:
    Serializable - Empty Serializable interface instance

##### type SerializableStruct struct {
##### }

##### func (this *SerializableStruct) GetAsMap() map[string]interface{}
    Retrive the Map of elements in the input value
#####     Returns:
    map[string]interface{} Map of attribute names and values
     
##### func (this *SerializableStruct) SetFromMap(attributes map[string]interface{}) (interface{}, error)
    Set value from an attributes map
#####     Returns:
    (interface{} Modified Element value,
     error Any arror that can occur during the computation)

##### type StructureCleaner interface {
##### 	CleanupInterfaceArray(in []interface{}) []interface{}
##### 	CleanupInterfaceMap(in map[interface{}]interface{}) map[string]interface{}
##### 	CleanupMapValue(v interface{}) interface{}
##### 	MapToInterface(v map[interface{}]interface{}) interface{}
##### 	CleanupInterfaceObject(in interface{}) interface{}
##### }

##### func NewStructureCleaner() StructureCleaner
      Return a new Structure Cleaner
#####     Returns:
    StructureCleaner Is instance of StructureCleaner interface

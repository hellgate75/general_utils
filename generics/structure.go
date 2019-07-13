package generics

import (
	"errors"
	"fmt"
	errs "github.com/hellgate75/general_utils/errors"
	"reflect"
)

// Generic Serializable Type
type Serializable interface {
	//Retrive attributes map
	GetAsMap() map[string]interface{}
	//Set attributes from attributes map map
	SetFromMap(map[string]interface{}) (interface{}, error)
}

type SerializableStruct struct {
}

func (this *SerializableStruct) GetAsMap() map[string]interface{} {
	return GetAsMapImpl(this)
}

func (this *SerializableStruct) SetFromMap(attributes map[string]interface{}) (interface{}, error) {
	return SetFromMapImpl(this, attributes)
}

// Type containing most relevant reflection information for any of the attributes in your interface
type FieldConfig struct {
	Name           string
	Type           reflect.Type
	Value          reflect.Value
	IsNil          bool
	FieldStructure reflect.StructField
}

// Implement Serializable#GetAsMap standard solution
//
// Parameters:
//   yourInterface (interface{} NOT POINTER) recover information from your interface and return the attributes map
//
// Returns:
//   map[string]interface{} That represent the attributes of the interface as map
func GetAsMapImpl(yourInterface interface{}) map[string]interface{} {
	var attributes map[string]interface{} = make(map[string]interface{})
	objVal := reflect.ValueOf(yourInterface)
	//	fmt.Println(objVal.Type().Kind().String())
	noFlds := objVal.Type().NumField()
	for i := 0; i < noFlds; i++ {
		fld := objVal.Type().Field(i)
		k := fld.Name
		//		fmt.Println(objVal.Field(i).Type().Kind().String())
		v := objVal.Field(i).Interface()
		attributes[fmt.Sprintf("%v", k)] = v
	}
	return attributes
}

// Extract FieldConfig from interace is not Nil, and it's a structure
//
// Parameters:
//    element (interface{}) Element we want inspect into
//
// Returns:
//   []FieldConfig - Field configuration list accordingly to interface structure

func PrepareFileStructList(element interface{}) ([]FieldConfig, error) {
	var err error
	defer func() {
		itf := recover()
		if errs.IsError(itf) {
			err = itf.(error)
			if logger != nil {
				logger.ErrorS(fmt.Sprintf("Error in PrepareFileStructList, message : %s", err.Error()))
			}

		} else {
			err = errors.New(fmt.Sprintf("%v", itf))
			if logger != nil {
				logger.ErrorS(fmt.Sprintf("Error in PrepareFileStructList, message : %v", itf))
			}
		}
	}()
	var configList []FieldConfig = []FieldConfig{}
	if element != nil && reflect.ValueOf(element).Kind() == reflect.Struct {
		elemRef := reflect.Indirect(reflect.ValueOf(element))
		for i := 0; i < elemRef.NumField(); i++ {
			fs := elemRef.Type().Field(i)
			configList = append(configList, FieldConfig{
				Name:           fs.Name,
				Type:           fs.Type,
				Value:          elemRef.Field(i),
				IsNil:          false,
				FieldStructure: fs,
			})
		}
	}
	return configList, err
}

// Extract FieldConfig from interace is not Nil and it's a structure
//
// Parameters:
//    element (interface{}) Element we want inspect into
//
// Returns:
//   map[string]FieldConfig - Field configuration map by field name, accordingly to interface structure

func PrepareFileStructMap(element interface{}) (map[string]FieldConfig, error) {
	var configMap map[string]FieldConfig = make(map[string]FieldConfig)
	var err error
	defer func() {
		itf := recover()
		if logger != nil {
			if errs.IsError(itf) {
				err = itf.(error)
				if logger != nil {
					logger.ErrorS(fmt.Sprintf("Error in PrepareFileStructMap, message : %s", err.Error()))
				}
			} else {
				err = errors.New(fmt.Sprintf("%v", itf))
				if logger != nil {
					logger.ErrorS(fmt.Sprintf("Error in PrepareFileStructMap, message : %v", itf))
				}
			}
		}
	}()
	if element != nil && reflect.ValueOf(element).Kind() == reflect.Struct {
		elemRef := reflect.Indirect(reflect.ValueOf(element))
		for i := 0; i < elemRef.NumField(); i++ {
			fs := elemRef.Type().Field(i)
			configMap[fs.Name] = FieldConfig{
				Name:           fs.Name,
				Type:           fs.Type,
				Value:          elemRef.Field(i),
				IsNil:          false,
				FieldStructure: fs,
			}
		}
	}
	return configMap, err
}

// Implement Serializable#SetFromMap standard solution
//
// Parameters:
//   yourInterface (interface{} pointer) set attributes into your interface
//   attributes (map[string]interface{}) Map of attribute names and values to set up into the interface
//
// Returns:
//   error Any arror that can occur during the computation
func SetFromMapImpl(yourInterface interface{}, attributes map[string]interface{}) (interface{}, error) {
	//	var fieldList []FieldConfig
	//	var wrongFieldList []FieldConfig
	//	var discardedFieldList []FieldConfig
	//	var yourValue reflect.Value = reflect.Indirect(reflect.ValueOf(yourInterface))
	//	var yourElement reflect.Value = reflect.Indirect(reflect.ValueOf(yourInterface)).Elem()
	yourElement := reflect.ValueOf(&yourInterface).Elem()
	//	var yourType reflect.Type = yourValue.Type()
	//	var attributesMap []reflect.StructField
	var err error
	defer func() {
		itf := recover()
		if errs.IsError(itf) {
			err = itf.(error)
			logger.ErrorS(fmt.Sprintf("Error in SetFromMapImpl(), message : %s", err.Error()))
		} else {
			err = errors.New(fmt.Sprintf("%v", itf))
			logger.ErrorS(fmt.Sprintf("Error in SetFromMapImpl(), message : %v", itf))
		}
	}()
	for k, v := range attributes {
		if fd := yourElement.FieldByName(k); fd.IsValid() {
			if !fd.CanSet() {
				return yourInterface, errors.New(fmt.Sprintf("Field Name '%s' cannot be setted up!!!", fd.Type().Name()))
			}
			fd.Set(reflect.ValueOf(v))
		}
	}
	//	for _, field := range fieldList {
	//		if strFld, ok := yourType.FieldByName(field.Name); ok {
	//			if strFld.Type.Name() == field.Type.Name() {
	//				if field.IsNil {
	//					yourValue.FieldByName(field.Name).Set(reflect.Zero(strFld.Type))
	//				} else {
	//					fmt.Println(field.Value.Type().Kind().String())
	//					fmt.Println(field.Value)
	//					//					Invalid Kind = iota
	//					//					Bool
	//					//					Int
	//					//					Int8
	//					//					Int16
	//					//					Int32
	//					//					Int64
	//					//					Uint
	//					//					Uint8
	//					//					Uint16
	//					//					Uint32
	//					//					Uint64
	//					//					Uintptr
	//					//					Float32
	//					//					Float64
	//					//					Complex64
	//					//					Complex128
	//					//					Array
	//					//					Chan
	//					//					Func
	//					//					Interface
	//					//					Map
	//					//					Ptr
	//					//					Slice
	//					//					String
	//					//					Struct
	//					//					UnsafePointer
	//					if !yourValue.FieldByName(field.Name).CanSet() {
	//						fmt.Println(fmt.Sprintf("Cannot set Value for field %s!!!", field.Name))
	//						continue
	//					}
	//					switch field.Value.Type().Kind() {
	//					case reflect.String:
	//						var value string = string(field.Value.String())
	//						reflect.ValueOf(yourInterface).SetString(value)
	//						//						yourValue.FieldByName(field.Name).SetString(value)
	//					case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int:
	//						var value int64 = int64(field.Value.Int())
	//						yourValue.FieldByName(field.Name).SetInt(value)
	//					case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
	//						var value uint64 = uint64(field.Value.Uint())
	//						yourValue.FieldByName(field.Name).SetUint(value)
	//					case reflect.Float32, reflect.Float64:
	//						var value float64 = float64(field.Value.Float())
	//						yourValue.FieldByName(field.Name).SetFloat(value)
	//					case reflect.Complex64, reflect.Complex128:
	//						var value complex128 = complex128(field.Value.Complex())
	//						yourValue.FieldByName(field.Name).SetComplex(value)
	//						//					case reflect.Array:
	//						//						var value complex128 = complex128(field.Value.)
	//						//						yourValue.FieldByName(field.Name).SetComplex(value)
	//					case reflect.Bool:
	//						var value bool = field.Value.Bool()
	//						yourValue.FieldByName(field.Name).SetBool(value)
	//					case reflect.Ptr:
	//						var value uintptr = field.Value.Pointer()
	//						yourValue.FieldByName(field.Name).SetPointer(unsafe.Pointer(value))
	//					default:
	//						yourValue.FieldByName(field.Name).Set(field.Value)
	//
	//					}
	//				}
	//			} else {
	//				//Wrong type
	//				wrongFieldList = append(wrongFieldList, field)
	//			}
	//		} else {
	//			// discarded
	//			discardedFieldList = append(discardedFieldList, field)
	//		}
	//	}
	//	var messages []string
	//	if len(wrongFieldList) > 0 {
	//		//Report wrong (nor present in the object) fields in the map
	//		for _, wf := range wrongFieldList {
	//			messages = append(messages, fmt.Sprintf("Error parsing field %s : Wrong type in target interface", wf.Name))
	//		}
	//	}
	//	if len(discardedFieldList) > 0 {
	//		//Report discarded fields in the map
	//		for _, df := range discardedFieldList {
	//			messages = append(messages, fmt.Sprintf("Error parsing field %s : Not present in target interface", df.Name))
	//		}
	//	}
	//	if len(messages) > 0 {
	//		var errorMessage string = "Follwong errors has been caught in  : "
	//		for _, message := range messages {
	//			errorMessage = fmt.Sprintf("%s\n%s", errorMessage, message)
	//		}
	//		err = errors.New(errorMessage)
	//	}
	return yourInterface, err
}

// Returns a new Emty serializable object instance
//
// Returns:
//   Serializable - Empty Serializable interface instance
func NewSerializable() Serializable {
	return &SerializableStruct{}
}

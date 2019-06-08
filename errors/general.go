package errors

import (
	"reflect"
)

type error interface {
	Error() string
}

type baseErrorStruct struct {
	s string
}

func (e *baseErrorStruct) Error() string {
	return e.s
}

func New(text string) error {
	return &baseErrorStruct{text}
}

func IsError(obj interface{}) bool {
	if obj == nil {
		return false
	}
	st := reflect.TypeOf(obj)
	_, ok := st.MethodByName("Error")
	return ok

}

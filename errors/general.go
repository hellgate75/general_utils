package errors

import ()

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

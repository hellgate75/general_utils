package streams

import (
	errors "github.com/hellgate75/general_utils/errors"
)

type Type interface{}

type pipelineStruct struct {
	_in      chan Type
	_out     chan Type
	_running bool
}

type Pipeline interface {
	GetOutChannel() *chan Type
	Start()
	Stop()
	Running() bool
	Write(Type)
	Read() (Type, errors.Exception)
}

func (this *pipelineStruct) GetOutChannel() *chan Type {
	return &this._out
}

func (this *pipelineStruct) Running() bool {
	return this._running
}

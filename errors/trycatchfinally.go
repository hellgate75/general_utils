package errors

import ()

type TryCatchFinallyBlock struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

type Exception interface{}

func Throw(up Exception) {
	panic(up)
}

func (tcf TryCatchFinallyBlock) Do() {
	if tcf.Finally != nil {
		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			} else {
				tcf.Catch(New("Undefined Error"))
			}
		}()
	}
	tcf.Try()
}

func DoTryCatch(try func(), catch func(e Exception)) {
	TryCatchFinallyBlock{
		Try:     try,
		Catch:   catch,
		Finally: nil,
	}.Do()
}

func DoTryCatchFinally(try func(), catch func(e Exception), finally func()) {
	TryCatchFinallyBlock{
		Try:     try,
		Catch:   catch,
		Finally: finally,
	}.Do()
}

package errors

import ()

//Generic exception descriptor
type Exception interface{}

//Trow new Exception
func Throw(up Exception) {
	panic(up)
}

//Try function descriptor with signature : func(...interface{}) interface{}
// Executes Code (Java like code)
type DoTry func(...interface{}) interface{}

//Catch function descriptor with signature : func(Exception)
// Executes Exception Catch (Java like code)
type DoCatch func(Exception)

//Finally function descriptor with signature : func()
// Executes Post Block operations (Java like code)
type DoFinally func()

type tryCatchFinallyStruct struct {
	TryFunc      DoTry
	CatchFuncs   []DoCatch
	FinallyFuncs []DoFinally
}

//Describes Try..Catch..Finally Block features
type TryCatchFinallyBlock interface {
	// Add New Exception Catch Function
	//
	//  Parameters:
	//    catch (errors.DoCatch) Exception Code Catch Function
	//
	//  Returns:
	//    TryCatchFinallyBlock the Try..Catch..Finally Block
	Catch(catch DoCatch) TryCatchFinallyBlock
	// Add New Finally Post Block execution Function
	//
	//  Parameters:
	//    finally (errors.DoFinally) Post Execution Code Function
	//
	//  Returns:
	//    TryCatchFinallyBlock the Try..Catch..Finally Block
	Finally(finally DoFinally) TryCatchFinallyBlock
	//Executes the entire Try..Catch..Finally Block
	//
	//  Parameters:
	//    args (...interface{}) Try Function Arguments
	//
	//  Returns:
	//    intrface{} Try Function Execution outcome
	Do(args ...interface{}) interface{}
}

func (tcf *tryCatchFinallyStruct) Catch(catch DoCatch) TryCatchFinallyBlock {
	tcf.CatchFuncs = append(tcf.CatchFuncs, catch)
	return tcf
}

func (tcf *tryCatchFinallyStruct) Finally(finally DoFinally) TryCatchFinallyBlock {
	tcf.FinallyFuncs = append(tcf.FinallyFuncs, finally)
	return tcf
}
func (tcf *tryCatchFinallyStruct) Do(args ...interface{}) interface{} {
	defer func() {
		for _, finally := range tcf.FinallyFuncs {
			finally()
		}
	}()
	defer func() {
		if r := recover(); r != nil {
			for _, catch := range tcf.CatchFuncs {
				catch(r)
			}
		} else {
			for _, catch := range tcf.CatchFuncs {
				catch(New("Undefined Error"))
			}
		}
	}()
	return tcf.TryFunc(args)
}

//Creates a new Try..Catch..Finally Block
//
//  Parameters:
//     try (errors.DoTry) Try Code Block
//
//  Returns:
//     TryCatchFinallyBlock Try..Catch..Finally Block
//  Usage:
//  Try(....).Catch(...).Finally(...).Do(...) for full exception lifecycle or
//  Try(....).Finally(...).Do(...) for making just post execution computation or
//  Try(....).Catch(...).Do(...) for catching errors or even
//  Try(....).Do(...) for simply executing code, in a further time
func Try(try DoTry) TryCatchFinallyBlock {
	return &tryCatchFinallyStruct{
		TryFunc: try,
	}
}

type multiTryCatchFinallyStruct struct {
	TryFuncs     []DoTry
	CatchFuncs   []DoCatch
	FinallyFuncs []DoFinally
}

//Describes Multiple Try Block for Try..Catch..Finally Block features
type MultiTryCatchFinallyBlock interface {
	// Add New Code Try Function
	//
	//  Parameters:
	//    try (errors.DoTry) Code Try Function
	//
	//  Returns:
	//    MultiTryCatchFinallyBlock the Try..Catch..Finally Block
	Try(try DoTry) MultiTryCatchFinallyBlock
	// Add New Exception Catch Function
	//
	//  Parameters:
	//    catch (errors.DoCatch) Exception Code Catch Function
	//
	//  Returns:
	//    MultiTryCatchFinallyBlock the Try..Catch..Finally Block
	Catch(catch DoCatch) MultiTryCatchFinallyBlock
	// Add New Finally Post Block execution Function
	//
	//  Parameters:
	//    finally (errors.DoFinally) Post Execution Code Function
	//
	//  Returns:
	//    MultiTryCatchFinallyBlock the Try..Catch..Finally Block
	Finally(finally DoFinally) MultiTryCatchFinallyBlock
	//Executes the entire Try..Catch..Finally Block
	//
	//  Parameters:
	//    args (...interface{}) Try Function Arguments
	//
	//  Returns:
	//    intrface{} Try Function Execution outcome
	Do(args ...interface{}) []interface{}
}

func (tcf *multiTryCatchFinallyStruct) Try(try DoTry) MultiTryCatchFinallyBlock {
	tcf.TryFuncs = append(tcf.TryFuncs, try)
	return tcf
}

func (tcf *multiTryCatchFinallyStruct) Catch(catch DoCatch) MultiTryCatchFinallyBlock {
	tcf.CatchFuncs = append(tcf.CatchFuncs, catch)
	return tcf
}

func (tcf *multiTryCatchFinallyStruct) Finally(finally DoFinally) MultiTryCatchFinallyBlock {
	tcf.FinallyFuncs = append(tcf.FinallyFuncs, finally)
	return tcf
}
func (tcf *multiTryCatchFinallyStruct) Do(args ...interface{}) []interface{} {
	defer func() {
		for _, finally := range tcf.FinallyFuncs {
			finally()
		}
	}()
	defer func() {
		if r := recover(); r != nil {
			for _, catch := range tcf.CatchFuncs {
				catch(r)
			}
		} else {
			for _, catch := range tcf.CatchFuncs {
				catch(New("Undefined Error"))
			}
		}
	}()
	var out []interface{}
	for _, try := range tcf.TryFuncs {
		out = append(out, try(args))
	}

	return out
}

//Creates a new Mlti Code Try : Try..Catch..Finally Block
//
//  Parameters:
//     try (errors.DoTry) Try Code Block to be added
//
//  Returns:
//     TryCatchFinallyBlock Try..Catch..Finally Block
//  Usage:
//  Try(....).Try(....).Catch(...).Finally(...).Do(...) for full exception lifecycle or
//  Try(....).Try(....).Finally(...).Do(...) for making just post execution computation or
//  Try(....).Try(....).Catch(...).Do(...) for catching errors or even
//  Try(....).Try(....).Try(....).Do(...) for simply executing code, in a further time
func MultiTry(try DoTry) MultiTryCatchFinallyBlock {
	var mtcfb multiTryCatchFinallyStruct = multiTryCatchFinallyStruct{}
	return mtcfb.Try(try)
}

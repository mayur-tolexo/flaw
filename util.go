package flaw

import (
	"fmt"
	"runtime"
	"strings"
)

var (
	debug      bool
	traceDepth int
)

const packageName = "github.com/mayur-tolexo/flaw"

func init() {
	debug = true
	traceDepth = 3
}

//SetDebug will set debug value for error
func SetDebug(flag bool) {
	debug = flag
}

//newError : Create new *Error object
func newError(err error, code int, debugMsg ...string) *Error {
	tracePath := ""
	//If ty debug is enable then only stackTrace i.e. header is sending DEBUG-TACHYON:true
	if debug {
		funcName, fileName, line := stackTrace(traceDepth)
		tracePath = fmt.Sprintf("%v:%v %v()", fileName, line, funcName)
	}

	dMsg := strings.Join(debugMsg, " ")
	msg, exists := errMap[code]
	if !exists {
		msg = dMsg
	}
	return &Error{
		Msg:      msg,
		DebugMsg: dMsg,
		Trace:    tracePath,
		Code:     code,
		Info:     make(map[string]interface{}),
	}
}

//Error : Implement Error method of error interface
func (e *Error) Error() string {
	return fmt.Sprintf("\nCode:\t\t[%d]\nMessage:\t[%v]\nStackTrace:\t[%v]\nDebugMsg:\t[%v]\n", e.Code, e.Msg, e.Trace, e.DebugMsg)
}

//SetMsg will overwrite msg in error
func (e *Error) SetMsg(msg string) *Error {
	if msg != "" {
		e.Msg = msg
	}
	return e
}

//IfCodeSetMsg will set msg if error code matches
func (e *Error) IfCodeSetMsg(errCode int, msg string) *Error {
	if e.Code == errCode {
		e.SetMsg(msg)
	}
	return e
}

//stackTrace : Get function name, file name and line no of the caller function
//Depth is the value from which it will start searching in the stack
func stackTrace(depth int) (funcName string, file string, line int) {
	var (
		ok bool
		pc uintptr
	)
	for i := depth; ; i++ {
		if pc, file, line, ok = runtime.Caller(i); ok {
			if strings.Contains(file, packageName) {
				continue
			}
			fileName := strings.Split(file, "github.com")
			if len(fileName) > 1 {
				file = fileName[1]
			}
			_, funcName = packageFuncName(pc)
			break
		} else {
			break
		}
	}
	return
}

//packageFuncName : Package and function name from package counter
func packageFuncName(pc uintptr) (packageName string, funcName string) {
	if f := runtime.FuncForPC(pc); f != nil {
		funcName = f.Name()
		if ind := strings.LastIndex(funcName, "/"); ind > 0 {
			packageName += funcName[:ind+1]
			funcName = funcName[ind+1:]
		}
		if ind := strings.Index(funcName, "."); ind > 0 {
			packageName += funcName[:ind]
			funcName = funcName[ind+1:]
		}
	}
	return
}

//AppendDebug will append debug msg
func AppendDebug(err error, dMsg string) error {
	switch err.(type) {
	case *Error:
		nErr := err.(*Error)
		nErr.DebugMsg += " " + dMsg
		err = nErr
	}
	return err
}

//GetMsg will return error msg
func GetMsg(err error) (msg string) {
	switch err.(type) {
	case *Error:
		msg = err.(*Error).Msg
	default:
		msg = err.Error()
	}
	return
}

//GetInfo will return error info
func GetInfo(err error) (info map[string]interface{}) {
	switch err.(type) {
	case *Error:
		info = err.(*Error).Info
	default:
	}
	return
}

//GetDebug will return debug msg and trace
func GetDebug(err error) (dMsg, trace string) {
	switch err.(type) {
	case *Error:
		e := err.(*Error)
		dMsg, trace = e.DebugMsg, e.Trace
	default:
		dMsg = err.Error()
	}
	return
}

func snakeName(name string) (capName string) {
	capName = ""
	for i := range name {
		v := name[i]
		if name[i] <= 'Z' {
			if i != 0 {
				capName += " "
			}
		}
		capName += string(v)
	}
	return
}
